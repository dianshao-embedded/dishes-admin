<template>
  <a-modal v-model="visible" :width="480" title="编辑示例信息" on-ok="handleOk">
    <a-form-model>
      <a-space>
      <a-form-model-item label="版本" prop="version" style="width: 200px">
          <a-input v-model="form.version"/>
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
      </a-space>
    </a-form-model>
    <a-space>
    <a-button @click="createVersion">确定固件版本及阶段</a-button>

    <a-upload
      name="file"
      :multiple="true"
      :action="url"
      :headers="headers"
      @change="handleChange"
    >
      <a-button> <a-icon type="upload" /> 选择上传文件 </a-button>
    </a-upload>
    </a-space>
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
  import { firmwareAdd } from '@/services/firmware'
  export default {
    name: 'FirmwareEdit',
    props: ["productID"],
    data() {
      return {
        headers: {
          authorization: 'authorization-text',
        },
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
          size: '',
          status: '',
          failure: 0,
          stage: '',
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
      createVersion() {
        this.url = ('/api/v1/firmwares/upload/' + this.form.product_id + '/' + this.form.stage + '/' + this.form.version)
        this.$message.success('选择成功');
      },
      handleChange(info) {
        if (info.file.status !== 'uploading') {
          console.log(info.file, info.fileList);
        }
        if (info.file.status === 'done') {
          this.$message.success(`${info.file.name} 上传成功`);
          this.form.name = info.file.name
          this.form.size =  info.file.size.toString()
          this.form.status = "pending"
        } else if (info.file.status === 'error') {
          this.$message.error(`${info.file.name} file upload failed.`);
          this.visible = false
          this.$emit('fetch-data')
        }
      },
      handleCancel() {
        this.visible = false
        this.$emit('fetch-data')
      },
      async handleOk() {
        try {
          await firmwareAdd(this.form)
          this.$message.success('添加成功')

          this.visible = false
          this.$emit('fetch-data')
        } catch (error) {
            console.log(error)
            this.$message.error('添加失败')
            this.visible = false
            this.$emit('fetch-data')
        }
      } 
    },
  }
</script>
