import { defineStore } from 'pinia';
import { ref } from 'vue';

export const useStoreMain = defineStore('main', () => {
  const refresh = ref(Date.now());
  const refreshTopicFollow = ref(Date.now());
  const theme = ref(localStorage.getItem('PAOPAO_THEME'));
  const authModalShow = ref(false);
  const authModelTab = ref('signin');
  const unreadMsgCount = ref(0);
  const composeModalShow = ref(false);

  function doRefresh(val?: number) {
    refresh.value = val || Date.now();
  }

  function doRefreshTopicFollow() {
    refreshTopicFollow.value = Date.now();
  }

  function updateUnreadMsgCount(count: number) {
    unreadMsgCount.value = count;
  }

  function triggerTheme(t: string) {
    theme.value = t;
  }

  function triggerAuth(status: boolean) {
    authModalShow.value = status;
  }

  function triggerAuthKey(key: string) {
    authModelTab.value = key;
  }

  function triggerCompose(status: boolean) {
    composeModalShow.value = status;
  }

  return {
    refresh,
    refreshTopicFollow,
    theme,
    authModalShow,
    authModelTab,
    unreadMsgCount,
    composeModalShow,
    doRefresh,
    doRefreshTopicFollow,
    updateUnreadMsgCount,
    triggerTheme,
    triggerAuth,
    triggerAuthKey,
    triggerCompose,
  };
});
