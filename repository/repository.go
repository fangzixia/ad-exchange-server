package repository

import (
	"context"
	"gorm.io/gorm"
	"log/slog"
	"time"
)

var DB *gorm.DB

type Space struct {
	Id         int
	AdFormat   int
	FloorPrice float32
	Timeout    int
}

type SpaceConfig struct {
	SpaceId        int
	TargetSpaceId  string
	Priority       int
	BidUrl         string
	BidHandlerName string
}

var spacesIdMapping map[int]Space

var spacesIds map[string]int

var configSpacesIdMapping map[int][]*SpaceConfig

func InitDB() {

}

func WarmUp() {
	slog.Info("warm up start", "start_time", time.Now())
	spaces, err := gorm.G[Space](DB).Raw("select id,media_id,ad_format,floor_price,timeout from ssp_space where status = 1").Find(context.Background())
	if err != nil {
		slog.Info("warm up error: ", err)
	}

	spaceConfigs, err := gorm.G[SpaceConfig](DB).Raw(`SELECT
			stc.space_id,
			stcd.target_space_id,
			stc.priority,
			dc.bid_url,
			p.deliver_class_name 
		FROM
			ssp_space sp
			JOIN ssp_space_traffic_config stc ON stc.space_id = sp.id
			JOIN ssp_space_traffic_config_docking stcd ON stcd.config_id = stc.id
			JOIN platform p ON p.id = stc.platform_id
			JOIN platform d ON d.id = p.dsp_id
			JOIN dsp_config dc ON dc.dsp_id = d.id 
		WHERE
			sp.status = 1 
			AND stc.status = 1
	`).Find(context.Background())

	if err != nil {
		slog.Info("warm up error: ", err)
	}

	sim := make(map[int]Space)
	for _, space := range spaces {
		sim[space.Id] = space
	}
	spacesIdMapping = sim

	csim := make(map[int][]*SpaceConfig)
	for _, spaceConfig := range spaceConfigs {
		csim[spaceConfig.SpaceId] = append(csim[spaceConfig.SpaceId], &spaceConfig)
	}
	configSpacesIdMapping = csim
	slog.Info("warm up end", "end", time.Now())
}

func GetSpace(spaceId int) Space {
	return spacesIdMapping[spaceId]
}

func GetSpaceConfig(spaceId int) []*SpaceConfig {
	return configSpacesIdMapping[spaceId]
}
