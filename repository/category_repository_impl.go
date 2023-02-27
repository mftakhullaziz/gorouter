package repository

import (
	"context"
	"database/sql"
	"errors"
	"gohttp/model/domain"
	"gohttp/utils/helper"
)

type CategoryRepositoryImpl struct {
}

func NewCategoryRepositoryImpl() CategoryRepository {
	return &CategoryRepositoryImpl{}
}

func (c CategoryRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	//TODO implement me
	SQL := "insert into category(`name`, `desc`) values (?,?)"
	result, err := tx.ExecContext(ctx, SQL, category.Name, category.Desc)
	helper.PanicIfError(err)
	id, err := result.LastInsertId()
	helper.PanicIfError(err)
	category.Id = int(id)
	return category
}

func (c CategoryRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	//TODO implement me
	SQL := "update category set `name` = ?, `desc` = ? where id = ?"
	_, err := tx.ExecContext(ctx, SQL, category.Name, category.Desc, category.Id)
	helper.PanicIfError(err)
	return category
}

func (c CategoryRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, category domain.Category) {
	//TODO implement me
	SQL := "delete from category where id = ?"
	_, err := tx.ExecContext(ctx, SQL, category.Id)
	helper.PanicIfError(err)
}

func (c CategoryRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, catId int) (domain.Category, error) {
	//TODO implement me
	SQL := "select id, name, `desc` from category where id = ?"
	rows, err := tx.QueryContext(ctx, SQL, catId)
	helper.PanicIfError(err)
	defer rows.Close()

	category := domain.Category{}
	if rows.Next() {
		err := rows.Scan(&category.Id, &category.Name, &category.Desc)
		helper.PanicIfError(err)
		return category, nil
	} else {
		return category, errors.New("category is not found")
	}
}

func (c CategoryRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Category {
	//TODO implement me
	SQL := "select id, name, `desc` from category"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var categories []domain.Category
	for rows.Next() {
		category := domain.Category{}
		err := rows.Scan(&category.Id, &category.Name, &category.Desc)
		helper.PanicIfError(err)
		categories = append(categories, category)
	}
	return categories
}
