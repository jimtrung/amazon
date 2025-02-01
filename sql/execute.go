package sql

import "github.com/jimtrung/amazon/sql/migrations"

func SetupDatabase() error {
	// schema.go
	err := CreateSchema()
	if err != nil {
		return err
	}

	err = migrations.BindDefaultProducts()
	if err != nil {
		return err
	}

	err = migrations.BindRuleProducts()
	if err != nil {
		return err
	}

	err = migrations.BindDefaultCart()
	if err != nil {
		return err
	}

	err = migrations.BindRuleCart()
	if err != nil {
		return err
	}

	err = migrations.CreateAddToCart()
	if err != nil {
		return err
	}

	return nil
}
