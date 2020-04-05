package images
//
//import(
//	"errors"
//	docker "github.com/fsouza/go-dockerclient"
//)
//
///**
// *
// * @Author: zhangxiaohu
// * @File: mgr.go.go
// * @Version: 1.0.0
// * @Time: 2020/1/14
// */
//
//type DockerImageInfo struct{
//	Id string
//	path string
//}
//
//type Manager struct {}
//
//
//func (mgr * Manager) List()([]*DockerImageInfo , error)  {
//	// connect docker
//	// get docker images
//	var images[] *DockerImageInfo
//	client, err := docker.NewClientFromEnv()
//	if err != nil {
//		panic(err)
//	}
//	imgs, err := client.ListImages(docker.ListImagesOptions{All: false})
//	if err != nil {
//		panic(err)
//	}
//	for index, img := range imgs {
//		images[index].Id = img.ID
//	}
//	return images,errors.New("Unsupported method")
//}
//
//func (mgr * Manager) Set(image DockerImageInfo)(error){
//	client, err := docker.NewClientFromEnv()
//	if err != nil {
//		panic(err)
//	}
//	options := new (docker.ImportImageOptions)
//	options.Source = image.path
//	options.
//	client.ImportImage(*options)
//}