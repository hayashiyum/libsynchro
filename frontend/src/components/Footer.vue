<template>
    <footer class="app-footer">
        <div class="footer-content">
            <nav class="footer-nav">
                <a class="footer-text">v0.0.2</a>
                <a href="https://github.com/hayashiyum/libsynchro" target="_blank" class="footer-link"><img
                        src="../assets/images/github-mark-white.png" alt="GitHub icon" width="23" height="23"></a>
                <a href="https://github.com/hayashiyum/libsynchro/blob/v0.0.2/README.md#how-to-use" target="_blank" class="footer-link">How to use</a>
            </nav>
            <p class="copyright">&copy; 2025 hayashiyum. All Rights Reserved.</p>
        </div>
    </footer>
</template>

<script>
export default {
  name: 'Footer'
}
</script>

<style>
.app-footer {
  background-color: #333;
  color: #fff;
  padding: 10px 30px;
  box-shadow: 0 -2px 4px rgba(0, 0, 0, 0.1);
}
.footer-content {
  display: flex;
  flex-direction: column;
  align-items: center;
  max-width: 1200px;
  margin: 0 auto;
}
.footer-nav {
  display: flex;
  justify-content: center;
  align-items: center;
  gap: 1.5rem;
  margin-bottom: 10px;
}
.footer-nav a {
  display: flex;
  align-items: center;
  text-decoration: none;
}
.footer-text {
  color: #fff;
  text-decoration: none;
  font-weight: 600;
}
.footer-link {
  color: #fff;
  text-decoration: none;
  font-weight: 600;
}
.footer-link:hover {
  color: #aaa;
}
.copyright {
  margin: 0;
  font-size: 0.9rem;
  color: #aaa;
}
</style>