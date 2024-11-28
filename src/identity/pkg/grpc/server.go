package grpc

import (
	"context"

	"github.com/aghyad-khlefawi/identity/pkg/jwthelper"
)

type IdentityServiceImplementation struct{
	UnimplementedIdentityServiceServer
}


func(imp *IdentityServiceImplementation) ValidateToken(ctx *context.Context, request *ValidateTokenRequest) (*ValidateTokenResponse,error){
 isValid,err := jwthelper.VerifyToken(request.Token)
	if err!=nil{
		return &ValidateTokenResponse{
			IsValid:false,
			Message: err.Error(),
		},nil
	}

	return &ValidateTokenResponse{
		IsValid: isValid,
		Message: "verified",
	},nil
}
