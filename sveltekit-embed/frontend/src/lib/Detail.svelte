<script>
  export let detail;
</script>

<table class="table-auto text-left w-full">
  <tbody>
    <tr>
      <th class="px-4 py-2">IP Address</th>
      <td class="px-4 py-2">{detail ? detail.ip : "-"}</td>
    </tr>
    <tr class="bg-blue-500 bg-opacity-5">
      <th class="px-4 py-2">Country</th>
      <td class="px-4 py-2">{detail ? detail.country_name : "-"}</td>
    </tr>
    <tr>
      <th class="px-4 py-2">Region</th>
      <td class="px-4 py-2">{detail ? detail.region : "-"}</td>
    </tr>
    <tr class="bg-blue-500 bg-opacity-5">
      <th class="px-4 py-2">City</th>
      <td class="px-4 py-2">{detail ? detail.city : "-"}</td>
    </tr>
    <tr>
      <th class="px-4 py-2">Zip/Postal Code</th>
      <td class="px-4 py-2">{detail ? detail.postal : "-"}</td>
    </tr>
    <tr class="bg-blue-500 bg-opacity-5">
      <th class="px-4 py-2">ASN</th>
      <td class="px-4 py-2">{detail ? detail.asn : "-"}</td>
    </tr>
    <tr>
      <th class="px-4 py-2">Organization</th>
      <td class="px-4 py-2">{detail ? detail.org : "-"}</td>
    </tr>
  </tbody>
</table>
