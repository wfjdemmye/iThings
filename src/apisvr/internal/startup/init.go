package startup

import (
	"context"
	"fmt"
	"github.com/i-Things/things/shared/utils"
	"github.com/zeromicro/go-zero/core/logx"

	"github.com/i-Things/things/src/apisvr/internal/svc"
	"github.com/i-Things/things/src/dmsvr/client/firmwaremanage"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func StartOtaChanWalk(s *svc.ServiceContext) {
	s.FileChan = make(chan int64, 100)
	go FileChanWalk(s)
}

func FileChanWalk(s *svc.ServiceContext) {
	ctx := context.Background()
	//处理因为宕机未执行的file
	old := &firmwaremanage.OtaFirmwareFileIndexReq{
		Size: &wrapperspb.Int64Value{
			Value: 0,
		},
	}
	fileList, err := s.FirmwareM.OtaFirmwareFileIndex(ctx, old)
	if err != nil {
		logx.Errorf("%v.OtaFirmwareFileIndex err:%v", utils.FuncName(), err)
		return
	}
	for _, f := range fileList.List {
		s.FileChan <- f.FirmwareID
	}
	//chan
	for {
		firmwareID := <-s.FileChan
		in := &firmwaremanage.FirmwareInfoReadReq{
			FirmwareID: firmwareID,
		}
		firmwareInfo, err := s.FirmwareM.FirmwareInfoRead(ctx, in)
		if err != nil {
			fmt.Println(err)
			continue
		}
		for _, f := range firmwareInfo.Files {
			//fmt.Println(f.FilePath, " chain dir")
			storageInfo, _ := s.OssClient.PrivateBucket().GetObjectInfo(context.Background(), f.FilePath)

			fileIn := &firmwaremanage.OtaFirmwareFileReq{
				FileID:    f.FileID,
				Size:      storageInfo.Size,
				Signature: storageInfo.Md5,
			}
			s.FirmwareM.OtaFirmwareFileUpdate(context.Background(), fileIn)
		}

		if err != nil {
			fmt.Println(err)
			continue
		}
	}
}
