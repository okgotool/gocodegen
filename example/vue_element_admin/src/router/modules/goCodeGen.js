/** When your routing table is too long, you can split it into small modules **/

import Layout from '@/layout'

const goCodeGenRouter = {
  path: '/gocodegen',
  component: Layout,
  redirect: 'noRedirect',
  name: 'GoCodeGen',
  meta: {
    title: '生成代码',
    icon: 'guide'
  },
  children: [
    {
      path: 'sysrole',
      component: () => import('@/views/gen/sysrole'),
      name: 'SysRole',
      meta: { title: 'SysRole' }
    },
    {
      path: 'testcat',
      component: () => import('@/views/gen/testcat'),
      name: 'TestCat',
      meta: { title: 'TestCat' }
    }
  ]
}

export default goCodeGenRouter
