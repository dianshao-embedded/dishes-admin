<template>
  <common-layout>
    <div class="top" style="margin-top: 60px">
      <div class="header">
        <span class="title">{{systemName}}</span>
      </div>
    </div>

    <div class="login" style="margin-top: 120px">
      <a-form-model ref="ruleForm" :model="form" :rules="rules">
        <a-alert
          v-show="error"
          type="error"
          :closable="true"
          :message="error"
          show-icon
          style="margin-bottom: 24px"
        />
      </a-form-model>
      <a-form-model-item ref="name" prop="user_name">
        <a-input
          v-model="form.user_name"
          autocomplete="autocomplete"
          size="large"
        >
          <a-icon slot="prefix" type="user" />
        </a-input>
      </a-form-model-item>
      <a-form-model-item ref="password" prop="password">
        <a-input
          v-model="form.password"
          size="large"
          autocomplete="autocomplete"
          type="password"
        >
          <a-icon slot="prefix" type="lock" />
        </a-input>
      </a-form-model-item>
      <a-form-model-item>
        <a-button
          :loading="logging"
          style="width: 100%; margin-top: 24px"
          size="large"
          html-type="submit"
          type="primary"
          @click="onSubmit"
        >
          登录
        </a-button>
      </a-form-model-item>
    </div>
  </common-layout>
</template>

<script>
  import CommonLayout from '@/layouts/CommonLayout'
  import { login, getUser } from '@/services/user'
  import { setAuthorization } from '@/utils/request'
  import { mapMutations } from 'vuex'
  export default {
    name: 'Login',
    components: { CommonLayout},
    data() {
      return {
        identifyCode: 'ab1v',
        identifyCodes: {},
        logging: false,
        error: '',
        form: {
          user_name: '',
          password: '',
        },
        rules: {
          user_name: [
            { required: true, message: '请输入用户名', trigger: 'blur' },
          ],
          password: [
            { required: true, message: '请输入密码', trigger: 'change' },
          ],
        },
      }
    },
    computed: {
      systemName() {
        return this.$store.state.setting.systemName
      },
    },
    methods: {
      ...mapMutations('account', ['setUser', 'setPermissions', 'setRoles']),
      onSubmit() {
        this.$refs.ruleForm.validate((valid) => {
          if (valid) {
            this.logging = true
            // this.form.password = sha1(this.form.password)
            login(this.form).then(this.afterLogin)
            setTimeout(() => {
              this.logging = false
            }, 2000)
          } else {
            console.log('error submit!!')
            return false
          }
        })
      },

      resetForm() {
        this.$refs.ruleForm.resetFields()
      },
      async afterLogin(res) {
        this.logging = false
        const loginRes = res.data
        if (res.status === 200) {
          setAuthorization({
            token: loginRes?.access_token,
            expireAt: new Date(loginRes?.expiresAt),
          })
          const userRes = await getUser().then((res) => {
            console.log(res, 'rrrrrr')
          })
          console.log(userRes, loginRes, 'loginRes--------->')
          const { user, permissions, roles } = loginRes
          this.setUser(user)
          this.setPermissions(permissions)
          this.setRoles(roles)
          this.$router.push({ path: '/dashboard/workplace' }, () => {})
          this.$message.success(loginRes.message, 3)
        } else {
          this.error = loginRes.message
        }
      },
    },
  }
</script>

<style lang="less" scoped>
  .common-layout {
    .alpha {
      background-color: rgba(25, 25, 112, 0.3);
    }
    .top {
      text-align: center;
      .header {
        height: 44px;
        line-height: 44px;
        a {
          text-decoration: none;
        }
        .logo {
          height: 44px;
          vertical-align: top;
          margin-right: 16px;
        }
        .title {
          font-size: 33px;
          color: @title-color;
          font-family: 'Myriad Pro', 'Helvetica Neue', Arial, Helvetica,
            sans-serif;
          font-weight: 600;
          position: relative;
          top: 2px;
        }
      }
      .desc {
        font-size: 14px;
        color: @text-color-second;
        margin-top: 12px;
        margin-bottom: 40px;
      }
    }
    .login {
      width: 368px;
      margin: 0 auto;
      @media screen and (max-width: 576px) {
        width: 95%;
      }
      @media screen and (max-width: 320px) {
        .captcha-button {
          font-size: 14px;
        }
      }
      .icon {
        font-size: 24px;
        color: @text-color-second;
        margin-left: 16px;
        vertical-align: middle;
        cursor: pointer;
        transition: color 0.3s;

        &:hover {
          color: @primary-color;
        }
      }
    }
  }
</style>
