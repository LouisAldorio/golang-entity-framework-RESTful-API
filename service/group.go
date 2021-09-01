package service

import (
	"context"
	"fmt"
	"log"
	"myapp/ent"
	"myapp/ent/group"
)

func GroupsGetAll(ctx context.Context, client *ent.Client) ([]*ent.Group, error) {

	groups, err := client.Group.Query().WithUsers(
		func(uq *ent.UserQuery) {
			uq.WithCars(
				func(cq *ent.CarQuery) {
					cq.WithCarUtilities()
				},
			)
		},
	).All(ctx)
	if err != nil {
		return nil, err
	}

	return groups, nil
}

func GroupsGetByID(ctx context.Context, client *ent.Client, groupID int) (*ent.Group, error) {
	
	group, err := client.Group.Query().WithUsers(
		func(uq *ent.UserQuery) {
			uq.WithCars(
				func(cq *ent.CarQuery) {
					cq.WithCarUtilities()
				},
			)
		},
	).Where(group.ID(groupID)).
		Only(ctx)
	if err != nil {
		return nil, err
	}

	return group, nil
}

// Get all groups that have users
func QueryGroupWithUsers(ctx context.Context, client *ent.Client) error {
	groups, err := client.Group.
		Query().
		Where(group.HasUsers()).
		All(ctx)
	if err != nil {
		return fmt.Errorf("failed getting groups: %w", err)
	}
	log.Println("groups returned:", groups)
	// Output: (Group(Name=GitHub), Group(Name=GitLab),)
	return nil
}

func QueryGroupByName(ctx context.Context, client *ent.Client, name string) error {
	cars, err := client.Group.
		Query().
		Where(group.Name(name)). // (Group(Name=GitHub),)
		QueryUsers().                // (User(Name=Ariel, Age=30),)
		QueryCars().                 // (Car(Model=Tesla, RegisteredAt=<Time>), Car(Model=Mazda, RegisteredAt=<Time>),)
		All(ctx)
	if err != nil {
		return fmt.Errorf("failed getting cars: %w", err)
	}
	log.Println("cars returned:", cars)
	// Output: (Car(Model=Tesla, RegisteredAt=<Time>), Car(Model=Mazda, RegisteredAt=<Time>),)
	return nil
}
