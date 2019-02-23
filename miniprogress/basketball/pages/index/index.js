//index.js
//获取应用实例
var app = getApp()
var latitude, longitude
//经纬度参数
Page({
  data: {
    latitude: 22.543099,
    longitude: 114.057868,
    lockhidden: true,
    markers: [{
      //摩拜单车的位置
      iconPath: "../../images/bike@red.png",
        id: 1,
        latitude: 22.540069,
        longitude: 114.059878,
        width: 40,
        height: 40,
        callout: {
          content:"爱你呀！1"
        }
      }, {
        iconPath: "../../images/bike@white.png",
        id: 2,
        latitude: 22.544069,
        longitude: 114.050878,
        width: 40,
        height: 40,
        callout: {
          content: "爱你呀！2"
        }
      }, {
        iconPath: "../../images/bike@red.png",
        id: 3,
        latitude: 22.546069,
        longitude: 114.056878,
        width: 40,
        height: 40,
        callout: {
          content: "爱你呀！3"
        }
      },{
        iconPath: "../../images/tome.png",
        id: 4,
        latitude: 22.546099,
        longitude: 114.057868,
        width: 70,
        height: 30
      },{
        iconPath: "../../images/hongbao.png",
        id: 5,
        latitude: 22.543069,
        longitude: 114.056878,
        width: 40,
        height: 40
      },{
      // 我的位置
        iconPath: "../../images/lable@icon.png",
        id: 0,
        latitude: 22.543099,
        longitude: 114.057868,
        width: 40,
        height: 40
     }],
     //控件
    controls: [{
      // 我的位置控件
      id: 0,
      iconPath: "../../images/imgs_main_location@2x.png",
      position: {
        left: 10,
        top: 590,
        width: 50,
        height: 50
      },
      clickable: true
    }, {
      // 红包控件
      id: 1,
      iconPath: "../../images/hongbao.png",
      position: {
        left: 350,
        top: 525,
        width: 50,
        height: 50
      },
      clickable: true
      }, {
        // 充值控件
        id: 2,
        iconPath: "../../images/chongzhi.png",
        position: {
          left: 350,
          top: 590,
          width: 50,
          height: 50
        },
        clickable: true
    }, {
      //二维码控件
      id: 3,
      iconPath: "../../images/lock.png",
      position: {
        left: 150,
        top: 580,
        width: 120,
        height: 60
      },
      clickable: true
    }]
  },
  onLoad: function () {
    var that = this
    // 获取当前经纬度
    wx.getLocation({
      type: 'wgs84',
      success: function (res) {
        latitude = res.latitude
        longitude = res.longitude
        console.log(latitude)
        console.log(longitude)
      }
    })

    //获取附近的篮球场地list
    //  不在以下 request 合法域名列表中  ==> 假设成功吧
    //  测试阶段，在微信官方IDE中，需要设置才能请求 未在小程序中备案的域名。。。。（设置===》项目====》项目设置 ===》 不校验域名 ）
    // wx.request({
    //   url: 'https://www.easy-mock.com/mock/5c6f5c85f5c55f016b1d7617/nice/nice',
    //   method : "GET",
    //   header : {
    //     'Content-Type': 'application/json',
    //     'Accept':'application/json, */*'
    //   },
    //   success: function(res){
    //     console.log(res.data);

    //   }
    // })    




    //调用应用实例的方法获取全局数据
    // app.getUserInfo(function(userInfo){
    //   //更新数据
    //   that.setData({
    //     latitude: latitude,
    //     longitude: longitude,
    //     userInfo:userInfo
    //   })
    // })
  },
  // 控件处理程序
  controltap(e) {
    // wx.showToast({
    //   title: '到这里 or 查看详情',
    // })

    var id = e.markerId

    wx.showToast({
      title: id.toString(),
    })
    this.setData({
      lockhidden: false
    });
  },

  // 控件处理程序
  controltapre(e) {
    // wx.showToast({
    //   title: '视野发生变化',
    // })
    // 获取中心点，请求服务器 获取球场list
    console.log("视野发生变化")
  },


  // 控件处理程序
  controltapcall(e) {
    wx.showToast({
      title: 'call',
    })

    // 小气泡处理 1、查看详情 2、到这里去 
    console.assert("nice")
  },


  //事件处理函数
  bindViewTap: function () {
    wx.navigateTo({
      url: '../logs/logs'
    })
  },
  // 扫码开锁弹出层显示隐藏
  confirm: function () {
    this.setData({
      lockhidden: true
    });
  }
})
