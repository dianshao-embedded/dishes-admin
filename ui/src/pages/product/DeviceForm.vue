<template>
  <a-card :bordered="false">
    <div>
      <a-space>
        <a-button slot="extra" @click="createDevice">新增设备</a-button>
      </a-space>
    </div>

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
      row-key="deviceId"
      :loading="listLoading"
    >
      <div slot="action" slot-scope="{ record }">
        <router-link :to="`/update/product/${record.product_id}/device/${record.id}/upgrade`" >更新记录</router-link>
        <a-divider type="vertical" />
        <a @click="handleShowDialog(record)">编辑</a>
        <a-divider type="vertical" />
        <a-popconfirm
          title="确定要删除吗"
          @confirm="handleDelete(record)"
        >
          <a>删除</a>
        </a-popconfirm>
      </div>
      <template slot="statusTitle">
        <a-icon type="info-circle" @click.native="onStatusTitleClick" />
      </template>
    </standard-table>

    <device-edit ref="deviceEdit" :productID="queryForm.productID" @fetch-data="fetchData" />
  </a-card>
</template>

<script>
  import {
    deviceQuery,
    deviceDelete,
  } from '@/services/device'
  import StandardTable from '@/components/table/StandardTable'
  import DeviceEdit from './components/DeviceEdit.vue'
  const columns = [
    {
      title: '名称',
      dataIndex: 'name',
      scopedSlots: { customRender: 'name' },
    },
    {
      title: '版本',
      dataIndex: 'version',
      scopedSlots: { customRender: 'version' },
    },
    {
      title: '阶段',
      dataIndex: 'stage',
      scopedSlots: { customRender: 'stage' },
    },
    {
      title: '序列号',
      dataIndex: 'sn',
      scopedSlots: { customRender: 'sn' },
    },
    {
      title: '操作',
      dataIndex: 'id',
      scopedSlots: { customRender: 'action' },
    },
  ]

  export default {
    name: 'QueryList',
    components: { StandardTable, DeviceEdit },
    data() {
      return {
        advanced: false,
        columns: columns,
        selectedRows: [],
        queryForm: {
          pageNum: 1,
          pageSize: 50,
          loginId: '',
          productID: '',
        },
        listLoading: false,
        list: [],
        total: 0,
        pwdVisible: false,
        newPwd: '',
        confirmLoading: false,
        confirmDeviceId: '',
        createTime: null,
      }
    },
    mounted() {
      this.queryForm.productID = this.$route.params.id
      this.fetchData()
    },
    methods: {
      async onChange() {
        console.log('change')

      },
      async fetchData() {
        this.listLoading = true
        try {
          const { data } = await deviceQuery({
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
      async handleShowDialog(record) {
        // const res = await getDeviceById(record.id)
        this.$refs.deviceEdit.showDialog(record)
      },
      createDevice() {
        this.$refs.deviceEdit.showDialog()
      },
      async handleDelete(record) {
        await deviceDelete(record.id)
        this.fetchData()
      },
      async handleConfirmReset() {
        this.confirmLoading = true
        try {
          let params = {
            newPwd: this.newPwd,
            deviceId: this.confirmDeviceId,
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
