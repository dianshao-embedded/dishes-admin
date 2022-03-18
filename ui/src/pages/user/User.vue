<template>
  <a-card title="用户列表" :bordered="false">
    <a-button slot="extra" @click="createUser">新建</a-button>

    <standard-table
      :columns="columns"
      :data-source="list"
      :pagination="{
        current: queryForm.pageNum,
        pageSize: queryForm.pageSize,
        total: total,
        showSizeChanger: true,
        showLessItems: true,
        showQuickJumper: true,
        onChange: onPageChange,
        onShowSizeChange: onSizeChange,
        showTotal: (total, range) =>
          `第 ${range[0]}-${range[1]} 条，总计 ${total} 条`,
      }"
      row-key="userId"
      :loading="listLoading"
    >
      <template slot="status" slot-scope="{ text, record }">
        <span>
          <a-switch
            v-if="text === 1"
            default-checked
            @change="onChange(record)"
          />
          <a-switch v-if="text === 2" @change="onChange(record)" />
        </span>
      </template>

      <template slot="email" slot-scope="{ text }">
        <span>{{ text || '--' }}</span>
      </template>

      <div slot="action" slot-scope="{ record }">
        <a @click="handleShowDialog(record)">编辑</a>
        <a-divider type="vertical" />
        <a-popconfirm
          title="确定要删除该用户吗"
          @confirm="handleDelete(record)"
        >
          <a>删除</a>
        </a-popconfirm>
      </div>
      <template slot="statusTitle">
        <a-icon type="info-circle" @click.native="onStatusTitleClick" />
      </template>
    </standard-table>

    <user-edit ref="userEdit" @fetch-data="fetchData" />
  </a-card>
</template>

<script>
  import {
    userQuery,
    userDelete,
    userDisable,
    userEnable,
  } from '@/services/user'
  import StandardTable from '@/components/table/StandardTable'
  import UserEdit from './componets/UserEdit.vue'
  const columns = [
    {
      title: '用户名',
      dataIndex: 'user_name',
    },
    {
      title: '昵称',
      dataIndex: 'real_name',
      scopedSlots: { customRender: 'real_name' },
    },
    {
      title: '电话',
      dataIndex: 'phone',
      scopedSlots: { customRender: 'phone' },
    },
    {
      title: '邮箱',
      dataIndex: 'email',
      scopedSlots: { customRender: 'email' },
    },
    {
      title: '状态',
      dataIndex: 'status',
      scopedSlots: { customRender: 'status' },
    },
    {
      title: '操作',
      dataIndex: 'userId',
      scopedSlots: { customRender: 'action' },
    },
  ]

  export default {
    name: 'QueryList',
    components: { StandardTable, UserEdit },
    data() {
      return {
        advanced: false,
        columns: columns,
        selectedRows: [],
        queryForm: {
          pageNum: 1,
          pageSize: 10,
          loginId: '',
        },
        listLoading: false,
        list: [],
        total: 0,
        pwdVisible: false,
        newPwd: '',
        confirmLoading: false,
        confirmUserId: '',
        createTime: null,
      }
    },
    mounted() {
      this.fetchData()
    },
    methods: {
      async onChange(record) {
        if (record.status === 1) {
          await userDisable(record.id)
        } else {
          await userEnable(record.id)
        }
      },
      async fetchData() {
        this.listLoading = true
        try {
          const { data } = await userQuery({
            ...this.queryForm,
            ...this.extraQueryForm,
          })
          this.list = data?.list || []
          this.totalCount = data?.totalCount || 0
        } catch (error) {
          console.log(error)
        } finally {
          this.listLoading = false
        }
      },
      onSizeChange(current, size) {
        this.queryForm.pageNum = 1
        this.queryForm.pageSize = size
        this.fetchData()
      },
      onPageChange(page, pageSize) {
        this.queryForm.pageNum = page
        this.queryForm.pageSize = pageSize
        this.fetchData()
      },
      resetPwd(text) {
        // let params = {
        //   newPwd,
        // }
        // resetPwd({userResetPwdAdminBo: })
        this.pwdVisible = true
        this.confirmUserId = text
      },
      async handleShowDialog(record) {
        // const res = await getUserById(record.id)
        this.$refs.userEdit.showDialog(record)
      },
      createUser() {
        this.$refs.userEdit.showDialog()
      },
      handleDelete(record) {
        console.log('delete user')
        userDelete(record.id)
        this.fetchData()
      },
      async handleConfirmReset() {
        this.confirmLoading = true
        try {
          let params = {
            newPwd: this.newPwd,
            userId: this.confirmUserId,
          }
          this.resetPwd(params)
          this.$message.success('重置密码成功')
          this.pwdVisible = false
          this.newPwd = ''
        } catch (error) {
          console.log(error)
        } finally {
          this.confirmLoading = false
        }
      },
      resetFields() {
        this.queryForm.loginId = ''
        this.createTime = null
        this.fetchData()
      },
      toggleAdvanced() {
        this.advanced = !this.advanced
      },
    },
  }
</script>

<style lang="less" scoped>
  .search {
    margin-bottom: 54px;
  }
  .fold {
    width: calc(100% - 216px);
    display: inline-block;
  }
  .operator {
    margin-bottom: 18px;
  }
  @media screen and (max-width: 900px) {
    .fold {
      width: 100%;
    }
  }
</style>
