<script>
  import Drawer, {
    AppContent,
    Header,
    Title,
    Content,
    Subtitle,
    Scrim,
  } from "@smui/drawer";
  import List, { Item, Text, PrimaryText, SecondaryText } from "@smui/list";

  let { open } = $props();;
</script>

<Drawer variant="modal" fixed={false} bind:open>
  <!-- <Header>
    <Title>Phones</Title>
  </Header> -->
  <Content>
    <List>
      <Item href="/">
        <Text>Home</Text>
      </Item>
      <Item href="/phones">
        <Text>Phones</Text>
      </Item>
    </List>
  </Content>
</Drawer>

<Scrim />
