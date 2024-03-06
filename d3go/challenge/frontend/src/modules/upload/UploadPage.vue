<template>
  <el-card class="upload-container position-center">
    <el-upload accept=".zip" drag action="/admin/upload" :show-file-list="false" :on-success="showTree">
      <el-icon class="el-icon--upload">
        <upload-filled/>
      </el-icon>
      <div class="el-upload__text">
        {{ $t('upload.drag') }}<em>{{ $t('upload.click') }}</em>
      </div>
      <template #tip>
        <div class="el-upload__tip tip">
          {{ $t('upload.hint') }}
        </div>
      </template>
    </el-upload>
    <el-tree :data="fileTree" @node-click="downloadFile" :props="defaultProps"></el-tree>
  </el-card>
</template>

<script lang="ts" setup>
import { UploadFilled } from '@element-plus/icons-vue'
import {ref} from "vue";
import {ElMessage} from "element-plus";
import {useRouter} from "vue-router";
import i18n from "../../locales"
const {t} = i18n.global
let fileTree = ref();
const router = useRouter()
const defaultProps = {
  children: 'children',
  label: 'label',
}
const downloadFile = (node) => {
  if (node.children) {
    return
  }
  window.location.href = node.url
}
const showTree = async (res) => {
  if (res.status_code === -1) {
    ElMessage.warning(t('upload.unauthorized'))
    setTimeout(() => {
      router.push({path: "/admin"})
    }, 1000)
    console.log("unauthorized")
    return
  }
  if (res.status_code === -2) {
    ElMessage.warning(t('upload.forbidden'))
    console.log("forbidden")
    return
  }
  if (res.status_code !== 0) {
    ElMessage.warning(t("upload.failed") + ':' + res.status_msg)
    console.log(res)
    return
  }
  ElMessage.success(t("upload.success") )
  console.log(res.data)
  fileTree.value = res.data
}
</script>

<style scoped>
.upload-container {
  width: 700px;
  height: fit-content;
  padding: 30px 30px 20px 30px;
  z-index: 2;
}

.tip {
  position: relative;
  left: 200px;
}
</style>