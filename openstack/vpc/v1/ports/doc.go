/*


Sample Code, This interface is used to create a port.


    result, err := ports.Create(client, ports.CreateOpts{
        Name:         "EricTestPort",
        NetworkId:    "5ae24488-454f-499c-86c4-c0355704005d",
    }).Extract()

    if err != nil {
        panic(err)
    }

Sample Code, This interface is used to update a port.


    result, err := ports.Update(client,"5e56a480-f337-4985-8ca4-98546cb4fdae", ports.UpdateOpts{
      Name: "ModifiedPort",
    }).Extract()

    if err != nil {
      panic(err)
    }

Sample Code, This interface is used to query a single port.


    result, err := ports.Get(client, "5e56a480-f337-4985-8ca4-98546cb4fdae").Extract()

    if err != nil {
      panic(err)
    }

Sample Code, This interface is used to query ports and to display the ports in a list.


    result, err := ports.List(client, ports.ListOpts{
        Limit: 3,
    }).Extract()

    if err != nil {
        panic(err)
    }

Sample Code, This interface is used to delete a port.

    result := ports.Delete(client, "5e56a480-f337-4985-8ca4-98546cb4fdae")

    if err != nil {
      panic(err)
    }
*/
package ports
