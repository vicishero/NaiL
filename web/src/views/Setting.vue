<template>
    <div>
        <n-card title="基本信息" size="small" class="setting-card">
            <div class="base-line avatar">
                <n-avatar
                    class="avatar-img"
                    :size="80"
                    :src="userInfo.avatar"
                />
                <n-button size="small" @click="triggerFileInput">更改头像</n-button>
                <input
                    ref="fileInputRef"
                    type="file"
                    accept="image/png,image/jpg,image/jpeg"
                    style="display: none"
                    @change="onFileSelected"
                />
            </div>
            <!-- 展示图 -->
            <div class="base-line cover">
                <div class="cover-preview" @click="triggerCoverInput">
                    <img v-if="userInfo.cover_image" :src="userInfo.cover_image" class="cover-img" />
                    <div v-else class="cover-placeholder">+ 展示图</div>
                </div>
                <n-button size="small" @click="triggerCoverInput">更改展示图</n-button>
                <input ref="coverInputRef" type="file" accept="image/png,image/jpg,image/jpeg" style="display:none" @change="onCoverSelected" />
            </div>
            <!-- 展示图裁剪弹窗 -->
            <n-modal v-model:show="showCoverCropper" preset="card" title="裁剪展示图" style="width:560px" :mask-closable="false" @after-leave="resetCoverCropper">
                <div class="cropper-container">
                    <vue-cropper v-if="coverCropperSrc" ref="coverCropperRef" :src="coverCropperSrc" :aspect-ratio="3/4" :view-mode="2" :auto-crop-area="1" :guides="true" :background="true" :rotatable="true" :scalable="true" :zoomable="true" :zoom-on-touch="true" :zoom-on-wheel="true" :crop-box-movable="true" :crop-box-resizable="true" :min-crop-box-width="240" :min-crop-box-height="320" @ready="onCoverCropperReady" />
                </div>
                <template #footer>
                    <div class="cropper-footer">
                        <n-button quaternary round @click="showCoverCropper = false">取消</n-button>
                        <n-button type="primary" round :loading="coverCropping" @click="handleCoverCrop">确认</n-button>
                    </div>
                </template>
            </n-modal>
            <!-- 头像裁剪弹窗 -->
            <n-modal
                v-model:show="showCropper"
                preset="card"
                title="裁剪头像"
                style="width: 560px"
                :mask-closable="false"
                @after-leave="resetCropper"
            >
                <div class="cropper-container">
                    <vue-cropper
                        v-if="cropperSrc"
                        ref="cropperRef"
                        :src="cropperSrc"
                        :aspect-ratio="1"
                        :view-mode="2"
                        :auto-crop-area="1"
                        :guides="true"
                        :background="true"
                        :rotatable="true"
                        :scalable="true"
                        :zoomable="true"
                        :zoom-on-touch="true"
                        :zoom-on-wheel="true"
                        :crop-box-movable="true"
                        :crop-box-resizable="true"
                        :min-crop-box-width="256"
                        :min-crop-box-height="256"
                        @ready="onCropperReady"
                    />
                </div>
                <template #footer>
                    <div class="cropper-footer">
                        <n-button quaternary round @click="showCropper = false">取消</n-button>
                        <n-button type="primary" round :loading="cropping" @click="handleCrop">确认</n-button>
                    </div>
                </template>
            </n-modal>
            <div class="base-line">
                <span class="base-label">昵称</span>
                <div v-if="!showNicknameEdit">
                    {{ userInfo.nickname }}
                </div>
                <n-input
                    ref="inputInstRef"
                    v-show="showNicknameEdit"
                    class="nickname-input"
                    v-model:value="userInfo.nickname"
                    type="text"
                    size="small"
                    placeholder="请输入昵称"
                    @blur="handleNicknameChange"
                    :maxlength="16"
                />
                <n-button
                    quaternary
                    round
                    type="success"
                    size="small"
                    v-if="!showNicknameEdit && userInfo.status == 1"
                    @click="handleNicknameShow"
                >
                    <template #icon>
                        <n-icon>
                            <edit />
                        </n-icon>
                    </template>
                </n-button>
            </div>
            <div class="base-line">
                <span class="base-label">用户名</span> @{{
                    userInfo.username
                }}
            </div>
        </n-card>

        <n-card v-if="allowActivation" title="激活码" size="small" class="setting-card">
            <div
                v-if="
                    userInfo.activation &&
                    userInfo.activation.length > 0
                "
            >
                {{ userInfo.activation }}

                <n-button
                    quaternary
                    round
                    type="success"
                    v-if="!showActivation"
                    @click="showActivation = true"
                >
                    重新激活
                </n-button>
            </div>
            <div v-else>
                <n-alert title="激活码激活提示" type="warning">
                    成功激活后后，才能发（公开/好友可见）动态、回复~<br />
                    <a
                        class="hash-link"
                        @click="showActivation = true"
                        v-if="!showActivation"
                    >
                    立即激活
                    </a>
                </n-alert>
            </div>

            <div class="phone-bind-wrap" v-if="showActivation">
                <n-form
                    ref="activateFormRef"
                    :model="activateData"
                    :rules="activateRules"
                >
                    <n-form-item path="activate_code" label="激活码">
                        <n-input
                            :value="activateData.activate_code"
                            @update:value="(v: string) => (activateData.activate_code = v.trim())"
                            placeholder="请输入激活码"
                            @keydown.enter.prevent
                        />
                    </n-form-item>
                    <n-form-item path="img_captcha" label="图形验证码">
                        <div class="captcha-img-wrap">
                            <n-input
                                v-model:value="activateData.imgCaptcha"
                                placeholder="请输入图形验证码后获取验证码"
                            />
                            <div class="captcha-img">
                                <img
                                    v-if="activateData.b64s"
                                    :src="activateData.b64s"
                                    @click="loadCaptcha4Activate"
                                />
                            </div>
                        </div>
                    </n-form-item>
                    <n-row :gutter="[0, 24]">
                        <n-col :span="24">
                            <div class="form-submit-wrap">
                                <n-button
                                    quaternary
                                    round
                                    @click="showActivation = false"
                                >
                                    取消
                                </n-button>
                                <n-button
                                    secondary
                                    round
                                    type="primary"
                                    :loading="activating"
                                    @click="handleActivation"
                                >
                                    激活
                                </n-button>
                            </div>
                        </n-col>
                    </n-row>
                </n-form>
            </div>
        </n-card>

        <n-card title="账户安全" size="small" class="setting-card">
            您已设置密码
            <n-button
                quaternary
                round
                type="success"
                v-if="!showPasswordSetting"
                @click="showPasswordSetting = true"
            >
                重置密码
            </n-button>
            <div class="phone-bind-wrap" v-if="showPasswordSetting">
                <n-form ref="formRef" :model="modelData" :rules="passwordRules">
                    <n-form-item path="old_password" label="旧密码">
                        <n-input
                            v-model:value="modelData.old_password"
                            type="password"
                            placeholder="请输入当前密码"
                            @keydown.enter.prevent
                        />
                    </n-form-item>
                    <n-form-item path="password" label="新密码">
                        <n-input
                            v-model:value="modelData.password"
                            type="password"
                            placeholder="请输入新密码"
                            @input="handlePasswordInput"
                            @keydown.enter.prevent
                        />
                    </n-form-item>
                    <n-form-item
                        ref="rPasswordFormItemRef"
                        first
                        path="reenteredPassword"
                        label="重复密码"
                    >
                        <n-input
                            v-model:value="modelData.reenteredPassword"
                            :disabled="!modelData.password"
                            type="password"
                            placeholder="请再次输入密码"
                            @keydown.enter.prevent
                        />
                    </n-form-item>
                    <n-row :gutter="[0, 24]">
                        <n-col :span="24">
                            <div class="form-submit-wrap">
                                <n-button
                                    quaternary
                                    round
                                    @click="showPasswordSetting = false"
                                >
                                    取消
                                </n-button>
                                <n-button
                                    secondary
                                    round
                                    type="primary"
                                    :loading="passwordSetting"
                                    @click="handleValidateButtonClick"
                                >
                                    更新
                                </n-button>
                            </div>
                        </n-col>
                    </n-row>
                </n-form>
            </div>
        </n-card>
    </div>
</template>

<script setup lang="ts">
import { onMounted, ref, reactive } from 'vue';
import { useStoreMain } from '@/store/main';
import { Edit } from '@vicons/tabler';
import type {
  FormItemRule,
  FormItemInst,
  FormInst,
  InputInst,
} from 'naive-ui';
import { TOKEN_KEY, useStoreUser } from '@/store/user';
import { storeToRefs } from 'pinia';
import { Api } from '@/utils/request';
import VueCropper from 'vue-cropperjs';

const allowActivation =
  import.meta.env.VITE_ALLOW_ACTIVATION.toLowerCase() === 'true';

const storeMain = useStoreMain();
const storeUser = useStoreUser();
const { userInfo } = storeToRefs(storeUser);

const activating = ref(false);
const fileInputRef = ref<HTMLInputElement>();
const cropperRef = ref<any>();
const inputInstRef = ref<InputInst>();
const showNicknameEdit = ref(false);
const showCropper = ref(false);
const cropperSrc = ref('');
const cropping = ref(false);
const passwordSetting = ref(false);
const showPasswordSetting = ref(false);
const showActivation = ref(false);
const activateFormRef = ref<FormInst>();
const formRef = ref<FormInst>();
const rPasswordFormItemRef = ref<FormItemInst>();
const modelData = reactive({
  password: '',
  old_password: '',
  reenteredPassword: '',
});

const activateData = reactive({
  id: '',
  b64s: '',
  imgCaptcha: '',
  activate_code: '',
});

// 触发文件选择
const triggerFileInput = () => {
  fileInputRef.value?.click();
};

// 文件选择后打开裁剪弹窗
const onFileSelected = (e: Event) => {
  const input = e.target as HTMLInputElement;
  const file = input.files?.[0];
  if (!file) return;

  // 校验图片类型
  if (!['image/png', 'image/jpg', 'image/jpeg'].includes(file.type)) {
    window.$message.warning('头像仅允许 png/jpg 格式');
    return;
  }

  // 校验图片大小
  if (file.size > 5 * 1024 * 1024) {
    window.$message.warning('图片大小不能超过5MB');
    return;
  }

  // 读取为 data URL，打开裁剪弹窗
  const reader = new FileReader();
  reader.onload = (ev) => {
    cropperSrc.value = ev.target?.result as string;
    showCropper.value = true;
  };
  reader.readAsDataURL(file);

  // 清空 input 以便重复选择同一文件
  input.value = '';
};

// 裁剪确认，裁剪后上传
const handleCrop = async () => {
  cropping.value = true;
  try {
    const canvas = cropperRef.value?.getCroppedCanvas({
      width: 256,
      height: 256,
    });
    if (!canvas) {
      window.$message.error('裁剪失败');
      cropping.value = false;
      return;
    }

    const blob = await new Promise<Blob>((resolve) => {
      canvas.toBlob(resolve as BlobCallback, 'image/png');
    });

    // 上传裁剪后的图片到 attachment 接口
    const formData = new FormData();
    formData.append('type', 'public/avatar');
    formData.append('file', blob, 'avatar.png');

    const token = 'Bearer ' + localStorage.getItem(TOKEN_KEY);
    const uploadUrl = import.meta.env.VITE_HOST + '/v1/attachment';

    const response = await fetch(uploadUrl, {
      method: 'POST',
      headers: { Authorization: token },
      body: formData,
    });
    const data = await response.json();

    if (data.code === 0) {
      // 更新用户头像
      await Api.v1.user.post.avatar({
        avatar: data.data.content,
      });
      window.$message.success('头像更新成功');
      storeUser.updateUserinfo({
        ...userInfo.value,
        avatar: data.data.content,
      });
      showCropper.value = false;
    } else {
      window.$message.error(data.msg || '上传失败');
    }
  } catch (err) {
    console.error(err);
    window.$message.error('上传失败');
  } finally {
    cropping.value = false;
  }
};

// 关闭裁剪弹窗后清理
const resetCropper = () => {
  cropperSrc.value = '';
};

// cropper 初始化完成后，将选择框默认设为 256x256 并居中
const onCropperReady = () => {
  const cropper = cropperRef.value;
  if (!cropper) return;
  const container = cropper.getContainerData();
  // 如果容器小于 256，则填满容器
  const size = Math.min(256, container.width, container.height);
  const left = (container.width - size) / 2;
  const top = (container.height - size) / 2;
  cropper.setCropBoxData({ left, top, width: size, height: size });
};

// === 展示图裁剪 ===
const coverInputRef = ref<any>();
const coverCropperRef = ref<any>();
const coverCropperSrc = ref('');
const showCoverCropper = ref(false);
const coverCropping = ref(false);

const triggerCoverInput = () => { coverInputRef.value?.click() };

const onCoverSelected = (e: Event) => {
  const input = e.target as HTMLInputElement;
  const file = input.files?.[0];
  if (!file) return;
  if (!['image/png', 'image/jpg', 'image/jpeg'].includes(file.type)) { window.$message.warning('仅允许 png/jpg 格式'); return }
  if (file.size > 5 * 1024 * 1024) { window.$message.warning('图片大小不能超过5MB'); return }
  const reader = new FileReader();
  reader.onload = (ev) => { coverCropperSrc.value = ev.target?.result as string; showCoverCropper.value = true }
  reader.readAsDataURL(file);
  input.value = '';
};

const handleCoverCrop = async () => {
  coverCropping.value = true;
  try {
    const canvas = coverCropperRef.value?.getCroppedCanvas({ width: 480, height: 640 });
    if (!canvas) { window.$message.error('裁剪失败'); coverCropping.value = false; return }
    const blob = await new Promise<Blob>((r) => canvas.toBlob(r as BlobCallback, 'image/jpeg', 0.9));
    const formData = new FormData();
    formData.append('type', 'public/avatar');
    formData.append('file', blob, 'cover.jpg');
    const token = 'Bearer ' + localStorage.getItem(TOKEN_KEY);
    const res = await fetch(import.meta.env.VITE_HOST + '/v1/attachment', { method: 'POST', headers: { Authorization: token }, body: formData });
    const data = await res.json();
    if (data.code === 0) {
      await Api.v1.user.post.avatar({ cover_image: data.data.content });
      storeUser.updateUserinfo({ ...userInfo.value, cover_image: data.data.content });
      window.$message.success('展示图更新成功');
      showCoverCropper.value = false;
    } else { window.$message.error(data.msg || '上传失败') }
  } catch (err) { console.error(err); window.$message.error('上传失败') }
  finally { coverCropping.value = false }
};

const resetCoverCropper = () => { coverCropperSrc.value = '' };

const onCoverCropperReady = () => {
  const c = coverCropperRef.value; if (!c) return;
  const d = c.getContainerData();
  const w = Math.min(240, d.width), h = Math.min(320, d.height);
  c.setCropBoxData({ left: (d.width - w) / 2, top: (d.height - h) / 2, width: w, height: h });
};

const validatePasswordStartWith = (rule: FormItemRule, value: any) => {
  return (
    !!modelData.password &&
    (modelData.password as any).startsWith(value) &&
    (modelData.password as any).length >= value.length
  );
};

const validatePasswordSame = (rule: FormItemRule, value: any) => {
  return value === modelData.password;
};

const handlePasswordInput = () => {
  if (modelData.reenteredPassword) {
    rPasswordFormItemRef.value?.validate({ trigger: 'password-input' });
  }
};

const handleValidateButtonClick = (e: MouseEvent) => {
  e.preventDefault();
  formRef.value?.validate((errors) => {
    if (!errors) {
      passwordSetting.value = true;
      Api.v1.user.post.password({
        password: modelData.password,
        old_password: modelData.old_password,
      })
        .then((res) => {
          passwordSetting.value = false;
          showPasswordSetting.value = false;
          window.$message.success('密码重置成功');

          // 用户退出登录
          storeUser.userLogout();
          storeMain.triggerAuth(true);
          storeMain.triggerAuthKey('signin');
        })
        .catch((err) => {
          passwordSetting.value = false;
        });
    }
  });
};

const handleActivation = (e: MouseEvent) => {
  e.preventDefault();
  activateFormRef.value?.validate((errors) => {
    if (activateData.imgCaptcha === '') {
      window.$message.warning('请输入图片验证码');
      return;
    }
    if (!errors) {
      activating.value = true;
      Api.v1.user.post.activate({
        activate_code: activateData.activate_code,
        captcha_id: activateData.id,
        imgCaptcha: activateData.imgCaptcha,
      })
        .then((res) => {
          activating.value = false;
          showActivation.value = false;
          window.$message.success('激活成功');

          storeUser.updateUserinfo({
            ...userInfo.value,
            activation: activateData.activate_code,
          });

          activateData.id = '';
          activateData.b64s = '';
          activateData.imgCaptcha = '';
          activateData.activate_code = '';
        })
        .catch((err) => {
          activating.value = false;
          if (err.code === 20012) {
            loadCaptcha4Activate();
          }
        });
    }
  });
};

const loadCaptcha4Activate = () => {
  Api.v1.captcha.get._self({})
    .then((res) => {
      activateData.id = res.id;
      activateData.b64s = res.b64s;
    })
    .catch((err) => {
      console.log(err);
    });
};

const handleNicknameChange = () => {
  Api.v1.user.post.nickname({
    nickname: userInfo.value.nickname || '',
  })
    .then((res) => {
      showNicknameEdit.value = false;
      window.$message.success('昵称修改成功');
    })
    .catch((err) => {
      showNicknameEdit.value = true;
    });
};

const activateRules = {
  activate_code: [
    {
      required: true,
      message: '请输入激活码',
      trigger: ['input'],
      validator: (rule: FormItemRule, value: any) => {
        return /\d{6}$/.test(value);
      },
    },
  ],
};
const passwordRules = {
  password: [
    {
      required: true,
      message: '请输入新密码',
    },
  ],
  old_password: [
    {
      required: true,
      message: '请输入旧密码',
    },
  ],
  reenteredPassword: [
    {
      required: true,
      message: '请再次输入密码',
      trigger: ['input', 'blur'],
    },
    {
      validator: validatePasswordStartWith,
      message: '两次密码输入不一致',
      trigger: 'input',
    },
    {
      validator: validatePasswordSame,
      message: '两次密码输入不一致',
      trigger: ['blur', 'password-input'],
    },
  ],
};
const handleNicknameShow = () => {
  showNicknameEdit.value = true;
  setTimeout(() => {
    inputInstRef.value?.focus();
  }, 30);
};
onMounted(() => {
  if (userInfo.value.id == '0') {
    storeMain.triggerAuth(true);
    storeMain.triggerAuthKey('signin');
  }
  loadCaptcha4Activate();
});
</script>

<style>
@import 'cropperjs/dist/cropper.css';
</style>

<style lang="less" scoped>
.setting-card {
    margin-top: -1px;
    border-radius: 0;
    .form-submit-wrap {
        display: flex;
        justify-content: flex-end;
    }

    .base-line {
        line-height: 2;
        display: flex;
        align-items: center;
        .base-label {
            opacity: 0.75;
            margin-right: 12px;
        }

        .nickname-input {
            margin-right: 10px;
            width: 120px;
        }
    }

    .cover {
        display: flex; flex-direction: column; align-items: flex-start; gap: 8px; margin-top: 16px;
        .cover-preview { width: 160px; height: 213px; border-radius: 8px; overflow: hidden; border: 2px dashed var(--border-color,#ddd); cursor: pointer; display: flex; align-items: center; justify-content: center; background: var(--body-color,#f5f5f5); }
        .cover-img { width: 100%; height: 100%; object-fit: cover; }
        .cover-placeholder { color: #999; font-size: 14px; }
    }

    .avatar {
        display: flex;
        flex-direction: column;
        align-items: flex-start;
        margin-bottom: 20px;
        .avatar-img {
            margin-bottom: 10px;
        }
    }

    .hash-link {
        margin-left: 12px;
    }

    .phone-bind-wrap {
        margin-top: 20px;
        .captcha-img-wrap {
            width: 100%;
            display: flex;
            align-items: center;
        }
        .captcha-img {
            width: 125px;
            height: 34px;
            border-radius: 3px;
            margin-left: 10px;
            overflow: hidden;
            cursor: pointer;
            img {
                width: 100%;
                height: 100%;
            }
        }
    }
}
.cropper-container {
    max-width: 100%;
    max-height: 60vh;
    overflow: hidden;
    .cropper-crop-box,
    .cropper-view-box {
        border-radius: 50%;
    }
}
.cropper-footer {
    display: flex;
    justify-content: flex-end;
    gap: 12px;
}
.dark {
    .setting-card {
        background-color: rgba(16, 16, 20, 0.75);
    }
}
</style>