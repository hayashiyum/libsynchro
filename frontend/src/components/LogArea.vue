<template>
    <div class="logs-section">
        <div class="log-output">
            <pre
                class="log-text"><span v-for="(log, index) in logs" :key="index" :class="{ 'error-log': log.type === 'ERROR' }">{{ log.message }}<br/></span>
            </pre>
        </div>
    </div>
</template>

<script>
export default {
    name: 'LogArea',
    props: {
      // objects: { type: 'INFO' | 'ERROR', message: string }
      logs: []
    },
}
</script>

<style>
.logs-section {
  margin-bottom: 20px;
  margin-left: 30px;
  margin-right: 30px;
  padding: 16px;
  background-color: #f8f8f8;
  border-radius: 8px;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.05);
  flex-grow: 0;
}
.log-output {
  height: 100%;
  min-height: 200px;
  max-height: 200px;
  overflow-y: auto;
  background-color: #cdcdcd;
  border: 1px solid #ddd;
  border-radius: 4px;
  padding: 10px;
  font-family: 'Courier New', Courier, monospace;
  white-space: pre-wrap;
}
.log-text {
  margin: 0;
  color: #333;
  text-align: left;
  font-size: 0.9rem;
}
.error-log {
  color: red;
}
</style>