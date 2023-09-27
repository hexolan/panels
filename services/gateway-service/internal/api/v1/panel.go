package v1

import (
	"time"
	"context"

	"github.com/gofiber/fiber/v2"

	"github.com/hexolan/panels/gateway-service/internal/rpc"
	"github.com/hexolan/panels/gateway-service/internal/rpc/panelv1"
)

func getPanelById(id string) (*panelv1.Panel, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	return rpc.Svcs.GetPanelSvc().GetPanel(
		ctx,
		&panelv1.GetPanelByIdRequest{Id: id},
	)
}

func getPanelIDFromName(name string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	panel, err := rpc.Svcs.GetPanelSvc().GetPanelByName(
		ctx,
		&panelv1.GetPanelByNameRequest{Name: name},
	)
	if err != nil {
		return "", err
	}
	
	return panel.GetId(), nil
}

func GetPanelById(c *fiber.Ctx) error {
	panel, err := getPanelById(c.Params("id"))
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{"status": "success", "data": panel})
}

func GetPanelByName(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	panel, err := rpc.Svcs.GetPanelSvc().GetPanelByName(
		ctx,
		&panelv1.GetPanelByNameRequest{Name: c.Params("name")},
	)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{"status": "success", "data": panel})
}

func UpdatePanelById(c *fiber.Ctx) error {
	// check user can update panels
	currentUser, err := getCurrentUser(c)
	if err != nil {
		return err
	}

	if !currentUser.IsAdmin {
		return fiber.NewError(fiber.StatusForbidden, "no permissions to update that panel")
	}

	// update the panel
	patchData := new(panelv1.PanelMutable)
	if err := c.BodyParser(patchData); err != nil {
		fiber.NewError(fiber.StatusBadRequest, "malformed request")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	panel, err := rpc.Svcs.GetPanelSvc().UpdatePanel(
		ctx, 
		&panelv1.UpdatePanelByIdRequest{Id: c.Params("id"), Data: patchData},
	)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{"status": "success", "data": panel})
}

func UpdatePanelByName(c *fiber.Ctx) error {
	// check user can update panels
	currentUser, err := getCurrentUser(c)
	if err != nil {
		return err
	}

	if !currentUser.IsAdmin {
		return fiber.NewError(fiber.StatusForbidden, "no permissions to update that panel")
	}

	// update the panel
	patchData := new(panelv1.PanelMutable)
	if err := c.BodyParser(patchData); err != nil {
		fiber.NewError(fiber.StatusBadRequest, "malformed request")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	panel, err := rpc.Svcs.GetPanelSvc().UpdatePanelByName(
		ctx, 
		&panelv1.UpdatePanelByNameRequest{Name: c.Params("name"), Data: patchData},
	)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{"status": "success", "data": panel})
}

func DeletePanelById(c *fiber.Ctx) error {
	// check user can delete panels
	currentUser, err := getCurrentUser(c)
	if err != nil {
		return err
	}

	if !currentUser.IsAdmin {
		return fiber.NewError(fiber.StatusForbidden, "no permissions to delete that panel")
	}

	// delete the panel
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	_, err = rpc.Svcs.GetPanelSvc().DeletePanel(
		ctx,
		&panelv1.DeletePanelByIdRequest{Id: c.Params("id")},
	)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{"status": "success", "msg": "panel deleted"})
}

func DeletePanelByName(c *fiber.Ctx) error {
	// check user can delete panels
	currentUser, err := getCurrentUser(c)
	if err != nil {
		return err
	}

	if !currentUser.IsAdmin {
		return fiber.NewError(fiber.StatusForbidden, "no permissions to update that panel")
	}

	// delete the panel
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	_, err = rpc.Svcs.GetPanelSvc().DeletePanelByName(
		ctx,
		&panelv1.DeletePanelByNameRequest{Name: c.Params("name")},
	)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{"status": "success", "msg": "panel deleted"})
}

func CreatePanel(c *fiber.Ctx) error {
	newPanel := new(panelv1.PanelMutable)
	if err := c.BodyParser(newPanel); err != nil {
		fiber.NewError(fiber.StatusBadRequest, "malformed request")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	panel, err := rpc.Svcs.GetPanelSvc().CreatePanel(
		ctx,
		&panelv1.CreatePanelRequest{Data: newPanel},
	)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{"status": "success", "data": panel})
}