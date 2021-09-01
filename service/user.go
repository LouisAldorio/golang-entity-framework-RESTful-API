package service

import (
	"context"
	"fmt"
	"log"
	"myapp/ent"
	"myapp/model"
)

func CreateUser(ctx context.Context, client *ent.Client, param model.UserCreateInput) (*ent.User, error) {

	user, err := client.User.
		Create().
		SetAge(param.Age).
		SetName(param.Username).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating user: %w", err)
	}
	log.Println("user was created: ", user)

	for _, v := range param.Cars {
		_, err := CreateCar(ctx, user, client, v)
		if err != nil {
			return nil, fmt.Errorf("error creating cars %s", err)
		}
	}

	return user, nil
}

func QueryUser(ctx context.Context, client *ent.Client, name string) ([]*ent.User, error) {

	u, err := client.User.
		Query().
		WithCars(func(q *ent.CarQuery) {
			q.WithCarUtilities()
		}).
		// `Only` fails if no user found,
		// or more than 1 user returned.
		All(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying user: %w", err)
	}

	log.Println("user returned: ", u)
	return u, nil
}
