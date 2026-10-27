# Forms and inputs

Generic rules (labels, types, 16px, focus, autofill, validation timing, Enter to
submit, loading buttons): see `emil-forms-and-inputs`. This file holds the
Hummingbird and PrestaShop specifics.

Reference: `assets/css/custom.css` (end of the forms block, "One brand focus ring"),
`assets/js/custom.js` ("Checkout autofill", "Account forms").

## Audit script (run per page, mobile viewport)

```js
[...document.querySelectorAll('input,select,textarea')]
  .filter(e => e.offsetParent && !['hidden','submit','button'].includes(e.type))
  .map(e => ({
    n: e.name, t: e.type, fs: parseFloat(getComputedStyle(e).fontSize),
    label: !!((e.id && document.querySelector('label[for="'+e.id+'"]')) || e.closest('label') || e.getAttribute('aria-label')),
    ac: e.getAttribute('autocomplete')
  }))
```

Fail: `fs < 16` (iOS zooms the page on focus), no label, missing `autocomplete`
on identity or address fields. Pages to check: header search, newsletter in the
footer, login, registration, checkout (guest and address), product (size select,
"notify me when available" email), cart.

## 16px and one focus ring

```css
/* One brand focus ring for every field. Bootstrap's default is blue. Inputs stay at 16px so iOS does not zoom */
.form-control:focus,.form-select:focus,.form-check-input:focus,.ps-searchbar__input:focus{border-color:var(--pies-primary);box-shadow:var(--pies-focus-ring);outline:0}
.form-control,.form-select,.ps-searchbar__input{font-size:16px}
```

- Override `:focus`, not only `:focus-visible`: Bootstrap's blue ring is on `:focus`.
- Fields with `border:0` (header search) get a ring from the same rule, so they have a visible state.
- Replace every earlier ad-hoc ring (`.2rem`, `.15rem`, 4px with other alphas) with `var(--pies-focus-ring)`.
- Module-specific inputs are often `font-size:.8125rem`: raise them (newsletter in the footer was 13px).
- Keep any smaller design size with a different technique (padding, smaller label), never a smaller input.

## Autofill tokens

The one-page checkout module renders address fields without `autocomplete`, so
browsers guess. Map by field name, once and after each re-render:

```js
(function () {
  if (!document.body || document.body.id !== 'checkout') return;
  var tokens = {
    email: 'email', firstname: 'given-name', lastname: 'family-name', company: 'organization',
    address1: 'address-line1', address2: 'address-line2', postcode: 'postal-code', city: 'address-level2',
    phone: 'tel', id_country: 'country'
  };
  function apply() {
    Object.keys(tokens).forEach(function (name) {
      document.querySelectorAll('.one-page-checkout [name="' + name + '"]:not([autocomplete])').forEach(function (el) {
        el.setAttribute('autocomplete', tokens[name]);
      });
    });
  }
  apply();
  var root = document.querySelector('.one-page-checkout, #content') || document.body;
  new MutationObserver(apply).observe(root, { childList: true, subtree: true });
})();
```

- Observe `childList` only: setting an attribute then does not retrigger the observer.
- The fields are inside `.one-page-checkout`, not inside `.one-page-checkout__section`: check the real DOM.
- Alias and VAT number stay without a token (no standard token fits).
- The core registration form already ships tokens (`given-name`, `new-password`...) except `birthday`.
- Test by reading the attribute in the DOM. Real browser autofill needs a Chrome profile with saved data (not automatable).

## Birthday field (registration)

A datepicker is the wrong tool for a date of birth (years of month-by-month
navigation). Keep a text field with the right keyboard and typed slashes:

```js
document.querySelectorAll('input[name="birthday"]').forEach(function (el) {
  if (!el.getAttribute('autocomplete')) el.setAttribute('autocomplete', 'bday');
  el.setAttribute('inputmode', 'numeric');
  el.addEventListener('input', function (e) {
    if (e.inputType && e.inputType.indexOf('delete') === 0) return;
    var d = el.value.replace(/\D/g, '').slice(0, 8);
    el.value = d.length > 4 ? d.slice(0, 2) + '/' + d.slice(2, 4) + '/' + d.slice(4)
      : d.length > 2 ? d.slice(0, 2) + '/' + d.slice(2) : d;
  });
});
```

The format `DD/MM/YYYY` is the shop's (es-ES). Do not switch to `<input type="date">`:
it submits `YYYY-MM-DD` and the core validator expects the locale format.

## Other checks that were already fine in the reference

- Every input has a label that focuses it, checkbox rows are clickable from the label.
- Types are right (`email`, `tel`, `password`), login has `username` and `current-password`.
- Search has `autocomplete="off"`. The mobile search overlay does not autofocus on
  open (touch rule). If the product wants the keyboard to appear right after the
  user taps the magnifier, focus on that user gesture.
- Disabled checkout fields (address before the email step) do not take focus: test with an enabled field.
