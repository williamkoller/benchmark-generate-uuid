<?php
function generateUuid() {
  return sprintf(
    '%04x%04x-%04x-%04x-%04x-%04x%04x%04x',
    mt_rand(0, 0xffff),
    mt_rand(0, 0xffff),
    mt_rand(0, 0xffff),
    mt_rand(0, 0x0fff) | 0x4000,
    mt_rand(0, 0x3fff) | 0x8000,
    mt_rand(0, 0xffff),
    mt_rand(0, 0xffff),
    mt_rand(0, 0xffff)
  );
}

$start = microtime(true);
for ($i = 0; $i < 1000000; $i++) {
    generateUUID();
}
$elapsed = microtime(true) - $start;

echo "\nPHP: " . number_format($elapsed, 5) . " seconds\n";