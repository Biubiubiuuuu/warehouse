// 菜单 侧边栏
export default [{
    path: '/index',
    title: '首页',
    icon: 'home'
  },
  {
    title: '角色管理',
    icon: 'folder-o',
    children: [{
        path: '/admin/queryAdmins',
        title: '管理员列表'
      },
      {
        path: '/admin/updateAdminPassword',
        title: '修改密码'
      }
    ]
  },
  {
    title: '商品管理',
    icon: 'folder-o',
    children: [{
        path: '/admin/goodsType/queryGoodsTypesByLimitOffset',
        title: '商品种类'
      },
      {
        path: '/admin/goodsStock/queryGoodsStocksByLimitOffset',
        title: '商品库存'
      }
    ]
  },
  {
    title: '订单管理',
    icon: 'folder-o',
    children: [{
      path: '/admin/order/queryOrderByLimitOffset',
      title: '用户订单'
    }]
  },
  {
    title: '用户管理',
    icon: 'folder-o',
    children: [{
      path: '/admin/users/queryUserByLimitOffset',
      title: '用户列表'
    }]
  }
]