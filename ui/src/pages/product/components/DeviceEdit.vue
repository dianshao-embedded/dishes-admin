<template>
  <a-modal v-model="visible" :width="480" title="编辑示例信息" on-ok="handleOk">
    <a-form-model>
      <a-form-model-item label="序列号" prop="sn" style="width: 200px">
          <a-input v-model="form.sn"/>
      </a-form-model-item>
      <a-form-model-item label="设备名" prop="name" style="width: 200px">
          <a-input v-model="form.name"/>
      </a-form-model-item>
      <a-form-model-item label="阶段" prop="stage" style="width: 200px">
        <a-select v-model="form.stage">
          <a-select-option value="dev">
            dev
          </a-select-option>
          <a-select-option value="test">
            test
          </a-select-option>
          <a-select-option value="release">
            release
          </a-select-option>
        </a-select>
      </a-form-model-item>
    </a-form-model>
    <template slot="footer">
      <a-button key="back" @click="handleCancel">取消</a-button>
      <a-button
        key="submit"
        type="primary"
        :loading="confirmLoading"
        @click="handleOk"
      >
        确定
      </a-button>
    </template>
  </a-modal>
</template>

<script>
  import { deviceAdd,  deviceUpdate } from '@/services/device'
  export default {
    name: 'DeviceEdit',
    props: ["productID"],
    data() {
      return {
        url: '',
        title: '',
        visible: false,
        labelCol: { span: 8 },
        wrapperCol: { span: 8 },
        other: '',
        roles: [],
        form: {
          product_id: '',
          name: '',
          version: '',
          stage: '',
          sn: '',
          status: '',
          failure: 0,
        },
        confirmLoading: false,
      }
    },
    methods: {
      showDialog(row) {
        this.visible = true
        this.form.product_id = (this.productID)
        if (row) {
          this.title = '修改示例'
          this.form = { ...this.form, ...row }
        } else {
          this.title = '添加示例'
        }
      },
      handleCancel() {
        this.visible = false
        this.$emit('fetch-data')
      },
      async handleOk() {
        console.log('OK')
          this.confirmLoading = true
          try {
              if (this.form.id) {
                  await deviceUpdate(this.form.id, this.form)
                  this.$message.success('编辑成功')
              } else {
                  await deviceAdd(this.form)
                  this.$message.success('添加成功')
              }

              this.visible = false
              this.$emit('fetch-data')
          } catch (error) {
              console.log(error)
          } finally {
              this.confirmLoading = false
          }
      },
    },
  }
</script>
