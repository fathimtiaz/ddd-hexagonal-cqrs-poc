package usecase

import (
	"context"
	"fmt"

	"ddd-hexagonal-cqrs-poc/services/common/io/input"
	"ddd-hexagonal-cqrs-poc/services/serviceA/io/input/command"
	"ddd-hexagonal-cqrs-poc/services/serviceA/outbound/repo"
	serviceb "ddd-hexagonal-cqrs-poc/services/serviceA/outbound/serviceB"
	serviceBCmd "ddd-hexagonal-cqrs-poc/services/serviceB/io/input/command"
)

type EntityAMoveUseCase struct {
	reader repo.Reader
	writer repo.Writer
	svcB   serviceb.CommandCallers
}

func NewEntityAMoveUseCase(r repo.Reader, w repo.Writer, s serviceb.CommandCallers) *EntityAMoveUseCase {
	return &EntityAMoveUseCase{reader: r, writer: w, svcB: s}
}

func (u *EntityAMoveUseCase) Execute(ctx context.Context, commander input.Commander) error {
	cmd, ok := commander.(command.EntityAMoveCommand)
	if !ok {
		return fmt.Errorf("invalid command type")
	}
	entityA, err := u.reader.EntityA(ctx, cmd.Query)
	if err != nil {
		return fmt.Errorf("read entity: %w", err)
	}

	if !entityA.CanMove() {
		return fmt.Errorf("entity cannot move")
	}

	entityA.Move(cmd.Payload.Direction)

	u.writer.EntityASave(ctx, entityA)

	// Synchronous inter-service call
	sbCmd := serviceBCmd.EntityAHasMovedCommand{}
	if err := u.svcB.EntityAHasMoved(ctx, sbCmd); err != nil {
		return fmt.Errorf("serviceB call failed: %w", err)
	}

	return nil
}
