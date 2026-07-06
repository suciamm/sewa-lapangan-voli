export function formatRupiah(value) {
  if (!value) return '';
  const number = value.toString().replace(/[^0-9]/g, '');
  return new Intl.NumberFormat('id-ID').format(number);
}

export function parseRupiah(value) {
  if (!value) return 0;
  return parseInt(value.toString().replace(/[^0-9]/g, '')) || 0;
}
