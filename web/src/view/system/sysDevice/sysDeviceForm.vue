
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="设备id:" prop="device_id">
    <el-input v-model="formData.device_id" :clearable="true" placeholder="请输入设备id" />
</el-form-item>
        <el-form-item label="设备通信id:" prop="subscribe_id">
    <el-input v-model="formData.subscribe_id" :clearable="true" placeholder="请输入设备通信id" />
</el-form-item>
        <el-form-item label="设备token:" prop="token">
    <el-input v-model="formData.token" :clearable="true" placeholder="请输入设备token" />
</el-form-item>
        <el-form-item label="注册地ip:" prop="register_ip">
    <el-input v-model="formData.register_ip" :clearable="true" placeholder="请输入注册地ip" />
</el-form-item>
        <el-form-item label="软件版本号:" prop="version">
    <el-input v-model="formData.version" :clearable="true" placeholder="请输入软件版本号" />
</el-form-item>
        <el-form-item label="硬件序列号:" prop="mac">
    <el-input v-model="formData.mac" :clearable="true" placeholder="请输入硬件序列号" />
</el-form-item>
        <el-form-item>
          <el-button :loading="btnLoading" type="primary" @click="save">保存</el-button>
          <el-button type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script setup>
import {
  createDevice,
  updateDevice,
  findDevice
} from '@/api/system/sysDevice'

defineOptions({
    name: 'DeviceForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'


const route = useRoute()
const router = useRouter()

// 提交按钮loading
const btnLoading = ref(false)

const type = ref('')
const formData = ref({
            device_id: '',
            subscribe_id: '',
            token: '',
            register_ip: '',
            version: '',
            mac: '',
        })
// 验证规则
const rule = reactive({
               device_id : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               subscribe_id : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               token : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findDevice({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
}

init()
// 保存按钮
const save = async() => {
      btnLoading.value = true
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return btnLoading.value = false
            let res
           switch (type.value) {
             case 'create':
               res = await createDevice(formData.value)
               break
             case 'update':
               res = await updateDevice(formData.value)
               break
             default:
               res = await createDevice(formData.value)
               break
           }
           btnLoading.value = false
           if (res.code === 0) {
             ElMessage({
               type: 'success',
               message: '创建/更改成功'
             })
           }
       })
}

// 返回按钮
const back = () => {
    router.go(-1)
}

</script>

<style>
</style>
