/*
 * ServiceDesk
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type Request struct {
	Id int64 `json:"id,omitempty"`

	Head string `json:"head,omitempty"`

	Body string `json:"body,omitempty"`

	Email string `json:"email,omitempty"`
}
