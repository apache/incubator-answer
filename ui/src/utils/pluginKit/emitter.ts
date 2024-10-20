type EventName = string;
type EventHandler = () => void;

class SimpleEventEmitter {
  events: Record<EventName, EventHandler[]> = {};

  on(name: EventName, handler: EventHandler) {
    if (!this.events[name]) {
      this.events[name] = [];
    }

    this.events[name].push(handler);
  }

  off(name: EventName, handler?: EventHandler) {
    const handlers = this.events[name];

    if (!handlers || handlers.length === 0) {
      return;
    }

    if (handler) {
      this.events[name] = handlers.filter((func) => func !== handler);
    } else {
      delete this.events[name];
    }
  }

  emit(name: EventName) {
    (this.events[name] || []).forEach((handler) => handler());
  }
}

export default SimpleEventEmitter;
