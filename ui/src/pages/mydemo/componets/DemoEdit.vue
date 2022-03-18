<template>
  <a-modal v-model="visible" :width="620" title="编辑示例信息" on-ok="handleOk">
    <a-form-model
      ref="ruleForm"
      :model="form"
      :rules="rules"
      :label-col="labelCol"
      :wrapper-col="wrapperCol"
    >
      <a-form-model-item label="任务编号" prop="code">
        <a-input v-model="form.code"/>
      </a-form-model-item>
      <a-form-model-item label="任务名称" prop="name">
        <a-input v-model="form.name"/>
      </a-form-model-item>
      <a-form-model-item label="任务备注" prop="memo">
        <a-input v-model="form.memo"/>
      </a-form-model-item>
      <a-form-model-item label="整数" prop="number">
        <a-input-number v-model="form.number"/>
      </a-form-model-item>
      <a-form-model-item label="无符号整数" prop="unsigned_number">
        <a-input-number v-model="form.unsigned_number"/>
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
  import { demoAdd, demoUpdate } from '@/services/demo'

  export default {
    name: 'DemoEdit',
    data() {
      return {
        title: '',
        visible: false,
        labelCol: { span: 4 },
        wrapperCol: { span: 17 },
        other: '',
        roles: [],
        form: {
          code: '',
          name: '',
          memo: '',
          number: undefined,
          unsigned_number: undefined,
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
                    await demoUpdate(this.form.id, this.form)
                    this.$message.success('编辑成功')
                } else {
                    await demoAdd(this.form)
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
