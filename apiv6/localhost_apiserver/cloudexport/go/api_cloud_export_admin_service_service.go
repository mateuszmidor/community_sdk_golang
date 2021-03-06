/*
 * Cloud Export Admin API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 202101
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package cloudexportstub

import (
	"context"
	"errors"
	"fmt"
	"net/http"
)

// CloudExportAdminServiceApiService is a service that implents the logic for the CloudExportAdminServiceApiServicer
// This service should implement the business logic for every endpoint for the CloudExportAdminServiceApi API.
// Include any external packages or services that will be required by this service.
type CloudExportAdminServiceApiService struct {
	repo *CloudExportRepo
}

// NewCloudExportAdminServiceApiService creates a default api service
func NewCloudExportAdminServiceApiService(repo *CloudExportRepo) CloudExportAdminServiceApiServicer {
	return &CloudExportAdminServiceApiService{
		repo: repo,
	}
}

// CloudExportAdminServiceCreateCloudExport -
func (s *CloudExportAdminServiceApiService) CloudExportAdminServiceCreateCloudExport(ctx context.Context, v202101beta1CreateCloudExportRequest V202101beta1CreateCloudExportRequest) (ImplResponse, error) {
	if export, err := s.repo.Create(v202101beta1CreateCloudExportRequest.Export); err != nil {
		return errorResponse(http.StatusBadRequest, "cloud export CREATE failed", err), nil
	} else {
		resp := V202101beta1CreateCloudExportResponse{Export: *export}
		return Response(http.StatusOK, &resp), nil
	}
}

// CloudExportAdminServiceUpdateCloudExport -
func (s *CloudExportAdminServiceApiService) CloudExportAdminServiceUpdateCloudExport(ctx context.Context, exportId string, v202101beta1UpdateCloudExportRequest V202101beta1UpdateCloudExportRequest) (ImplResponse, error) {
	v202101beta1UpdateCloudExportRequest.Export.Id = exportId
	if export, err := s.repo.Update(v202101beta1UpdateCloudExportRequest.Export); err != nil {
		return errorResponse(http.StatusBadRequest, "cloud export UPDATE failed", err), nil
	} else {
		resp := V202101beta1UpdateCloudExportResponse{Export: *export}
		return Response(http.StatusOK, &resp), nil
	}
}

// CloudExportAdminServiceDeleteCloudExport -
func (s *CloudExportAdminServiceApiService) CloudExportAdminServiceDeleteCloudExport(ctx context.Context, exportId string) (ImplResponse, error) {
	if err := s.repo.Delete(exportId); err != nil {
		return errorResponse(http.StatusNotFound, "cloud export DELETE failed", err), nil
	} else {
		resp := map[string]interface{}{}
		return Response(http.StatusOK, &resp), nil
	}
}

// CloudExportAdminServiceGetCloudExport -
func (s *CloudExportAdminServiceApiService) CloudExportAdminServiceGetCloudExport(ctx context.Context, exportId string) (ImplResponse, error) {
	if export := s.repo.Get(exportId); export == nil {
		err := fmt.Errorf("no such cloud export %q", exportId)
		return errorResponse(http.StatusNotFound, "cloud export GET failed", err), nil
	} else {
		resp := V202101beta1GetCloudExportResponse{Export: *export}
		return Response(http.StatusOK, &resp), nil
	}
}

// CloudExportAdminServiceListCloudExport -
func (s *CloudExportAdminServiceApiService) CloudExportAdminServiceListCloudExport(ctx context.Context) (ImplResponse, error) {
	resp := V202101beta1ListCloudExportResponse{Exports: s.repo.List()}
	return Response(http.StatusOK, &resp), nil
}

// CloudExportAdminServicePatchCloudExport - not used.
func (s *CloudExportAdminServiceApiService) CloudExportAdminServicePatchCloudExport(ctx context.Context, exportId string, v202101beta1PatchCloudExportRequest V202101beta1PatchCloudExportRequest) (ImplResponse, error) {
	// TODO - update CloudExportAdminServicePatchCloudExport with the required logic for this service method.
	// Add api_cloud_export_admin_service_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, V202101beta1PatchCloudExportResponse{}) or use other options such as http.Ok ...
	//return Response(200, V202101beta1PatchCloudExportResponse{}), nil

	//TODO: Uncomment the next line to return response Response(0, GooglerpcStatus{}) or use other options such as http.Ok ...
	//return Response(0, GooglerpcStatus{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("CloudExportAdminServicePatchCloudExport method not implemented")
}

func errorResponse(httpCode int, message string, err error) ImplResponse {
	const grpcCodeUnknown = 2 // translation httpCode -> grpcCode not relevant here

	grpcResponse := GooglerpcStatus{
		Code:    grpcCodeUnknown,
		Message: fmt.Sprintf("%v: %v", message, err.Error()),
		Details: []ProtobufAny{},
	}
	return Response(httpCode, grpcResponse)
}
