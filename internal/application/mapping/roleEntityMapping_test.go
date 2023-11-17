package mapping

import (
	utlits "github.com/3110Y/cc-utlits"
	"github.com/3110Y/profile/internal/application/dto"
	"github.com/3110Y/profile/internal/domain/entity"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRoleEntityMapping_ToRoleDTO(t *testing.T) {
	entityProfile := entity.Role{
		Id:   "123456789",
		Name: "Name",
	}
	profileDTO := dto.RoleDTO{
		Id:   entityProfile.Id,
		Name: entityProfile.Name,
	}
	profileDTOFilled := utlits.Pointer(RoleEntityMapping{Entity: entityProfile}).ToRoleDTO()
	assert.Equal(t, profileDTO, profileDTOFilled)
}
