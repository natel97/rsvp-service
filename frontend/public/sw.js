/// <reference no-default-lib="true"/>
/// <reference lib="esnext" />
/// <reference lib="webworker" />
/**@type {ServiceWorkerGlobalScope} sw */
const sw = self;

sw.addEventListener("install", () => {
  self.skipWaiting();
});

sw.addEventListener("activate", async (event) => {
  event.waitUntil(self.clients.claim());
});

self.addEventListener("push", (e) => {
  const body = e.data.json();
  if (body.kind === "push-notify") {
    self.registration.showNotification("Update", {
      body: body.body,
      icon: "/icon.svg",
      data: body.url,
      vibrate: [200, 100, 200, 100, 200, 100, 200],
    });
  }
});

sw.addEventListener("notificationclick", function (event) {
  let url = event.notification.data;
  event.notification.close();

  event.waitUntil(
    clients.matchAll({ type: "window" }).then((windowClients) => {
      if (windowClients && windowClients.length) {
        for (var i = 0; i < windowClients.length; i++) {
          var client = windowClients[i];
          if (client.url.includes(url) && "focus" in client) {
            const focus = client.focus();
            client.postMessage({
              type: "OPEN_MODAL",
              data: {
                title: event.notification.title,
                body: event.notification.body,
              },
            });
            return focus;
          }
        }
      }
      if (clients.openWindow) {
        const newWindow = clients.openWindow(url);

        newWindow.then((window) => {
          // IK it's bad, promise isn't actually when window is ready :(
          setTimeout(() => {
            window.postMessage({
              type: "OPEN_MODAL",
              data: {
                title: event.notification.title,
                body: event.notification.body,
              },
            });
          }, 1500);
        });

        return newWindow;
      }
    })
  );
});
