<template>
  <d2-container>
    <el-form
      :model="ruleForm"
      :rules="rules"
      show-icon
      ref="ruleForm"
      label-width="100px"
      class="demo-ruleForm"
    >
      <el-form-item label="旧密码" prop="oldpass">
        <el-input
          type="password"
          v-model="ruleForm.oldpass"
          autocomplete="off"
        ></el-input>
      </el-form-item>
      <el-form-item label="新密码" prop="newpass">
        <el-input
          type="password"
          v-model="ruleForm.newpass"
          autocomplete="off"
        ></el-input>
      </el-form-item>
      <el-form-item label="确认新密码" prop="checkPass">
        <el-input
          type="password"
          v-model="ruleForm.checkPass"
          autocomplete="off"
        ></el-input>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="submitForm('ruleForm')"
          >提交</el-button
        >
        <el-button @click="resetForm('ruleForm')">重置</el-button>
      </el-form-item>
    </el-form>
  </d2-container>
</template>

<script>
import router from '@/router'
import { UpdateAdminPassword } from '@api/admin'
export default {
  data() {
    var validateOldPass = (rule, value, callback) => {
      if (value === '') {
        callback(new Error('请输入旧密码'))
      } else {
        callback()
      }
    }
    var validateNewPass = (rule, value, callback) => {
      if (value === '') {
        callback(new Error('请输入新密码'))
      } else {
        if (this.ruleForm.checkPass !== '') {
          this.$refs.ruleForm.validateField('checkPass')
        }
        callback()
      }
    }
    var validateCheckPass = (rule, value, callback) => {
      if (value === '') {
        callback(new Error('请再次输入密码'))
      } else if (value !== this.ruleForm.newpass) {
        callback(new Error('两次输入密码不一致!'))
      } else {
        callback()
      }
    }
    return {
      ruleForm: {
        oldpass: '',
        newpass: '',
        checkPass: '',
      },
      rules: {
        oldpass: [{ validator: validateOldPass, trigger: 'blur' }],
        newpass: [{ validator: validateNewPass, trigger: 'blur' }],
        checkPass: [{ validator: validateCheckPass, trigger: 'blur' }],
      },
    }
  },
  methods: {
    submitForm(formName) {
      this.$refs[formName].validate(async (valid) => {
        if (valid) {
          UpdateAdminPassword({
            oldPassword: this.ruleForm.oldpass,
            newPassword: this.ruleForm.newpass,
          }).then((res) => {
            this.$message({
              message: '修改成功',
              type: 'success',
            })
            // 跳转路由
            router.push({
              name: 'login',
            })
          })
        } else {
          this.$message({
            message: '修改失败',
            type: 'warning',
          })
          return false
        }
      })
    },
    resetForm(formName) {
      this.$refs[formName].resetFields()
    },
  },
}
</script>
