<template>
  <div>
    <!-- the message's input -->
    <input id="input" type="text" />
    <!-- when clicked then a websocket event will be sent to the server, at this example we registered the 'chat' -->
    <button id="sendBtn" disabled>Send</button>
    &nbsp;
    <!-- the messages will be shown here -->
    <pre>{{outputTxt}}</pre>
  </div>
</template>

<script>
import neffos from 'neffos.js'
import { getToken } from '@/libs/util'
export default {
  name: "index",
  data() {
    return {
      outputTxt: ''
    }
  },
  methods: {
    addMessage(msg) {
      this.outputTxt += (msg + "\n")
    },
    handleNamespaceConnectedConn(nsConn) {
      nsConn.emit("Hello from browser client side!");

      let inputTxt = document.getElementById("input");
      let sendBtn = document.getElementById("sendBtn");

      sendBtn.disabled = false;
      sendBtn.onclick = () => {
        const input = inputTxt.value;
        inputTxt.value = "";
        nsConn.emit("chat", input);
        this.addMessage("Me: " + input);
      };
    },
    handleError(reason) {
      console.log(reason);
      window.alert("error: see the dev console");
    },
    async runExample() {
      let that = this
      let wsURL = 'ws://127.0.0.1:8888/echo'
      const enableJWT = true;
      if (enableJWT) {
        const token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjozMjEzMjF9.8waEX7-vPKACa-Soi1pQvW3Rl8QY-SUFcHKTLZI4mvU";
        // wsURL += "?token=" + token;
        // wsURL += "?token=" + getToken()
      }
      try {
        const conn = await neffos.dial(wsURL, {
          default: { // "default" namespace.
            _OnNamespaceConnected: (nsConn, msg) => {
              that.addMessage("connected to namespace: " + msg.Namespace);
              that.handleNamespaceConnectedConn(nsConn)
            },
            _OnNamespaceDisconnect: (nsConn, msg) => {
              that.addMessage("disconnected from namespace: " + msg.Namespace);
            },
            chat: (nsConn, msg) => { // "chat" event.
              that.addMessage(msg.Body);
            }
          }
        },{
          headers: {"X-Username": 'yhm',}
        });
        conn.connect("default");
      } catch (err) {
        this.handleError(err);
      }
    },
  },
  created () {
    this.runExample()
  }
}
</script>

<style scoped>

</style>
