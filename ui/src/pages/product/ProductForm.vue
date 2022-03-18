<template>
  <a-card :bordered="false">
    <div>
      <a-space>
        <a-button slot="extra" @click="createProduct">新增产品</a-button>
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
      row-key="productId"
      :loading="listLoading"
    >
      <div slot="action" slot-scope="{ record }">
        <router-link :to="`/update/product/${record.id}/device`" >设备</router-link>
        <a-divider type="vertical" />
        <router-link :to="`/update/product/${record.id}/firmware`" >固件</router-link>
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

    <product-edit ref="productEdit" @fetch-data="fetchData" />
  </a-card>
</template>

<script>
  import {
    productQuery,
    productDelete,
  } from '@/services/product'
  import StandardTable from '@/components/table/StandardTable'
  import ProductEdit from './components/ProductEdit.vue'
  const columns = [
    {
      title: '名称',
      dataIndex: 'name',
      scopedSlots: { customRender: 'name' },
    },
    {
      title: '描述',
      dataIndex: 'description',
      scopedSlots: { customRender: 'description' },
    },
    {
      title: '操作系统',
      dataIndex: 'os',
      scopedSlots: { customRender: 'os' },
    },
    {
      title: '更新方式',
      dataIndex: 'update_method',
      scopedSlots: { customRender: 'update_method' },
    },
    {
      title: '固件格式',
      dataIndex: 'format',
      scopedSlots: { customRender: 'format' },
    },
    {
      title: '网络连接方式',
      dataIndex: 'network',
      scopedSlots: { customRender: 'network' },
    },
    {
      title: '终端类型',
      dataIndex: 'type',
      scopedSlots: { customRender: 'type' },
    },
    {
      title: '操作',
      dataIndex: 'productId',
      scopedSlots: { customRender: 'action' },
    },
  ]

  export default {
    name: 'QueryList',
    components: { StandardTable, ProductEdit },
    data() {
      return {
        advanced: false,
        columns: columns,
        selectedRows: [],
        queryForm: {
          pageNum: 1,
          pageSize: 50,
          loginId: '',
        },
        listLoading: false,
        list: [],
        total: 0,
        pwdVisible: false,
        newPwd: '',
        confirmLoading: false,
        confirmProductId: '',
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
          const { data } = await productQuery({
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
        // const res = await getProductById(record.id)
        this.$refs.productEdit.showDialog(record)
      },
      createProduct() {
        this.$refs.productEdit.showDialog()
      },
      async handleDelete(record) {
        await productDelete(record.id)
        this.fetchData()
      },
      async handleConfirmReset() {
        this.confirmLoading = true
        try {
          let params = {
            newPwd: this.newPwd,
            productId: this.confirmProductId,
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
