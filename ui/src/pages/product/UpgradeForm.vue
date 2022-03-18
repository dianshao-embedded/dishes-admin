<template>
  <a-card :bordered="false">
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
      row-key="upgradeId"
      :loading="listLoading"
    >
    </standard-table>
  </a-card>
</template>

<script>
  import {
    upgradeQuery,
    upgradeDelete,
  } from '@/services/upgrade'
  import StandardTable from '@/components/table/StandardTable'
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
      title: '大小',
      dataIndex: 'size',
      scopedSlots: { customRender: 'size' },
    },
    {
      title: '状态',
      dataIndex: 'status',
      scopedSlots: { customRender: 'status' },
    },
    {
      title: '失败次数',
      dataIndex: 'failure',
      scopedSlots: { customRender: 'failure' },
    },
  ]

  export default {
    name: 'QueryList',
    components: { StandardTable },
    data() {
      return {
        advanced: false,
        columns: columns,
        selectedRows: [],
        queryForm: {
          pageNum: 1,
          pageSize: 50,
          loginId: '',
          device_id: '',
        },
        listLoading: false,
        list: [],
        total: 0,
        pwdVisible: false,
        newPwd: '',
        confirmLoading: false,
        confirmUpgradeId: '',
        createTime: null,
      }
    },
    mounted() {
      this.queryForm.device_id = this.$route.params.id2
      this.fetchData()
    },
    methods: {
      async onChange() {
        console.log('change')

      },
      async fetchData() {
        this.listLoading = true
        try {
          const { data } = await upgradeQuery({
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
        // const res = await getUpgradeById(record.id)
        this.$refs.upgradeEdit.showDialog(record)
      },
      createUpgrade() {
        this.$refs.upgradeEdit.showDialog()
      },
      async handleDelete(record) {
        await upgradeDelete(record.id)
        this.fetchData()
      },
      async handleConfirmReset() {
        this.confirmLoading = true
        try {
          let params = {
            newPwd: this.newPwd,
            upgradeId: this.confirmUpgradeId,
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
