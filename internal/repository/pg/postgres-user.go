package pg

import (
	"context"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"
	"github.com/veremchukvv/stonks-test/internal/models"
	"github.com/veremchukvv/stonks-test/pkg/logging"
)

var ErrVkUserNotFound = errors.New("VK user not found")
var ErrGoogleUserNotFound = errors.New("Google user not found")

type PostgresUserRepo struct {
	db  *pgxpool.Pool
	ctx context.Context
}

func NewPostgresUserRepo(pgpool *pgxpool.Pool, ctx context.Context) *PostgresUserRepo {
	return &PostgresUserRepo{
		pgpool,
		ctx}
}

func (ur *PostgresUserRepo) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	log := logging.FromContext(ctx)
	const query string = `SELECT user_id, user_name, lastname, email, password_hash FROM users WHERE email=$1`
	var u models.User
	err := ur.db.QueryRow(ctx, query, email).Scan(&u.Id, &u.Name,
		&u.Lastname, &u.Email, &u.Password)
	if err != nil {
		log.Infof("Can't get user: %v", err)
		return nil, err
	}
	return &u, nil
}

func (ur *PostgresUserRepo) GetVKUserByID(ctx context.Context, id int) (*models.User, error) {
	log := logging.FromContext(ctx)
	const query string = `SELECT user_id, user_auth_type, user_name, lastname, email FROM users WHERE (user_id=$1 and user_auth_type='vk')`
	var vu models.User
	err := ur.db.QueryRow(ctx, query, id).Scan(&vu.Id, &vu.AuthType, &vu.Name,
		&vu.Lastname, &vu.Email)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			log.Infof("Can't get VK User: %v", err)
			return nil, ErrVkUserNotFound
		}
		return nil, err
	}
	return &vu, nil
}

func (ur *PostgresUserRepo) GetGoogleUserByID(ctx context.Context, gid string) (*models.User, error) {
	log := logging.FromContext(ctx)
	const query string = `SELECT user_id, user_auth_type, user_name, lastname, email FROM users WHERE (google_id=$1 and user_auth_type='google')`
	var gu models.User
	err := ur.db.QueryRow(ctx, query, gid).Scan(&gu.Id, &gu.AuthType, &gu.Name,
		&gu.Lastname, &gu.Email)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			log.Infof("Can't get Google User: %v", err)
			return nil, ErrGoogleUserNotFound
		}
		return nil, err
	}
	return &gu, nil
}

func (ur *PostgresUserRepo) GetUserByID(ctx context.Context, id int, auth_type string) (*models.User, error) {
	log := logging.FromContext(ctx)

	if auth_type == "local" {
		const query string = `SELECT user_id, user_auth_type, user_name, lastname, email, password_hash FROM users 
                              WHERE (user_id=$1 and user_auth_type='local')`
		var u models.User
		err := ur.db.QueryRow(ctx, query, id).Scan(&u.Id, &u.AuthType, &u.Name,
			&u.Lastname, &u.Email, &u.Password)

		if err != nil {
			log.Infof("Can't get user: %v", err)
			return nil, err
		}
		return &u, nil
	}

	const query string = `SELECT user_id, user_auth_type, user_name, lastname FROM users WHERE user_id=$1`
	var u models.User
	err := ur.db.QueryRow(ctx, query, id).Scan(&u.Id, &u.AuthType, &u.Name,
		&u.Lastname)

	if err != nil {
		log.Infof("Can't get user: %v", err)
		return nil, err
	}
	return &u, nil
}

func (ur *PostgresUserRepo) CreateUser(ctx context.Context, user *models.User) (*models.User, error) {
	log := logging.FromContext(ctx)
	const query string = `INSERT INTO users 
		(user_name, user_auth_type, lastname, email, password_hash) 
		VALUES($1, $2, $3, $4, $5)
		returning user_id;`

	var uid int
	err := ur.db.QueryRow(ctx, query, user.Name, "local", user.Lastname, user.Email, user.Password).Scan(&uid)
	if err != nil {
		log.Errorf("Error on write user to database: %v", err)
		return nil, err
	}
	user.Id = uid
	return user, nil
}

func (ur *PostgresUserRepo) UpdateUser(ctx context.Context, user *models.User) (*models.User, error) {
	log := logging.FromContext(ctx)
	const query string = `UPDATE users SET (email, modified_at) = ($1, NOW()) WHERE user_id=$2 and user_auth_type=$3 returning user_id`

	var uid int
	err := ur.db.QueryRow(ctx, query, user.Email, user.Id, user.AuthType).Scan(&uid)
	if err != nil {
		log.Errorf("Error on write user to database: %v", err)
		return nil, err
	}
	user.Id = uid
	return user, nil
}

func (ur *PostgresUserRepo) DeleteUser(ctx context.Context, userId int, authType string) error {
	log := logging.FromContext(ctx)
	const query string = `DELETE FROM users WHERE user_id=$1 and user_auth_type=$2 returning user_id`

	var uid int
	err := ur.db.QueryRow(ctx, query, userId, authType).Scan(&uid)
	if err != nil {
		log.Errorf("Error on delete user %d from database: %v", userId, err)
		return err
	}
	return nil
}

func (ur *PostgresUserRepo) CreateVKUser(ctx context.Context, user *models.User) (*models.User, error) {
	log := logging.FromContext(ctx)
	const query string = `INSERT INTO users 
		(user_id, user_auth_type, user_name, lastname, email) 
		VALUES($1, $2, $3, $4, $5)
		returning user_id;`

	var uid int
	err := ur.db.QueryRow(ctx, query, user.Id, "vk", user.Name, user.Lastname, user.Email).Scan(&uid)
	if err != nil {
		log.Errorf("Error on write user to database: %v", err)
		return nil, err
	}
	user.Id = uid
	return user, nil
}

func (ur *PostgresUserRepo) CreateGoogleUser(ctx context.Context, user *models.User) (*models.User, error) {
	log := logging.FromContext(ctx)
	const query string = `INSERT INTO users 
		(google_id, user_auth_type, user_name, lastname, email) 
		VALUES($1, $2, $3, $4, $5)
		returning user_id;`

	var uid int
	err := ur.db.QueryRow(ctx, query, user.GoogleId, "google", user.Name, user.Lastname, user.Email).Scan(&uid)
	if err != nil {
		log.Errorf("Error on write user to database: %v", err)
		return nil, err
	}
	user.Id = uid
	return user, nil
}