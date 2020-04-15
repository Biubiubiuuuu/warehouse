<template>
  <d2-container>
    <d2-crud
      ref="d2Crud"
      selection-row
      add-title="新增管理员"
      :columns="columns"
      :data="data"
      :options="options"
      :rowHandle="rowHandle"
      :pagination="pagination"
      :form-options="formOptions"
      @selection-change="selsChange"
      @row-remove="handleRowRemove"
      @pagination-current-change="paginationCurrentChange"
      @row-add="handleRowAdd"
      @dialog-cancel="handleDialogCancel"
    >
      <el-button slot="header" style="margin-bottom: 5px" @click="addRow"
        >新增</el-button
      >
      <el-button
        slot="header"
        style="margin-bottom: 5px"
        @click="handleDelete(sels.map((i) => i.id))"
        >批量删除</el-button
      >
    </d2-crud>
  </d2-container>
</template>

<script>
import moment from 'moment'
import D2Crud from '@d2-projects/d2-crud'
import {
  AdminsByLimitOffset,
  AddAdmin,
  DeleteAdmin,
  DeleteBatchAdmin,
} from '@api/admin'
export default {
  components: { D2Crud },
  data() {
    return {
      sels: [],
      all: [],
      pagination: {
        currentPage: 1,
        pageSize: 15,
        total: 0,
      },
      columns: [
        {
          title: 'ID',
          key: 'id',
          width: '50px',
        },
        {
          title: '用户名',
          key: 'username',
        },
        {
          title: '权限（Y/N）',
          key: 'administrator',
        },
        {
          title: 'ip',
          key: 'ip',
        },
        {
          title: '创建时间',
          key: 'created_at',
          formatter: function(row) {
            if (row.created_at == undefined) {
              return ''
            }
            return moment(row.created_at).format('YYYY-MM-DD HH:mm:ss')
          },
        },
      ],
      data: [],
      options: {
        maxHeight: '80%',
      },
      rowHandle: {
        remove: {
          icon: 'el-icon-delete',
          text: '删除',
          size: 'small',
        },
      },
      formOptions: {
        labelWidth: '80px',
        labelPosition: 'left',
        saveLoading: false,
      },
    }
  },
  mounted() {
    this.fetchData()
  },
  methods: {
    addRow() {
      this.$refs.d2Crud.showDialog({
        mode: 'add',
        template: {
          username: {
            title: '用户名',
            value: '',
          },
          password: {
            title: '密码',
            value: '',
          },
          administrator: {
            title: '超级管理员',
            value: 'N',
            component: {
              name: 'el-select',
              options: [
                {
                  value: 'Y',
                  label: '是',
                },
                {
                  value: 'N',
                  label: '否',
                },
              ],
            },
          },
        },
      })
    },
    // 单选删除
    handleRowRemove({ index, row }, done) {
      DeleteAdmin(row.id)
        .then((res) => {
          this.$message({
            message: '删除成功',
            type: 'success',
          })
          this.fetchData()
          done()
        })
        .catch((err) => {
          Promise.reject(err)
        })
    },
    selsChange(sels) {
      this.sels = sels
    },
    // 多选删除
    handleDelete(idArray) {
      if (idArray.length > 0) {
        this.$confirm('确认删除吗？', '提示', {
          confirmButtonText: '确认',
          cancelButtonText: '取消',
          type: 'warning',
        })
          .then(() => {
            DeleteBatchAdmin({
              ids: idArray,
            })
              .then((res) => {
                this.$message({
                  message: '删除成功',
                  type: 'success',
                })
                this.fetchData()
              })
              .catch((err) => {
                Promise.reject(err)
              })
          })
          .catch(() => {
            this.$message({
              type: 'info',
              message: '已取消批量删除',
            })
          })
      } else {
        this.$message({
          message: '请选择要删除的数据',
          type: 'error',
        })
      }
    },
    // 分页
    paginationCurrentChange(currentPage) {
      this.pagination.currentPage = currentPage
      this.fetchData()
    },
    // 获取数据
    fetchData() {
      AdminsByLimitOffset(this.pagination.pageSize, this.pagination.currentPage)
        .then(async (res) => {
          this.data = res.users
          this.pagination.total = res.count
        })
        .catch((err) => {
          Promise.reject(err)
        })
    },
    // 新增
    handleRowAdd(row, done) {
      this.formOptions.saveLoading = true
      AddAdmin({
        username: row.username,
        password: row.password,
        administrator: row.administrator,
      })
        .then((res) => {
          this.$refs.d2Crud.closeDialog()
          this.$message({
            message: '新增成功',
            type: 'success',
          })
          this.fetchData()
        })
        .catch((err) => {
          Promise.reject(err)
          done(false)
        })
      this.formOptions.saveLoading = false
    },
    handleDialogCancel(done) {
      this.$message({
        message: '取消保存',
        type: 'warning',
      })
      done()
    },
  },
}
</script>
