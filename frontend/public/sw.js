/// <reference no-default-lib="true"/>
/// <reference lib="esnext" />
/// <reference lib="webworker" />
/**@type {ServiceWorkerGlobalScope} sw */
const sw = self;

sw.addEventListener("install", (e) => {
  console.log("Install", { e });
  self.skipWaiting();
});

sw.addEventListener("activate", async (event) => {
  event.waitUntil(self.clients.claim());
});

self.addEventListener("push", (e) => {
  const body = e.data.json();
  if (body.kind === "push-notify") {
    self.registration.showNotification("Update", { body: body.body });
  }
});

sw.addEventListener("notificationclick", function (event) {
  let url = `/`;
  event.notification.close(); // Android needs explicit close.
  event.waitUntil(
    clients.matchAll({ type: "window" }).then((windowClients) => {
      // Check if there is already a window/tab open with the target URL
      for (var i = 0; i < windowClients.length; i++) {
        var client = windowClients[i];
        // If so, just focus it.
        if (client.url === url && "focus" in client) {
          return client.focus();
        }
      }
      // If not, then open the target URL in a new window/tab.
      if (clients.openWindow) {
        return clients.openWindow(url);
      }
    })
  );
});
