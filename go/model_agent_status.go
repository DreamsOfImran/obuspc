/*
 * User Services Platform - Controller API
 *
 * This is the REST API for the User Services Platform (USP) Controller.  This is how an external entity can interact with a USP Controller to retrieve information from or configure a USP Agent.
 *
 * API version: 1.0.0
 * Contact: info@broadband-forum.org
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type AgentStatus struct {

	CommunicationsStatus string `json:"communicationsStatus,omitempty"`
}
