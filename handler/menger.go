package handler

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/wmsx/menger_svc/models"
	proto "github.com/wmsx/menger_svc/proto/menger"
	"github.com/xianghuzhao/kdfcrypt"
	"golang.org/x/crypto/argon2"
)

const (
	DefaultAvatar = ""
	RandSaltLen   = 16
	PasswordLen   = 32
)

type MengerHandler struct{}

func (*MengerHandler) Register(ctx context.Context, req *proto.RegisterRequest, res *proto.RegisterResponse) error {
	var (
		err             error
		encodedPassword string
	)
	if req.Avatar == "" {
		req.Avatar = DefaultAvatar
	}
	var old *models.Menger
	if old, err = models.GetMengerByEmail(req.Email); err != nil {
		log.Error("查询用户失败 err: ", err)
		return err
	}
	if old != nil {
		res.ErrorMsg = &proto.ErrorMsg{Msg: "邮箱已被注册"}
		return nil
	}
	if old, err = models.GetMengerByName(req.Name); err != nil {
		log.Error("查询用户失败 err: ", err)
		return err
	}
	if old != nil {
		res.ErrorMsg = &proto.ErrorMsg{Msg: "用户名已被使用"}
		return nil
	}

	encodedPassword, err = kdfcrypt.Encode(req.Password, &kdfcrypt.Option{
		Algorithm:        "argon2id",
		Param:            "m=65536,t=1,p=4",
		RandomSaltLength: RandSaltLen,
		HashLength:       PasswordLen,
	})

	if err = models.AddMenger(req.Name, req.Email, encodedPassword, req.Avatar); err != nil {
		log.Error("注册用户失败 err: ", err)
		return err
	}
	return nil
}

func (*MengerHandler) Login(ctx context.Context, req *proto.LoginRequest, res *proto.LoginResponse) error {
	var (
		err      error
		verified bool
	)
	var menger *models.Menger
	if menger, err = models.GetMengerByEmailOrName(req.Name, req.Email); err != nil {
		log.Error("根据邮箱或用户名查找用户失败 err: ", err)
		return err
	}
	if menger == nil {
		res.ErrorMsg = &proto.ErrorMsg{Msg: "信息不正确"}
		return nil
	}

	if verified, err = kdfcrypt.Verify(req.Password, menger.Password); err != nil {
		log.Error("验证密码失败 err: ", err)
		return err
	}
	if !verified {
		log.Error("密码验证未通过")
		res.ErrorMsg = &proto.ErrorMsg{Msg: "信息不正确"}
	}

	res.MengerInfo = &proto.MengerInfo{
		Id:     menger.ID,
		Name:   menger.Name,
		Email:  menger.Email,
		Avatar: menger.Avatar,
	}
	return nil
}

func (*MengerHandler) Logout(context.Context, *proto.LogoutRequest, *proto.LogoutResponse) error {
	panic("implement me")
}

func (*MengerHandler) GetMenger(context.Context, *proto.GetMengerRequest, *proto.GetMengerResponse) error {
	return nil
}

func (m *MengerHandler) GetMengerList(ctx context.Context, req *proto.GetMengerListRequest, res *proto.GetMengerListResponse) error {
	mengers, err := models.GetMengerByIds(req.MengerIds)
	if err != nil {
		return err
	}

	mengerInfos := make([]*proto.MengerInfo, 0)
	for _, menger := range mengers {
		mengerInfo := &proto.MengerInfo{
			Id:     menger.ID,
			Name:   menger.Name,
			Email:  menger.Email,
			Avatar: menger.Avatar,
		}
		mengerInfos = append(mengerInfos, mengerInfo)
	}
	res.MengerInfos = mengerInfos
	return nil
}

func cryptoPassword(password, salt string) string {
	return base64.StdEncoding.EncodeToString(argon2.IDKey([]byte(password), []byte(salt), 1, 64*1024, 4, PasswordLen))
}

// GenerateRandomBytes returns securely generated random bytes.
// It will return an error if the system's secure random
// number generator fails to function correctly, in which
// case the caller should not continue.
func generateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	// Note that err == nil only if we read len(b) bytes.
	if err != nil {
		return nil, err
	}

	return b, nil
}
