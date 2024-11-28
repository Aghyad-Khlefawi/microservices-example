package grpc

import (
	"context"

	"github.com/aghyad-khlefawi/identity/pkg/jwthelper"
	"github.com/aghyad-khlefawi/identity/pkg/users"
	"google.golang.org/protobuf/types/known/structpb"
)

type IdentityServiceImplementation struct {
	UnimplementedIdentityServiceServer
}

func NewIdentityService() *IdentityServiceImplementation {
	return &IdentityServiceImplementation{
		UnimplementedIdentityServiceServer{},
	}
}

func (imp *IdentityServiceImplementation) ValidateToken(ctx context.Context, request *ValidateTokenRequest) (*ValidateTokenResponse, error) {
	isValid, claims, err := jwthelper.VerifyToken(request.Token)
	if err != nil {
		return &ValidateTokenResponse{
			IsValid: false,
			Message: err.Error(),
			Claims:  nil,
		}, nil
	}

	claimsStruct, err := structpb.NewStruct(claims)
	return &ValidateTokenResponse{
		IsValid: isValid,
		Message: "verified",
		Claims:  claimsStruct,
	}, nil
}

func (imo *IdentityServiceImplementation) CreateUser(ctx context.Context, request *CreateUserRequest) (*CreateUserResponse,error){
 user,err:= users.CreateUser(request.Email,request.Password,request.GeneratePassword)
if err!=nil{
		return nil,err
	}
	return &CreateUserResponse{UserId: user.Id.String()},nil
}
