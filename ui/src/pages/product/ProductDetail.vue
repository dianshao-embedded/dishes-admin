<template>
  <a-card :bordered="false">
    <detail-list title="基础信息">
      <detail-list-item term="名称">{{product.name}}</detail-list-item>
      <detail-list-item term="描述">{{product.description}}</detail-list-item>
      <detail-list-item term="操作系统">{{product.os}}</detail-list-item>
      <detail-list-item term="更新方式">{{product.update_method}}</detail-list-item>
      <detail-list-item term="固件格式">{{product.format}}</detail-list-item>
      <detail-list-item term="网络连接方式">{{product.network}}</detail-list-item>
      <detail-list-item term="终端类型">{{product.type}}</detail-list-item>
    </detail-list>
    <br>
    <div>
      <a-space>
        固件列表
        <a-button slot="extra" @click="createFirmware">手动新增固件</a-button>
      </a-space>
    </div>
    <br>
    <div>
    <a-table
      row-key="id"
      style="margin-bottom: 24px"
      :columns="columns"
      :dataSource="product.firmware_list"
      :pagination="false"
    >
      <div slot="action" slot-scope="{ record }">
        <a-popconfirm
          title="确定要删除吗"
          @confirm="handleDelete(record)"
        >
          <a>删除</a>
        </a-popconfirm>
      </div>
    </a-table>
    </div>
  </a-card>
</template>

<script>
  import {
    productGet,
  } from '@/services/product'
  import {
    firmwareDelete,
  } from '@/services/firmware' 
  import DetailList from '@/components/tool/DetailList'

  const DetailListItem = DetailList.Item

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
      title: '操作',
      dataIndex: 'id',
      scopedSlots: { customRender: 'action' },
    },
  ]

  export default {
    name: 'ProductDetail',
    components: { DetailListItem, DetailList },
    data() {
      return {
        columns: columns,
        product: {
          name: '',
          description: '',
          os: '',
          update_method: '',
          format: '',
          network: '',
          type: '',
          firmware_list: [],
        },
        queryForm: {
          pageNum: 1,
          pageSize: 50,
          loginId: '',
        },
        listLoading: false,
      }
    },
    mounted() {
      this.getData(this.$route.params.id)
    },
    methods: {
      async getData(id) {
        try {
          let req = await productGet(id)
          console.log(req)
          this.product = req.data
        } catch (error) {
          console.log(error)
        } finally {
          this.listLoading = false
        }
      },
      async handleDelete(record) {
        await firmwareDelete(record.id)
        this.getData()
      },
    }
  }
</script>

<style lang="less" scoped>
  .title {
    color: @title-color;
    font-size: 16px;
    font-weight: 500;
    margin-bottom: 16px;
  }
</style>