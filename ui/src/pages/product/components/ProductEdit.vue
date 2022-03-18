<template>
  <a-modal v-model="visible" :width="620" title="编辑示例信息" on-ok="handleOk">
    <a-form-model
      ref="ruleForm"
      :model="form"
      :rules="rules"
      :label-col="labelCol"
      :wrapper-col="wrapperCol"
    >
      <a-form-model-item label="名称" prop="name">
        <a-input v-model="form.name"/>
      </a-form-model-item>
      <a-form-model-item label="描述" prop="description">
        <a-input v-model="form.description"/>
      </a-form-model-item>
      <a-form-model-item label="操作系统" prop="os">
        <a-select v-model="form.os">
          <a-select-option value="Linux">
            Linux
          </a-select-option>
          <a-select-option value="RTOS">
            RTOS
          </a-select-option>
        </a-select>
      </a-form-model-item>
      <a-form-model-item label="更新方式" prop="update_method">
        <a-select v-model="form.update_method">
          <a-select-option value="Auto">
            Auto
          </a-select-option>
          <a-select-option value="Manual">
            Manual
          </a-select-option>
        </a-select>
      </a-form-model-item>
      <a-form-model-item label="固件格式" prop="format">
        <a-input v-model="form.format"/>
      </a-form-model-item>
      <a-form-model-item label="网络连接方式" prop="network">
        <a-input v-model="form.network"/>
      </a-form-model-item>
      <a-form-model-item label="终端类型" prop="type">
        <a-input v-model="form.type"/>
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
  import { productAdd, productUpdate } from '@/services/product'

  export default {
    name: 'ProductEdit',
    data() {
      return {
        title: '',
        visible: false,
        labelCol: { span: 8 },
        wrapperCol: { span: 8 },
        other: '',
        roles: [],
        form: {
          name: '',
          description: '',
          os: '',
          format: '',
          network: '',
          type: '',
        },
        confirmLoading: false,
      }
    },
    methods: {
      showDialog(row) {
        this.visible = true
        if (row) {
          this.title = '修改示例'
          this.form = { ...this.form, ...row }
        } else {
          this.title = '添加示例'
        }
      },
      handleCancel() {
        this.visible = false
        this.$refs.ruleForm.resetFields()
      },
      handleOk() {
        this.$refs.ruleForm.validate(async (valid) => {
          if (valid) {
            this.confirmLoading = true
            try {
                if (this.form.id) {
                    await productUpdate(this.form.id, this.form)
                    this.$message.success('编辑成功')
                } else {
                    await productAdd(this.form)
                    this.$message.success('添加成功')
                }

                this.visible = false
                this.$emit('fetch-data')
            } catch (error) {
                console.log(error)
            } finally {
                this.confirmLoading = false
            }
          } else {
            console.log('error submit!!')
            return false
          }
        })
      },
    },
  }
</script>
