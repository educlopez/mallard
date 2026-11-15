/*! Env Favicon for PrestaShop - back office behaviour (colour sync, logo drop & preview). */
(function () {
  'use strict';
  var wrap = document.querySelector('.ef-wrap');
  if (!wrap) { return; }
  var HEX = /^#[0-9a-f]{6}$/i;

  // Colour pickers: the swatch, the hex field and the live badge preview stay in sync.
  Array.prototype.forEach.call(wrap.querySelectorAll('.ef-color'), function (box) {
    var picker = box.querySelector('input[type="color"]');
    var hex = box.querySelector('.ef-hex');
    var env = box.closest('.ef-env');
    function paint(v) { if (env) { env.style.setProperty('--c', v); } }
    picker.addEventListener('input', function () { hex.value = picker.value; hex.classList.remove('is-bad'); paint(picker.value); });
    hex.addEventListener('input', function () {
      var v = hex.value.trim();
      if (v.charAt(0) !== '#') { v = '#' + v; }
      var ok = HEX.test(v);
      hex.classList.toggle('is-bad', !ok);
      if (ok) { picker.value = v; paint(v); }
    });
    hex.addEventListener('blur', function () { if (HEX.test(hex.value)) { return; } hex.value = picker.value; hex.classList.remove('is-bad'); });
  });

  // Logo: drag & drop, file name and a quick preview of the chosen file.
  var drop = document.getElementById('ef-drop');
  var input = document.getElementById('ef-file');
  if (!drop || !input) { return; }
  var nameEl = document.getElementById('ef-file-name');
  var raw = document.getElementById('ef-raw');
  var rawImg = document.getElementById('ef-raw-img');

  function show(file) {
    if (!file) { return; }
    nameEl.textContent = file.name;
    nameEl.hidden = false;
    if (file.type && file.type.indexOf('image/') === 0 && window.FileReader) {
      var fr = new FileReader();
      fr.onload = function () { rawImg.src = fr.result; raw.hidden = false; };
      fr.readAsDataURL(file);
    }
  }
  input.addEventListener('change', function () { show(input.files && input.files[0]); });
  ['dragenter', 'dragover'].forEach(function (ev) {
    drop.addEventListener(ev, function (e) { e.preventDefault(); drop.classList.add('is-over'); });
  });
  ['dragleave', 'drop'].forEach(function (ev) {
    drop.addEventListener(ev, function (e) { e.preventDefault(); drop.classList.remove('is-over'); });
  });
  drop.addEventListener('drop', function (e) {
    var files = e.dataTransfer && e.dataTransfer.files;
    if (!files || !files.length) { return; }
    try { input.files = files; } catch (err) { return; }
    show(files[0]);
  });
})();
