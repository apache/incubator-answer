import { DRAFT_TIMESIGH_STORAGE_KEY as timeSign } from '@/common/constants';

const store = {
  storage: localStorage || window.localStorage,
  set(key: string, value, time?: number): void {
    const t = time || Date.now() + 1000 * 60 * 60 * 24 * 7; // default 7 days
    try {
      this.storage.setItem(key, `${t}${timeSign}${JSON.stringify(value)}`);
    } catch {
      // ignore
      console.error('set storage error: the key is', key);
    }
  },
  get(key: string): any {
    const timeSignLen = timeSign.length;
    let index = 0;
    let time = 0;
    let res;
    try {
      res = this.storage.getItem(key);
    } catch {
      console.error('get storage error: the key is', key);
    }
    if (res) {
      index = res.indexOf(timeSign);
      time = +res.slice(0, index);
      if (time > new Date().getTime()) {
        res = res.slice(index + timeSignLen);
        try {
          res = JSON.parse(res);
        } catch {
          // ignore
        }
      } else {
        // timeout remove storage
        res = null;
        this.storage.removeItem(key);
      }
      return res;
    }

    return res;
  },
  remove(key: string, callback?: () => void): void {
    this.storage.removeItem(key);
    callback?.();
  },
};

export default store;
