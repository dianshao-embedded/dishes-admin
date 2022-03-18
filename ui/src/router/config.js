import TabsView from '@/layouts/tabs/TabsView'
import BlankView from '@/layouts/BlankView'
import PageView from '@/layouts/PageView'

// 路由配置
const options = {
  routes: [
    {
      path: '/login',
      name: '登录页',
      component: () => import('@/pages/login')
    },
    {
      path: '*',
      name: '404',
      component: () => import('@/pages/exception/404'),
    },
    {
      path: '/403',
      name: '403',
      component: () => import('@/pages/exception/403'),
    },
    {
      path: '/',
      name: '首页',
      component: TabsView,
      redirect: '/login',
      children: [
        {
          path: 'dashboard',
          name: 'Dashboard',
          meta: {
            icon: 'dashboard'
          },
          component: BlankView,
          children: [
            {
              path: 'workplace',
              name: '工作台',
              meta: {
                page: {
                  closable: false
                }
              },
              component: () => import('@/pages/dashboard/workplace'),
            },
          ]
        },
        {
          path: 'update',
          name: '固件更新',
          meta: {
            icon: 'cloud-download',
          },
          component: PageView,
          children: [
            {
              path: 'product',
              name: '产品',
              component: () => import('@/pages/product/ProductForm.vue'),
            },
            {
              path: 'product/:id/firmware',
              name: '固件',
              meta: {
                highlight: '/update/product',
                invisible: true
              },
              component: () => import('@/pages/product/FirmwareForm.vue')
            },
            {
              path: 'product/:id/device',
              name: '设备',
              meta: {
                highlight: '/update/product',
                invisible: true
              },
              component: () => import('@/pages/product/DeviceForm.vue')
            },
            {
              path: 'product/:id/device/:id2/upgrade',
              name: '更新列表',
              meta: {
                highlight: '/update/product',
                invisible: true
              },
              component: () => import('@/pages/product/UpgradeForm.vue')
            },
          ],
        },
        {
          path: 'user',
          name: '用户管理',
          meta: {
            icon: 'table',
          },
          component: PageView,
          children: [
            {
              path: 'query',
              name: '用户列表',
              component: () => import('@/pages/user'),
            },
          ],
        },
      ]
    },
  ]
}

export default options
