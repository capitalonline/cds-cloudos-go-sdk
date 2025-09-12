package main

import (
	"testing"

	"github.com/capitalonline/cds-cloudos-go-sdk/services/ecs"
	"go.uber.org/mock/gomock"
)

func TestEcsExamples(t *testing.T) {
	// 创建GoMock控制器
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// 创建mock客户端
	mockClient := ecs.NewMockClient(ctrl)

	// 测试describeRegions函数
	t.Run("TestDescribeRegions", func(t *testing.T) {
		// 设置期望的调用
		mockClient.EXPECT().DescribeRegions().Return(&ecs.DescribeRegionsResult{}, nil)

		describeRegions(mockClient)
	})

	// 测试describeInstanceList函数
	t.Run("TestDescribeInstanceList", func(t *testing.T) {
		// 设置期望的调用
		mockClient.EXPECT().DescribeInstanceList(gomock.Any()).Return(&ecs.DescribeInstanceListResult{}, nil)

		// 执行测试
		describeInstanceList(mockClient)
	})

	// 测试operateInstance函数
	t.Run("TestOperateInstance", func(t *testing.T) {
		// 设置期望的调用
		mockClient.EXPECT().OperateInstance(gomock.Any()).Return(&ecs.OperateInstanceResult{}, nil)

		// 执行测试
		operateInstance(mockClient)
	})

	// 测试modifyInstanceName函数
	t.Run("TestModifyInstanceName", func(t *testing.T) {
		// 设置期望的调用
		mockClient.EXPECT().ModifyInstanceName(gomock.Any()).Return(&ecs.ModifyInstanceNameResult{}, nil)

		// 执行测试
		modifyInstanceName(mockClient)
	})

	// 测试describeInstance函数
	t.Run("TestDescribeInstance", func(t *testing.T) {
		// 设置期望的调用
		mockClient.EXPECT().DescribeInstance(gomock.Any()).Return(&ecs.DescribeInstanceResult{}, nil)

		// 执行测试
		describeInstance(mockClient)
	})

	// 测试describeTaskEvent函数
	t.Run("TestDescribeTaskEvent", func(t *testing.T) {
		// 设置期望的调用
		mockClient.EXPECT().DescribeTaskEvent(gomock.Any()).Return(&ecs.DescribeTaskEventResult{}, nil)

		// 执行测试
		describeTaskEvent(mockClient)
	})

	// 测试describeEcsFamilyInfo函数
	t.Run("TestDescribeEcsFamilyInfo", func(t *testing.T) {
		// 设置期望的调用
		mockClient.EXPECT().DescribeEcsFamilyInfo(gomock.Any()).Return(&ecs.DescribeEcsFamilyInfoResult{}, nil)

		// 执行测试
		describeEcsFamilyInfo(mockClient)
	})

	// 测试changeInstanceConfigure函数
	t.Run("TestChangeInstanceConfigure", func(t *testing.T) {
		// 设置期望的调用
		mockClient.EXPECT().ChangeInstanceConfigure(gomock.Any()).Return(&ecs.ChangeInstanceConfigureResult{}, nil)

		// 执行测试
		changeInstanceConfigure(mockClient)
	})

	// 测试extendDisk函数
	t.Run("TestExtendDisk", func(t *testing.T) {
		// 设置期望的调用
		mockClient.EXPECT().ExtendDisk(gomock.Any()).Return(&ecs.ExtendDiskResult{}, nil)

		// 执行测试
		extendDisk(mockClient)
	})
}
