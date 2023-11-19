package hubLogRepo

import (
	"context"
	"fmt"
)

func (d HubLogRepo) InitProduct(ctx context.Context, productID string) error {
	sql := fmt.Sprintf("CREATE STABLE IF NOT EXISTS %s "+
		"(`ts` timestamp,`content` BINARY(5000),`topic` BINARY(500), `action` BINARY(100),"+
		" `request_id` BINARY(100), `trance_id` BINARY(100), `result_type` BIGINT)"+
		"TAGS (`product_id` BINARY(50),`device_name`  BINARY(50));",
		d.GetLogStableName())
	if _, err := d.t.ExecContext(ctx, sql); err != nil {
		return err
	}
	return nil
}

func (d HubLogRepo) DropProduct(ctx context.Context, productID string) error {
	return nil
}

func (d HubLogRepo) DropDevice(ctx context.Context, productID string, deviceName string) error {
	sql := fmt.Sprintf("drop table if exists %s;", d.GetLogTableName(productID, deviceName))
	if _, err := d.t.ExecContext(ctx, sql); err != nil {
		return err
	}
	return nil
}
