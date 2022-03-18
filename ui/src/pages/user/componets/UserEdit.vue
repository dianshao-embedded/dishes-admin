<template>
  <a-modal v-model="visible" :width="620" title="编辑用户信息" on-ok="handleOk">
    <a-form-model
      ref="ruleForm"
      :model="form"
      :rules="rules"
      :label-col="labelCol"
      :wrapper-col="wrapperCol"
    >
      <a-form-model-item label="用户名" prop="user_name">
        <a-input v-model="form.user_name" placeholder="请输入用户名" />
      </a-form-model-item>
      <a-form-model-item label="用户密码" prop="password">
        <a-input v-model="form.password" placeholder="请输入用户密码" />
      </a-form-model-item>
      <a-form-model-item label="用户昵称" prop="real_name">
        <a-input v-model="form.real_name" placeholder="请输入用户昵称" />
      </a-form-model-item>
      <a-form-model-item label="邮箱" prop="email">
        <a-input v-model="form.email" placeholder="请输入邮箱" />
      </a-form-model-item>
      <a-form-model-item label="手机号" prop="phone">
        <a-input v-model="form.phone" placeholder="请输入正确手机号" />
      </a-form-model-item>
      <a-form-model-item label="用户角色" prop="user_roles">
        <a-select v-model="form.user_roles">
          <a-select-option v-for="role in roles" :key="role.id">
            {{ role.name }}
          </a-select-option>
        </a-select>
      </a-form-model-item>
      <a-form-model-item label="用户状态" prop="status">
        <a-radio-group v-model="form.status">
          <a-radio :value="1">正常</a-radio>
          <a-radio :value="2">禁用</a-radio>
        </a-radio-group>
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
  import { userAdd, userUpdate, getRoles } from '@/services/user'
  import {
    isPhone,
    isEmail,
    isIP,
    isBlank,
  } from '@/utils/validate'

  export default {
    name: 'UserEdit',
    data() {
      const validatePhone = (rule, value, callback) => {
        if (isBlank(value)) {
          callback()
        } else if (!isPhone(value)) {
          callback(new Error('请输入正确的手机号'))
        } else {
          callback()
        }
      }
      const validateEmail = (rule, value, callback) => {
        if (isBlank(value)) {
          callback()
        } else if (!isEmail(value)) {
          callback(new Error('请输入正确的邮箱地址'))
        } else {
          callback()
        }
      }
      const validateLastLoginIp = (rule, value, callback) => {
        if (isBlank(value)) {
          callback()
        } else if (!isIP(value)) {
          callback(new Error('请输入正确的IP地址'))
        } else {
          callback()
        }
      }
      return {
        title: '',
        visible: false,
        labelCol: { span: 4 },
        wrapperCol: { span: 17 },
        other: '',
        roles: [],
        form: {
          user_name: '',
          password: '',
          real_name: '',
          email: '',
          phone: '',
          status: '',
          user_roles: '',
        },
        rules: {
          userName: [
            {
              required: true,
              message: '请输入用户名',
              trigger: 'blur',
            },
          ],
          mobile: [
            {
              required: true,
              validator: validatePhone,
              trigger: 'blur',
            },
          ],
          email: [
            {
              validator: validateEmail,
              trigger: 'blur',
            },
          ],
          lastLoginIp: [
            {
              validator: validateLastLoginIp,
              trigger: 'blur',
            },
          ],
        },
        confirmLoading: false,
      }
    },
    created() {
      this.getRoles()
    },
    methods: {
      async getRoles() {
        const res = await getRoles()
        this.roles = res.data.list
      },
      showDialog(row) {
        this.visible = true
        if (row) {
          this.title = '修改用户'
          this.form = { ...this.form, ...row }
          this.form.user_roles = row.roles[0].id
        } else {
          this.title = '添加用户'
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
            this.form.user_roles = [{ role_id: this.form.user_roles }]
            try {
              if (this.form.id) {
                await userUpdate(this.form.id, this.form)
                this.$message.success('编辑成功')
              } else {
                await userAdd(this.form)
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
