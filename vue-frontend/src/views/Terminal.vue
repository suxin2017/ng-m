<template>
  <div ref="terRef" @keyup.esc="disableEsc"></div>
</template>
<script setup>
import "xterm/css/xterm.css";
import { onMounted, ref } from "vue";
import { Terminal } from "xterm";
import { FitAddon } from "xterm-addon-fit";

const terRef = ref(null);

onMounted(() => {
  document.addEventListener("keypress", (e) => {});
  var term;
  var websocket = new WebSocket(
    "ws://" +
      window.location.hostname +
      ":" +
      window.location.port +
      "/api/term"
  );
  websocket.binaryType = "arraybuffer";

  function ab2str(buf) {
    return new TextDecoder().decode(new Uint8Array(buf));
  }

  websocket.onopen = function (evt) {
    term = new Terminal({
      rows: 40,
      cols: 80,
      fontFamily: "Cascadia Code",
      screenKeys: true,
      cursorBlink: true,
    });
    const fitAddon = new FitAddon();
    term.loadAddon(fitAddon);

    term.attachCustomKeyEventHandler((event) => {
      if (["v", "c"].includes(event.key) && event.ctrlKey) {
        return false;
      }
    });

    term.onData(function (data) {
      websocket.send(new TextEncoder().encode("\x00" + data));
    });

    term.onResize(function (evt) {
      websocket.send(
        new TextEncoder().encode(
          "\x01" + JSON.stringify({ cols: evt.cols, rows: evt.rows })
        )
      );
    });

    term.onTitleChange(function (title) {
      document.title = title;
      console.log("title");
    });

    term.open(terRef.value);
    fitAddon.fit();
    // term.fit();
    websocket.onmessage = function (evt) {
      if (evt.data instanceof ArrayBuffer) {
        term.write(ab2str(evt.data));
      } else {
        alert(evt.data);
      }
    };

    websocket.onclose = function (evt) {
      term.write("Session terminated");
      term.destroy();
    };

    websocket.onerror = function (evt) {
      if (typeof console.log == "function") {
        console.log(evt);
      }
    };
  };
});
</script>

<style></style>
