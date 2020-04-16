<template>
  <d2-container>
    <d2-crud
      ref="d2Crud"
      selection-row
      add-title="新增商品"
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
import Vue from 'vue'
import moment from 'moment'
import D2Crud from '@d2-projects/d2-crud'
import {
  QueryGoodsTypesByLimitOffset,
  DeleteGoodsType,
  DeleteGoodsTypes,
  AddGoodsType,
} from '@api/goods'

Vue.component('custom-input', {
  props: ['value'],
  template: `
    <input
      v-bind:value="value"
      v-on:input="$emit('input', $event.target.value)"
    >
  `,
})

export default {
  components: { D2Crud },
  data() {
    return {
      sels: [],
      all: [],
      pagination: {
        currentPage: 1,
        pageSize: 7,
        total: 0,
      },
      columns: [
        {
          title: 'ID',
          key: 'id',
          width: '50px',
        },
        {
          title: '商品名称',
          key: 'goods_name',
        },
        {
          title: '商品规格',
          key: 'goods_specs',
        },
        {
          title: '商品成本价',
          key: 'goods_unitprince',
        },
        {
          title: '商品销售价',
          key: 'goods_prince',
        },
        {
          title: '生产批号',
          key: 'goods_batch_number',
        },
        {
          title: '商品状态',
          key: 'goods_state',
          formatter: function(row) {
            if (row.goods_state == undefined) {
              return ''
            }
            if (row.goods_state == '1') {
              return '已下架'
            }
            if (row.goods_state == '2') {
              return '在售'
            }
          },
        },
        {
          title: '生产时间',
          key: 'goods_date',
          formatter: function(row) {
            if (row.goods_date == undefined) {
              return ''
            }
            return moment(row.goods_date).format('YYYY-MM-DD HH:mm:ss')
          },
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
        maxHeight: '70%',
        stripe: true,
      },
      rowHandle: {
        edit: {
          icon: 'el-icon-edit',
          size: 'small',
        },
        remove: {
          icon: 'el-icon-delete',
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
          goods_name: {
            title: '商品名称',
            value: '',
          },
          goods_specs: {
            title: '规格',
            value: '1',
            component: {
              name: 'el-select',
              options: [
                {
                  value: '1',
                  label: '盒',
                },
                {
                  value: '2',
                  label: '瓶',
                },
                {
                  value: '3',
                  label: '支',
                },
              ],
            },
          },
          goods_prince: {
            title: '销售价',
            value: '',
          },
          goods_unitprince: {
            title: '成本价',
            value: '',
          },
          goods_batch_number: {
            title: '生成批号',
            value: '',
          },
          goods_date: {
            title: '生产日期',
            value: '',
          },
          goods_image: {
            title: '商品图片',
            value: '',
            component: {
              name: 'custom-input',
            },
          },
        },
      })
    },
    // 单选删除
    handleRowRemove({ index, row }, done) {
      DeleteGoodsType(row.id)
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
            DeleteGoodsTypes({
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
      QueryGoodsTypesByLimitOffset(
        this.pagination.pageSize,
        this.pagination.currentPage
      )
        .then(async (res) => {
          this.data = res.goodsTypes
          this.pagination.total = res.count
        })
        .catch((err) => {
          Promise.reject(err)
        })
    },
    // 新增
    handleRowAdd(row, done) {
      this.formOptions.saveLoading = true
      AddGoodsType({
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
