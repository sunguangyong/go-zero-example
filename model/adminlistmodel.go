package model

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ AdminListModel = (*customAdminListModel)(nil)

type AdminUser struct {
	UserName string `db:"user_name"` // tg user_name
	ChatId   int64  `db:"chat_id"`   // tg_id
}

type (
	// AdminListModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAdminListModel.
	AdminListModel interface {
		adminListModel
		CommonFind(ctx context.Context, querySql, orderSql, limitSql string) ([]*AdminList, error)
		CommonDistinctFind(ctx context.Context) ([]*AdminUser, error)
		WithName(userName string) AdminUserOption
		WithChatId(chatId int64) AdminUserOption
		CommonOptionFind(ctx context.Context, opts ...AdminUserOption) ([]*AdminList, error)
		WithOrder() AdminUserOption
	}

	customAdminListModel struct {
		*defaultAdminListModel
	}
)

// NewAdminListModel returns a model for the database table.
func NewAdminListModel(conn sqlx.SqlConn) AdminListModel {
	return &customAdminListModel{
		defaultAdminListModel: newAdminListModel(conn),
	}
}

func (m *defaultAdminListModel) CommonFind(ctx context.Context, querySql, orderSql,
	limitSql string) ([]*AdminList, error) {
	query := fmt.Sprintf("select %s from %s %s %s %s", adminListRows, m.table, querySql,
		orderSql, limitSql)
	var resp []*AdminList
	err := m.conn.QueryRowsCtx(ctx, &resp, query)
	switch err {
	case nil:
		return resp, nil
	case sqlx.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultAdminListModel) CommonDistinctFind(ctx context.Context) ([]*AdminUser, error) {
	query := fmt.Sprintf("select DISTINCT chat_id,user_name from admin_list where is_delete = 0 and user_name != ''")
	var resp []*AdminUser
	err := m.conn.QueryRowsCtx(ctx, &resp, query)
	switch err {
	case nil:
		return resp, nil
	case sqlx.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

type AdminUserOption func(qb squirrel.SelectBuilder) squirrel.SelectBuilder

func (m *defaultAdminListModel) WithName(userName string) AdminUserOption {
	return func(qb squirrel.SelectBuilder) squirrel.SelectBuilder {
		qb = qb.Where("user_name = ?", userName)
		return qb
	}
}

func (m *defaultAdminListModel) WithOrder() AdminUserOption {
	return func(qb squirrel.SelectBuilder) squirrel.SelectBuilder {
		qb = qb.OrderBy("id desc")
		return qb
	}
}

func (m *defaultAdminListModel) WithChatId(chatId int64) AdminUserOption {
	return func(qb squirrel.SelectBuilder) squirrel.SelectBuilder {
		qb = qb.Where("chat_id = ?", chatId)
		return qb
	}
}

func (m *defaultAdminListModel) CommonOptionFind(ctx context.Context, opts ...AdminUserOption) ([]*AdminList, error) {
	qb := squirrel.Select(adminListRows).From(m.table)

	for _, opt := range opts {
		qb = opt(qb)
	}
	//qb = qb.Where("name = 1")
	//qb = qb.Where("age = 2")
	//
	query, args, err := qb.ToSql()

	fmt.Println("err ==== ", err)
	fmt.Println("query ==== ", query, args)

	var resp []*AdminList

	err = m.conn.QueryRowsCtx(ctx, &resp, query, args...)

	switch err {
	case nil:
		return resp, nil
	case sqlx.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
