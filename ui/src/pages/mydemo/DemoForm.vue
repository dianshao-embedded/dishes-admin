<template>
  <a-card title="示例" :bordered="false">
    <a-button slot="extra" @click="createDemo">新建</a-button>

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
      row-key="demoId"
      :loading="listLoading"
    >
      <div slot="action" slot-scope="{ record }">
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

    <demo-edit ref="demoEdit" @fetch-data="fetchData" />
  </a-card>
</template>

<script>
  import {
    demoQuery,
    demoDelete,
  } from '@/services/demo'
  import StandardTable from '@/components/table/StandardTable'
  import DemoEdit from './componets/DemoEdit.vue'
  const columns = [
    {
      title: '任务编号',
      dataIndex: 'code',
      scopedSlots: { customRender: 'code' },
    },
    {
      title: '任务名称',
      dataIndex: 'name',
      scopedSlots: { customRender: 'name' },
    },
    {
      title: '任务备注',
      dataIndex: 'memo',
      scopedSlots: { customRender: 'memo' },
    },
    {
      title: '整数',
      dataIndex: 'number',
      scopedSlots: { customRender: 'number' },
    },
    {
      title: '无符号整数',
      dataIndex: 'unsigned_number',
      scopedSlots: { customRender: 'unsigned_number' },
    },
    {
      title: '操作',
      dataIndex: 'demoId',
      scopedSlots: { customRender: 'action' },
    },
  ]

  export default {
    name: 'QueryList',
    components: { StandardTable, DemoEdit },
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
        confirmDemoId: '',
        createTime: null,
      }
    },
    mounted() {
      this.fetchData()
    },
    methods: {
      async onChange() {
        console.log('change')

      },
      async fetchData() {
        this.listLoading = true
        try {
          const { data } = await demoQuery({
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
        // const res = await getDemoById(record.id)
        this.$refs.demoEdit.showDialog(record)
      },
      createDemo() {
        this.$refs.demoEdit.showDialog()
      },
      async handleDelete(record) {
        await demoDelete(record.id)
        this.fetchData()
      },
      async handleConfirmReset() {
        this.confirmLoading = true
        try {
          let params = {
            newPwd: this.newPwd,
            demoId: this.confirmDemoId,
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
