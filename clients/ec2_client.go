package clients

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

/*
* CODE GENERATED AUTOMATICALLY WITH github.com/felipemarinho97/go-mock-gen
 */

// IEC2Client generic client
type IEC2Client interface {
	// Accepts the Convertible Reserved Instance exchange quote described in the
	// GetReservedInstancesExchangeQuote call.
	//
	AcceptReservedInstancesExchangeQuote(arg1 context.Context, arg2 *ec2.AcceptReservedInstancesExchangeQuoteInput, arg3 ...func(*ec2.Options)) (*ec2.AcceptReservedInstancesExchangeQuoteOutput, error)
	// Accepts a request to associate subnets with a transit gateway multicast domain.
	//
	AcceptTransitGatewayMulticastDomainAssociations(arg1 context.Context, arg2 *ec2.AcceptTransitGatewayMulticastDomainAssociationsInput, arg3 ...func(*ec2.Options)) (*ec2.AcceptTransitGatewayMulticastDomainAssociationsOutput, error)
	// Accepts a transit gateway peering attachment request. The peering attachment
	// must be in the pendingAcceptance state.
	//
	AcceptTransitGatewayPeeringAttachment(arg1 context.Context, arg2 *ec2.AcceptTransitGatewayPeeringAttachmentInput, arg3 ...func(*ec2.Options)) (*ec2.AcceptTransitGatewayPeeringAttachmentOutput, error)
	// Accepts a request to attach a VPC to a transit gateway. The VPC attachment must
	// be in the pendingAcceptance state. Use DescribeTransitGatewayVpcAttachments to
	// view your pending VPC attachment requests. Use RejectTransitGatewayVpcAttachment
	// to reject a VPC attachment request.
	//
	AcceptTransitGatewayVpcAttachment(arg1 context.Context, arg2 *ec2.AcceptTransitGatewayVpcAttachmentInput, arg3 ...func(*ec2.Options)) (*ec2.AcceptTransitGatewayVpcAttachmentOutput, error)
	// Accepts one or more interface VPC endpoint connection requests to your VPC
	// endpoint service.
	//
	AcceptVpcEndpointConnections(arg1 context.Context, arg2 *ec2.AcceptVpcEndpointConnectionsInput, arg3 ...func(*ec2.Options)) (*ec2.AcceptVpcEndpointConnectionsOutput, error)
	// Accept a VPC peering connection request. To accept a request, the VPC peering
	// connection must be in the pending-acceptance state, and you must be the owner of
	// the peer VPC. Use DescribeVpcPeeringConnections to view your outstanding VPC
	// peering connection requests. For an inter-Region VPC peering connection request,
	// you must accept the VPC peering connection in the Region of the accepter VPC.
	//
	AcceptVpcPeeringConnection(arg1 context.Context, arg2 *ec2.AcceptVpcPeeringConnectionInput, arg3 ...func(*ec2.Options)) (*ec2.AcceptVpcPeeringConnectionOutput, error)
	// Advertises an IPv4 or IPv6 address range that is provisioned for use with your
	// Amazon Web Services resources through bring your own IP addresses (BYOIP). You
	// can perform this operation at most once every 10 seconds, even if you specify
	// different address ranges each time. We recommend that you stop advertising the
	// BYOIP CIDR from other locations when you advertise it from Amazon Web Services.
	// To minimize down time, you can configure your Amazon Web Services resources to
	// use an address from a BYOIP CIDR before it is advertised, and then
	// simultaneously stop advertising it from the current location and start
	// advertising it through Amazon Web Services. It can take a few minutes before
	// traffic to the specified addresses starts routing to Amazon Web Services because
	// of BGP propagation delays. To stop advertising the BYOIP CIDR, use
	// WithdrawByoipCidr.
	//
	AdvertiseByoipCidr(arg1 context.Context, arg2 *ec2.AdvertiseByoipCidrInput, arg3 ...func(*ec2.Options)) (*ec2.AdvertiseByoipCidrOutput, error)
	// Allocates an Elastic IP address to your Amazon Web Services account. After you
	// allocate the Elastic IP address you can associate it with an instance or network
	// interface. After you release an Elastic IP address, it is released to the IP
	// address pool and can be allocated to a different Amazon Web Services account.
	// You can allocate an Elastic IP address from an address pool owned by Amazon Web
	// Services or from an address pool created from a public IPv4 address range that
	// you have brought to Amazon Web Services for use with your Amazon Web Services
	// resources using bring your own IP addresses (BYOIP). For more information, see
	// Bring Your Own IP Addresses (BYOIP)
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-byoip.html) in the
	// Amazon Elastic Compute Cloud User Guide. [EC2-VPC] If you release an Elastic IP
	// address, you might be able to recover it. You cannot recover an Elastic IP
	// address that you released after it is allocated to another Amazon Web Services
	// account. You cannot recover an Elastic IP address for EC2-Classic. To attempt to
	// recover an Elastic IP address that you released, specify it in this operation.
	// An Elastic IP address is for use either in the EC2-Classic platform or in a VPC.
	// By default, you can allocate 5 Elastic IP addresses for EC2-Classic per Region
	// and 5 Elastic IP addresses for EC2-VPC per Region. For more information, see
	// Elastic IP Addresses
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/elastic-ip-addresses-eip.html)
	// in the Amazon Elastic Compute Cloud User Guide. You can allocate a carrier IP
	// address which is a public IP address from a telecommunication carrier, to a
	// network interface which resides in a subnet in a Wavelength Zone (for example an
	// EC2 instance).
	//
	AllocateAddress(arg1 context.Context, arg2 *ec2.AllocateAddressInput, arg3 ...func(*ec2.Options)) (*ec2.AllocateAddressOutput, error)
	// Allocates a Dedicated Host to your account. At a minimum, specify the supported
	// instance type or instance family, the Availability Zone in which to allocate the
	// host, and the number of hosts to allocate.
	//
	AllocateHosts(arg1 context.Context, arg2 *ec2.AllocateHostsInput, arg3 ...func(*ec2.Options)) (*ec2.AllocateHostsOutput, error)
	// Allocate a CIDR from an IPAM pool. In IPAM, an allocation is a CIDR assignment
	// from an IPAM pool to another resource or IPAM pool. For more information, see
	// Allocate CIDRs
	// (https://docs.aws.amazon.com/vpc/latest/ipam/allocate-cidrs-ipam.html) in the
	// Amazon VPC IPAM User Guide.
	//
	AllocateIpamPoolCidr(arg1 context.Context, arg2 *ec2.AllocateIpamPoolCidrInput, arg3 ...func(*ec2.Options)) (*ec2.AllocateIpamPoolCidrOutput, error)
	// Applies a security group to the association between the target network and the
	// Client VPN endpoint. This action replaces the existing security groups with the
	// specified security groups.
	//
	ApplySecurityGroupsToClientVpnTargetNetwork(arg1 context.Context, arg2 *ec2.ApplySecurityGroupsToClientVpnTargetNetworkInput, arg3 ...func(*ec2.Options)) (*ec2.ApplySecurityGroupsToClientVpnTargetNetworkOutput, error)
	// Assigns one or more IPv6 addresses to the specified network interface. You can
	// specify one or more specific IPv6 addresses, or you can specify the number of
	// IPv6 addresses to be automatically assigned from within the subnet's IPv6 CIDR
	// block range. You can assign as many IPv6 addresses to a network interface as you
	// can assign private IPv4 addresses, and the limit varies per instance type. For
	// information, see IP Addresses Per Network Interface Per Instance Type
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/using-eni.html#AvailableIpPerENI)
	// in the Amazon Elastic Compute Cloud User Guide. You must specify either the IPv6
	// addresses or the IPv6 address count in the request. You can optionally use
	// Prefix Delegation on the network interface. You must specify either the IPV6
	// Prefix Delegation prefixes, or the IPv6 Prefix Delegation count. For
	// information, see  Assigning prefixes to Amazon EC2 network interfaces
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-prefix-eni.html) in the
	// Amazon Elastic Compute Cloud User Guide.
	//
	AssignIpv6Addresses(arg1 context.Context, arg2 *ec2.AssignIpv6AddressesInput, arg3 ...func(*ec2.Options)) (*ec2.AssignIpv6AddressesOutput, error)
	// Assigns one or more secondary private IP addresses to the specified network
	// interface. You can specify one or more specific secondary IP addresses, or you
	// can specify the number of secondary IP addresses to be automatically assigned
	// within the subnet's CIDR block range. The number of secondary IP addresses that
	// you can assign to an instance varies by instance type. For information about
	// instance types, see Instance Types
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instance-types.html) in the
	// Amazon Elastic Compute Cloud User Guide. For more information about Elastic IP
	// addresses, see Elastic IP Addresses
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/elastic-ip-addresses-eip.html)
	// in the Amazon Elastic Compute Cloud User Guide. When you move a secondary
	// private IP address to another network interface, any Elastic IP address that is
	// associated with the IP address is also moved. Remapping an IP address is an
	// asynchronous operation. When you move an IP address from one network interface
	// to another, check network/interfaces/macs/mac/local-ipv4s in the instance
	// metadata to confirm that the remapping is complete. You must specify either the
	// IP addresses or the IP address count in the request. You can optionally use
	// Prefix Delegation on the network interface. You must specify either the IPv4
	// Prefix Delegation prefixes, or the IPv4 Prefix Delegation count. For
	// information, see  Assigning prefixes to Amazon EC2 network interfaces
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-prefix-eni.html) in the
	// Amazon Elastic Compute Cloud User Guide.
	//
	AssignPrivateIpAddresses(arg1 context.Context, arg2 *ec2.AssignPrivateIpAddressesInput, arg3 ...func(*ec2.Options)) (*ec2.AssignPrivateIpAddressesOutput, error)
	// Associates an Elastic IP address, or carrier IP address (for instances that are
	// in subnets in Wavelength Zones) with an instance or a network interface. Before
	// you can use an Elastic IP address, you must allocate it to your account. An
	// Elastic IP address is for use in either the EC2-Classic platform or in a VPC.
	// For more information, see Elastic IP Addresses
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/elastic-ip-addresses-eip.html)
	// in the Amazon Elastic Compute Cloud User Guide. [EC2-Classic, VPC in an
	// EC2-VPC-only account] If the Elastic IP address is already associated with a
	// different instance, it is disassociated from that instance and associated with
	// the specified instance. If you associate an Elastic IP address with an instance
	// that has an existing Elastic IP address, the existing address is disassociated
	// from the instance, but remains allocated to your account. [VPC in an EC2-Classic
	// account] If you don't specify a private IP address, the Elastic IP address is
	// associated with the primary IP address. If the Elastic IP address is already
	// associated with a different instance or a network interface, you get an error
	// unless you allow reassociation. You cannot associate an Elastic IP address with
	// an instance or network interface that has an existing Elastic IP address.
	// [Subnets in Wavelength Zones] You can associate an IP address from the
	// telecommunication carrier to the instance or network interface. You cannot
	// associate an Elastic IP address with an interface in a different network border
	// group. This is an idempotent operation. If you perform the operation more than
	// once, Amazon EC2 doesn't return an error, and you may be charged for each time
	// the Elastic IP address is remapped to the same instance. For more information,
	// see the Elastic IP Addresses section of Amazon EC2 Pricing
	// (http://aws.amazon.com/ec2/pricing/).
	//
	AssociateAddress(arg1 context.Context, arg2 *ec2.AssociateAddressInput, arg3 ...func(*ec2.Options)) (*ec2.AssociateAddressOutput, error)
	// Associates a target network with a Client VPN endpoint. A target network is a
	// subnet in a VPC. You can associate multiple subnets from the same VPC with a
	// Client VPN endpoint. You can associate only one subnet in each Availability
	// Zone. We recommend that you associate at least two subnets to provide
	// Availability Zone redundancy. If you specified a VPC when you created the Client
	// VPN endpoint or if you have previous subnet associations, the specified subnet
	// must be in the same VPC. To specify a subnet that's in a different VPC, you must
	// first modify the Client VPN endpoint (ModifyClientVpnEndpoint) and change the
	// VPC that's associated with it.
	//
	AssociateClientVpnTargetNetwork(arg1 context.Context, arg2 *ec2.AssociateClientVpnTargetNetworkInput, arg3 ...func(*ec2.Options)) (*ec2.AssociateClientVpnTargetNetworkOutput, error)
	// Associates a set of DHCP options (that you've previously created) with the
	// specified VPC, or associates no DHCP options with the VPC. After you associate
	// the options with the VPC, any existing instances and all new instances that you
	// launch in that VPC use the options. You don't need to restart or relaunch the
	// instances. They automatically pick up the changes within a few hours, depending
	// on how frequently the instance renews its DHCP lease. You can explicitly renew
	// the lease using the operating system on the instance. For more information, see
	// DHCP options sets
	// (https://docs.aws.amazon.com/vpc/latest/userguide/VPC_DHCP_Options.html) in the
	// Amazon Virtual Private Cloud User Guide.
	//
	AssociateDhcpOptions(arg1 context.Context, arg2 *ec2.AssociateDhcpOptionsInput, arg3 ...func(*ec2.Options)) (*ec2.AssociateDhcpOptionsOutput, error)
	// Associates an Identity and Access Management (IAM) role with an Certificate
	// Manager (ACM) certificate. This enables the certificate to be used by the ACM
	// for Nitro Enclaves application inside an enclave. For more information, see
	// Certificate Manager for Nitro Enclaves
	// (https://docs.aws.amazon.com/enclaves/latest/user/nitro-enclave-refapp.html) in
	// the Amazon Web Services Nitro Enclaves User Guide. When the IAM role is
	// associated with the ACM certificate, the certificate, certificate chain, and
	// encrypted private key are placed in an Amazon S3 bucket that only the associated
	// IAM role can access. The private key of the certificate is encrypted with an
	// Amazon Web Services managed key that has an attached attestation-based key
	// policy. To enable the IAM role to access the Amazon S3 object, you must grant it
	// permission to call s3:GetObject on the Amazon S3 bucket returned by the command.
	// To enable the IAM role to access the KMS key, you must grant it permission to
	// call kms:Decrypt on the KMS key returned by the command. For more information,
	// see  Grant the role permission to access the certificate and encryption key
	// (https://docs.aws.amazon.com/enclaves/latest/user/nitro-enclave-refapp.html#add-policy)
	// in the Amazon Web Services Nitro Enclaves User Guide.
	//
	AssociateEnclaveCertificateIamRole(arg1 context.Context, arg2 *ec2.AssociateEnclaveCertificateIamRoleInput, arg3 ...func(*ec2.Options)) (*ec2.AssociateEnclaveCertificateIamRoleOutput, error)
	// Associates an IAM instance profile with a running or stopped instance. You
	// cannot associate more than one IAM instance profile with an instance.
	//
	AssociateIamInstanceProfile(arg1 context.Context, arg2 *ec2.AssociateIamInstanceProfileInput, arg3 ...func(*ec2.Options)) (*ec2.AssociateIamInstanceProfileOutput, error)
	// Associates one or more targets with an event window. Only one type of target
	// (instance IDs, Dedicated Host IDs, or tags) can be specified with an event
	// window. For more information, see Define event windows for scheduled events
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/event-windows.html) in the
	// Amazon EC2 User Guide.
	//
	AssociateInstanceEventWindow(arg1 context.Context, arg2 *ec2.AssociateInstanceEventWindowInput, arg3 ...func(*ec2.Options)) (*ec2.AssociateInstanceEventWindowOutput, error)
	// Associates a subnet in your VPC or an internet gateway or virtual private
	// gateway attached to your VPC with a route table in your VPC. This association
	// causes traffic from the subnet or gateway to be routed according to the routes
	// in the route table. The action returns an association ID, which you need in
	// order to disassociate the route table later. A route table can be associated
	// with multiple subnets. For more information, see Route tables
	// (https://docs.aws.amazon.com/vpc/latest/userguide/VPC_Route_Tables.html) in the
	// Amazon Virtual Private Cloud User Guide.
	//
	AssociateRouteTable(arg1 context.Context, arg2 *ec2.AssociateRouteTableInput, arg3 ...func(*ec2.Options)) (*ec2.AssociateRouteTableOutput, error)
	// Associates a CIDR block with your subnet. You can only associate a single IPv6
	// CIDR block with your subnet. An IPv6 CIDR block must have a prefix length of
	// /64.
	//
	AssociateSubnetCidrBlock(arg1 context.Context, arg2 *ec2.AssociateSubnetCidrBlockInput, arg3 ...func(*ec2.Options)) (*ec2.AssociateSubnetCidrBlockOutput, error)
	// Associates the specified subnets and transit gateway attachments with the
	// specified transit gateway multicast domain. The transit gateway attachment must
	// be in the available state before you can add a resource. Use
	// DescribeTransitGatewayAttachments
	// (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribeTransitGatewayAttachments.html)
	// to see the state of the attachment.
	//
	AssociateTransitGatewayMulticastDomain(arg1 context.Context, arg2 *ec2.AssociateTransitGatewayMulticastDomainInput, arg3 ...func(*ec2.Options)) (*ec2.AssociateTransitGatewayMulticastDomainOutput, error)
	// Associates the specified transit gateway attachment with a transit gateway
	// policy table.
	//
	AssociateTransitGatewayPolicyTable(arg1 context.Context, arg2 *ec2.AssociateTransitGatewayPolicyTableInput, arg3 ...func(*ec2.Options)) (*ec2.AssociateTransitGatewayPolicyTableOutput, error)
	// Associates the specified attachment with the specified transit gateway route
	// table. You can associate only one route table with an attachment.
	//
	AssociateTransitGatewayRouteTable(arg1 context.Context, arg2 *ec2.AssociateTransitGatewayRouteTableInput, arg3 ...func(*ec2.Options)) (*ec2.AssociateTransitGatewayRouteTableOutput, error)
	// This API action is currently in limited preview only. If you are interested in
	// using this feature, contact your account manager. Associates a branch network
	// interface with a trunk network interface. Before you create the association, run
	// the create-network-interface
	// (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_CreateNetworkInterface.html)
	// command and set --interface-type to trunk. You must also create a network
	// interface for each branch network interface that you want to associate with the
	// trunk network interface.
	//
	AssociateTrunkInterface(arg1 context.Context, arg2 *ec2.AssociateTrunkInterfaceInput, arg3 ...func(*ec2.Options)) (*ec2.AssociateTrunkInterfaceOutput, error)
	// Associates a CIDR block with your VPC. You can associate a secondary IPv4 CIDR
	// block, an Amazon-provided IPv6 CIDR block, or an IPv6 CIDR block from an IPv6
	// address pool that you provisioned through bring your own IP addresses (BYOIP
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-byoip.html)). The IPv6
	// CIDR block size is fixed at /56. You must specify one of the following in the
	// request: an IPv4 CIDR block, an IPv6 pool, or an Amazon-provided IPv6 CIDR
	// block. For more information about associating CIDR blocks with your VPC and
	// applicable restrictions, see VPC and subnet sizing
	// (https://docs.aws.amazon.com/vpc/latest/userguide/VPC_Subnets.html#VPC_Sizing)
	// in the Amazon Virtual Private Cloud User Guide.
	//
	AssociateVpcCidrBlock(arg1 context.Context, arg2 *ec2.AssociateVpcCidrBlockInput, arg3 ...func(*ec2.Options)) (*ec2.AssociateVpcCidrBlockOutput, error)
	// Links an EC2-Classic instance to a ClassicLink-enabled VPC through one or more
	// of the VPC's security groups. You cannot link an EC2-Classic instance to more
	// than one VPC at a time. You can only link an instance that's in the running
	// state. An instance is automatically unlinked from a VPC when it's stopped - you
	// can link it to the VPC again when you restart it. After you've linked an
	// instance, you cannot change the VPC security groups that are associated with it.
	// To change the security groups, you must first unlink the instance, and then link
	// it again. Linking your instance to a VPC is sometimes referred to as attaching
	// your instance.
	//
	AttachClassicLinkVpc(arg1 context.Context, arg2 *ec2.AttachClassicLinkVpcInput, arg3 ...func(*ec2.Options)) (*ec2.AttachClassicLinkVpcOutput, error)
	// Attaches an internet gateway or a virtual private gateway to a VPC, enabling
	// connectivity between the internet and the VPC. For more information about your
	// VPC and internet gateway, see the Amazon Virtual Private Cloud User Guide
	// (https://docs.aws.amazon.com/vpc/latest/userguide/).
	//
	AttachInternetGateway(arg1 context.Context, arg2 *ec2.AttachInternetGatewayInput, arg3 ...func(*ec2.Options)) (*ec2.AttachInternetGatewayOutput, error)
	// Attaches a network interface to an instance.
	//
	AttachNetworkInterface(arg1 context.Context, arg2 *ec2.AttachNetworkInterfaceInput, arg3 ...func(*ec2.Options)) (*ec2.AttachNetworkInterfaceOutput, error)
	// Attaches an EBS volume to a running or stopped instance and exposes it to the
	// instance with the specified device name. Encrypted EBS volumes must be attached
	// to instances that support Amazon EBS encryption. For more information, see
	// Amazon EBS encryption
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSEncryption.html) in the
	// Amazon Elastic Compute Cloud User Guide. After you attach an EBS volume, you
	// must make it available. For more information, see Make an EBS volume available
	// for use
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ebs-using-volumes.html). If
	// a volume has an Amazon Web Services Marketplace product code:
	//
	// * The volume can
	// be attached only to a stopped instance.
	//
	// * Amazon Web Services Marketplace
	// product codes are copied from the volume to the instance.
	//
	// * You must be
	// subscribed to the product.
	//
	// * The instance type and operating system of the
	// instance must support the product. For example, you can't detach a volume from a
	// Windows instance and attach it to a Linux instance.
	//
	// For more information, see
	// Attach an Amazon EBS volume to an instance
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ebs-attaching-volume.html)
	// in the Amazon Elastic Compute Cloud User Guide.
	//
	AttachVolume(arg1 context.Context, arg2 *ec2.AttachVolumeInput, arg3 ...func(*ec2.Options)) (*ec2.AttachVolumeOutput, error)
	// Attaches a virtual private gateway to a VPC. You can attach one virtual private
	// gateway to one VPC at a time. For more information, see Amazon Web Services
	// Site-to-Site VPN (https://docs.aws.amazon.com/vpn/latest/s2svpn/VPC_VPN.html) in
	// the Amazon Web Services Site-to-Site VPN User Guide.
	//
	AttachVpnGateway(arg1 context.Context, arg2 *ec2.AttachVpnGatewayInput, arg3 ...func(*ec2.Options)) (*ec2.AttachVpnGatewayOutput, error)
	// Adds an ingress authorization rule to a Client VPN endpoint. Ingress
	// authorization rules act as firewall rules that grant access to networks. You
	// must configure ingress authorization rules to enable clients to access resources
	// in Amazon Web Services or on-premises networks.
	//
	AuthorizeClientVpnIngress(arg1 context.Context, arg2 *ec2.AuthorizeClientVpnIngressInput, arg3 ...func(*ec2.Options)) (*ec2.AuthorizeClientVpnIngressOutput, error)
	// [VPC only] Adds the specified outbound (egress) rules to a security group for
	// use with a VPC. An outbound rule permits instances to send traffic to the
	// specified IPv4 or IPv6 CIDR address ranges, or to the instances that are
	// associated with the specified source security groups. You specify a protocol for
	// each rule (for example, TCP). For the TCP and UDP protocols, you must also
	// specify the destination port or port range. For the ICMP protocol, you must also
	// specify the ICMP type and code. You can use -1 for the type or code to mean all
	// types or all codes. Rule changes are propagated to affected instances as quickly
	// as possible. However, a small delay might occur. For information about VPC
	// security group quotas, see Amazon VPC quotas
	// (https://docs.aws.amazon.com/vpc/latest/userguide/amazon-vpc-limits.html).
	//
	AuthorizeSecurityGroupEgress(arg1 context.Context, arg2 *ec2.AuthorizeSecurityGroupEgressInput, arg3 ...func(*ec2.Options)) (*ec2.AuthorizeSecurityGroupEgressOutput, error)
	// Adds the specified inbound (ingress) rules to a security group. An inbound rule
	// permits instances to receive traffic from the specified IPv4 or IPv6 CIDR
	// address range, or from the instances that are associated with the specified
	// destination security groups. You specify a protocol for each rule (for example,
	// TCP). For TCP and UDP, you must also specify the destination port or port range.
	// For ICMP/ICMPv6, you must also specify the ICMP/ICMPv6 type and code. You can
	// use -1 to mean all types or all codes. Rule changes are propagated to instances
	// within the security group as quickly as possible. However, a small delay might
	// occur. For more information about VPC security group quotas, see Amazon VPC
	// quotas
	// (https://docs.aws.amazon.com/vpc/latest/userguide/amazon-vpc-limits.html).
	//
	AuthorizeSecurityGroupIngress(arg1 context.Context, arg2 *ec2.AuthorizeSecurityGroupIngressInput, arg3 ...func(*ec2.Options)) (*ec2.AuthorizeSecurityGroupIngressOutput, error)
	// Bundles an Amazon instance store-backed Windows instance. During bundling, only
	// the root device volume (C:\) is bundled. Data on other instance store volumes is
	// not preserved. This action is not applicable for Linux/Unix instances or Windows
	// instances that are backed by Amazon EBS.
	//
	BundleInstance(arg1 context.Context, arg2 *ec2.BundleInstanceInput, arg3 ...func(*ec2.Options)) (*ec2.BundleInstanceOutput, error)
	// Cancels a bundling operation for an instance store-backed Windows instance.
	//
	CancelBundleTask(arg1 context.Context, arg2 *ec2.CancelBundleTaskInput, arg3 ...func(*ec2.Options)) (*ec2.CancelBundleTaskOutput, error)
	// Cancels the specified Capacity Reservation, releases the reserved capacity, and
	// changes the Capacity Reservation's state to cancelled. Instances running in the
	// reserved capacity continue running until you stop them. Stopped instances that
	// target the Capacity Reservation can no longer launch. Modify these instances to
	// either target a different Capacity Reservation, launch On-Demand Instance
	// capacity, or run in any open Capacity Reservation that has matching attributes
	// and sufficient capacity.
	//
	CancelCapacityReservation(arg1 context.Context, arg2 *ec2.CancelCapacityReservationInput, arg3 ...func(*ec2.Options)) (*ec2.CancelCapacityReservationOutput, error)
	// Cancels one or more Capacity Reservation Fleets. When you cancel a Capacity
	// Reservation Fleet, the following happens:
	//
	// * The Capacity Reservation Fleet's
	// status changes to cancelled.
	//
	// * The individual Capacity Reservations in the
	// Fleet are cancelled. Instances running in the Capacity Reservations at the time
	// of cancelling the Fleet continue to run in shared capacity.
	//
	// * The Fleet stops
	// creating new Capacity Reservations.
	//
	CancelCapacityReservationFleets(arg1 context.Context, arg2 *ec2.CancelCapacityReservationFleetsInput, arg3 ...func(*ec2.Options)) (*ec2.CancelCapacityReservationFleetsOutput, error)
	// Cancels an active conversion task. The task can be the import of an instance or
	// volume. The action removes all artifacts of the conversion, including a
	// partially uploaded volume or instance. If the conversion is complete or is in
	// the process of transferring the final disk image, the command fails and returns
	// an exception. For more information, see Importing a Virtual Machine Using the
	// Amazon EC2 CLI
	// (https://docs.aws.amazon.com/AWSEC2/latest/CommandLineReference/ec2-cli-vmimport-export.html).
	//
	CancelConversionTask(arg1 context.Context, arg2 *ec2.CancelConversionTaskInput, arg3 ...func(*ec2.Options)) (*ec2.CancelConversionTaskOutput, error)
	// Cancels an active export task. The request removes all artifacts of the export,
	// including any partially-created Amazon S3 objects. If the export task is
	// complete or is in the process of transferring the final disk image, the command
	// fails and returns an error.
	//
	CancelExportTask(arg1 context.Context, arg2 *ec2.CancelExportTaskInput, arg3 ...func(*ec2.Options)) (*ec2.CancelExportTaskOutput, error)
	// Cancels an in-process import virtual machine or import snapshot task.
	//
	CancelImportTask(arg1 context.Context, arg2 *ec2.CancelImportTaskInput, arg3 ...func(*ec2.Options)) (*ec2.CancelImportTaskOutput, error)
	// Cancels the specified Reserved Instance listing in the Reserved Instance
	// Marketplace. For more information, see Reserved Instance Marketplace
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ri-market-general.html) in
	// the Amazon EC2 User Guide.
	//
	CancelReservedInstancesListing(arg1 context.Context, arg2 *ec2.CancelReservedInstancesListingInput, arg3 ...func(*ec2.Options)) (*ec2.CancelReservedInstancesListingOutput, error)
	// Cancels the specified Spot Fleet requests. After you cancel a Spot Fleet
	// request, the Spot Fleet launches no new Spot Instances. You must specify whether
	// the Spot Fleet should also terminate its Spot Instances. If you terminate the
	// instances, the Spot Fleet request enters the cancelled_terminating state.
	// Otherwise, the Spot Fleet request enters the cancelled_running state and the
	// instances continue to run until they are interrupted or you terminate them
	// manually.
	//
	CancelSpotFleetRequests(arg1 context.Context, arg2 *ec2.CancelSpotFleetRequestsInput, arg3 ...func(*ec2.Options)) (*ec2.CancelSpotFleetRequestsOutput, error)
	// Cancels one or more Spot Instance requests. Canceling a Spot Instance request
	// does not terminate running Spot Instances associated with the request.
	//
	CancelSpotInstanceRequests(arg1 context.Context, arg2 *ec2.CancelSpotInstanceRequestsInput, arg3 ...func(*ec2.Options)) (*ec2.CancelSpotInstanceRequestsOutput, error)
	// Determines whether a product code is associated with an instance. This action
	// can only be used by the owner of the product code. It is useful when a product
	// code owner must verify whether another user's instance is eligible for support.
	//
	ConfirmProductInstance(arg1 context.Context, arg2 *ec2.ConfirmProductInstanceInput, arg3 ...func(*ec2.Options)) (*ec2.ConfirmProductInstanceOutput, error)
	// Copies the specified Amazon FPGA Image (AFI) to the current Region.
	//
	CopyFpgaImage(arg1 context.Context, arg2 *ec2.CopyFpgaImageInput, arg3 ...func(*ec2.Options)) (*ec2.CopyFpgaImageOutput, error)
	// Initiates the copy of an AMI. You can copy an AMI from one Region to another, or
	// from a Region to an Outpost. You can't copy an AMI from an Outpost to a Region,
	// from one Outpost to another, or within the same Outpost. To copy an AMI to
	// another partition, see CreateStoreImageTask
	// (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_CreateStoreImageTask.html).
	// To copy an AMI from one Region to another, specify the source Region using
	// the
	//
	// SourceRegion parameter, and specify the destination Region using its
	// endpoint. Copies of encrypted backing snapshots for the AMI are encrypted.
	// Copies of unencrypted backing snapshots remain unencrypted, unless you set
	// Encrypted during the copy operation. You cannot create an unencrypted copy of an
	// encrypted backing snapshot. To copy an AMI from a Region to an Outpost, specify
	// the source Region using the
	//
	// SourceRegion parameter, and specify the ARN of the
	// destination Outpost using DestinationOutpostArn. Backing snapshots copied to an
	// Outpost are encrypted by default using the default encryption key for the
	// Region, or a different key that you specify in the request using KmsKeyId.
	// Outposts do not support unencrypted snapshots. For more information,  Amazon EBS
	// local snapshots on Outposts
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/snapshots-outposts.html#ami)
	// in the Amazon Elastic Compute Cloud User Guide. For more information about the
	// prerequisites and limits when copying an AMI, see Copying an AMI
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/CopyingAMIs.html) in the
	// Amazon Elastic Compute Cloud User Guide.
	//
	CopyImage(arg1 context.Context, arg2 *ec2.CopyImageInput, arg3 ...func(*ec2.Options)) (*ec2.CopyImageOutput, error)
	// Copies a point-in-time snapshot of an EBS volume and stores it in Amazon S3. You
	// can copy a snapshot within the same Region, from one Region to another, or from
	// a Region to an Outpost. You can't copy a snapshot from an Outpost to a Region,
	// from one Outpost to another, or within the same Outpost. You can use the
	// snapshot to create EBS volumes or Amazon Machine Images (AMIs). When copying
	// snapshots to a Region, copies of encrypted EBS snapshots remain encrypted.
	// Copies of unencrypted snapshots remain unencrypted, unless you enable encryption
	// for the snapshot copy operation. By default, encrypted snapshot copies use the
	// default Key Management Service (KMS) KMS key; however, you can specify a
	// different KMS key. To copy an encrypted snapshot that has been shared from
	// another account, you must have permissions for the KMS key used to encrypt the
	// snapshot. Snapshots copied to an Outpost are encrypted by default using the
	// default encryption key for the Region, or a different key that you specify in
	// the request using KmsKeyId. Outposts do not support unencrypted snapshots. For
	// more information,  Amazon EBS local snapshots on Outposts
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/snapshots-outposts.html#ami)
	// in the Amazon Elastic Compute Cloud User Guide. Snapshots created by copying
	// another snapshot have an arbitrary volume ID that should not be used for any
	// purpose. For more information, see Copy an Amazon EBS snapshot
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ebs-copy-snapshot.html) in
	// the Amazon Elastic Compute Cloud User Guide.
	//
	CopySnapshot(arg1 context.Context, arg2 *ec2.CopySnapshotInput, arg3 ...func(*ec2.Options)) (*ec2.CopySnapshotOutput, error)
	// Creates a new Capacity Reservation with the specified attributes. Capacity
	// Reservations enable you to reserve capacity for your Amazon EC2 instances in a
	// specific Availability Zone for any duration. This gives you the flexibility to
	// selectively add capacity reservations and still get the Regional RI discounts
	// for that usage. By creating Capacity Reservations, you ensure that you always
	// have access to Amazon EC2 capacity when you need it, for as long as you need it.
	// For more information, see Capacity Reservations
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-capacity-reservations.html)
	// in the Amazon EC2 User Guide. Your request to create a Capacity Reservation
	// could fail if Amazon EC2 does not have sufficient capacity to fulfill the
	// request. If your request fails due to Amazon EC2 capacity constraints, either
	// try again at a later time, try in a different Availability Zone, or request a
	// smaller capacity reservation. If your application is flexible across instance
	// types and sizes, try to create a Capacity Reservation with different instance
	// attributes. Your request could also fail if the requested quantity exceeds your
	// On-Demand Instance limit for the selected instance type. If your request fails
	// due to limit constraints, increase your On-Demand Instance limit for the
	// required instance type and try again. For more information about increasing your
	// instance limits, see Amazon EC2 Service Quotas
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-resource-limits.html)
	// in the Amazon EC2 User Guide.
	//
	CreateCapacityReservation(arg1 context.Context, arg2 *ec2.CreateCapacityReservationInput, arg3 ...func(*ec2.Options)) (*ec2.CreateCapacityReservationOutput, error)
	// Creates a Capacity Reservation Fleet. For more information, see Create a
	// Capacity Reservation Fleet
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/work-with-cr-fleets.html#create-crfleet)
	// in the Amazon EC2 User Guide.
	//
	CreateCapacityReservationFleet(arg1 context.Context, arg2 *ec2.CreateCapacityReservationFleetInput, arg3 ...func(*ec2.Options)) (*ec2.CreateCapacityReservationFleetOutput, error)
	// Creates a carrier gateway. For more information about carrier gateways, see
	// Carrier gateways
	// (https://docs.aws.amazon.com/wavelength/latest/developerguide/how-wavelengths-work.html#wavelength-carrier-gateway)
	// in the Amazon Web Services Wavelength Developer Guide.
	//
	CreateCarrierGateway(arg1 context.Context, arg2 *ec2.CreateCarrierGatewayInput, arg3 ...func(*ec2.Options)) (*ec2.CreateCarrierGatewayOutput, error)
	// Creates a Client VPN endpoint. A Client VPN endpoint is the resource you create
	// and configure to enable and manage client VPN sessions. It is the destination
	// endpoint at which all client VPN sessions are terminated.
	//
	CreateClientVpnEndpoint(arg1 context.Context, arg2 *ec2.CreateClientVpnEndpointInput, arg3 ...func(*ec2.Options)) (*ec2.CreateClientVpnEndpointOutput, error)
	// Adds a route to a network to a Client VPN endpoint. Each Client VPN endpoint has
	// a route table that describes the available destination network routes. Each
	// route in the route table specifies the path for traﬃc to speciﬁc resources or
	// networks.
	//
	CreateClientVpnRoute(arg1 context.Context, arg2 *ec2.CreateClientVpnRouteInput, arg3 ...func(*ec2.Options)) (*ec2.CreateClientVpnRouteOutput, error)
	// Provides information to Amazon Web Services about your customer gateway device.
	// The customer gateway device is the appliance at your end of the VPN connection.
	// You must provide the IP address of the customer gateway device’s external
	// interface. The IP address must be static and can be behind a device performing
	// network address translation (NAT). For devices that use Border Gateway Protocol
	// (BGP), you can also provide the device's BGP Autonomous System Number (ASN). You
	// can use an existing ASN assigned to your network. If you don't have an ASN
	// already, you can use a private ASN. For more information, see Customer gateway
	// options for your Site-to-Site VPN connection
	// (https://docs.aws.amazon.com/vpn/latest/s2svpn/cgw-options.html) in the Amazon
	// Web Services Site-to-Site VPN User Guide. To create more than one customer
	// gateway with the same VPN type, IP address, and BGP ASN, specify a unique device
	// name for each customer gateway. An identical request returns information about
	// the existing customer gateway; it doesn't create a new customer gateway.
	//
	CreateCustomerGateway(arg1 context.Context, arg2 *ec2.CreateCustomerGatewayInput, arg3 ...func(*ec2.Options)) (*ec2.CreateCustomerGatewayOutput, error)
	// Creates a default subnet with a size /20 IPv4 CIDR block in the specified
	// Availability Zone in your default VPC. You can have only one default subnet per
	// Availability Zone. For more information, see Creating a default subnet
	// (https://docs.aws.amazon.com/vpc/latest/userguide/default-vpc.html#create-default-subnet)
	// in the Amazon Virtual Private Cloud User Guide.
	//
	CreateDefaultSubnet(arg1 context.Context, arg2 *ec2.CreateDefaultSubnetInput, arg3 ...func(*ec2.Options)) (*ec2.CreateDefaultSubnetOutput, error)
	// Creates a default VPC with a size /16 IPv4 CIDR block and a default subnet in
	// each Availability Zone. For more information about the components of a default
	// VPC, see Default VPC and default subnets
	// (https://docs.aws.amazon.com/vpc/latest/userguide/default-vpc.html) in the
	// Amazon Virtual Private Cloud User Guide. You cannot specify the components of
	// the default VPC yourself. If you deleted your previous default VPC, you can
	// create a default VPC. You cannot have more than one default VPC per Region. If
	// your account supports EC2-Classic, you cannot use this action to create a
	// default VPC in a Region that supports EC2-Classic. If you want a default VPC in
	// a Region that supports EC2-Classic, see "I really want a default VPC for my
	// existing EC2 account. Is that possible?" in the Default VPCs FAQ
	// (http://aws.amazon.com/vpc/faqs/#Default_VPCs).
	//
	CreateDefaultVpc(arg1 context.Context, arg2 *ec2.CreateDefaultVpcInput, arg3 ...func(*ec2.Options)) (*ec2.CreateDefaultVpcOutput, error)
	// Creates a set of DHCP options for your VPC. After creating the set, you must
	// associate it with the VPC, causing all existing and new instances that you
	// launch in the VPC to use this set of DHCP options. The following are the
	// individual DHCP options you can specify. For more information about the options,
	// see RFC 2132 (http://www.ietf.org/rfc/rfc2132.txt).
	//
	// * domain-name-servers - The
	// IP addresses of up to four domain name servers, or AmazonProvidedDNS. The
	// default DHCP option set specifies AmazonProvidedDNS. If specifying more than one
	// domain name server, specify the IP addresses in a single parameter, separated by
	// commas. To have your instance receive a custom DNS hostname as specified in
	// domain-name, you must set domain-name-servers to a custom DNS server.
	//
	// *
	// domain-name - If you're using AmazonProvidedDNS in us-east-1, specify
	// ec2.internal. If you're using AmazonProvidedDNS in another Region, specify
	// region.compute.internal (for example, ap-northeast-1.compute.internal).
	// Otherwise, specify a domain name (for example, ExampleCompany.com). This value
	// is used to complete unqualified DNS hostnames. Important: Some Linux operating
	// systems accept multiple domain names separated by spaces. However, Windows and
	// other Linux operating systems treat the value as a single domain, which results
	// in unexpected behavior. If your DHCP options set is associated with a VPC that
	// has instances with multiple operating systems, specify only one domain name.
	//
	// *
	// ntp-servers - The IP addresses of up to four Network Time Protocol (NTP)
	// servers.
	//
	// * netbios-name-servers - The IP addresses of up to four NetBIOS name
	// servers.
	//
	// * netbios-node-type - The NetBIOS node type (1, 2, 4, or 8). We
	// recommend that you specify 2 (broadcast and multicast are not currently
	// supported). For more information about these node types, see RFC 2132
	// (http://www.ietf.org/rfc/rfc2132.txt).
	//
	// Your VPC automatically starts out with a
	// set of DHCP options that includes only a DNS server that we provide
	// (AmazonProvidedDNS). If you create a set of options, and if your VPC has an
	// internet gateway, make sure to set the domain-name-servers option either to
	// AmazonProvidedDNS or to a domain name server of your choice. For more
	// information, see DHCP options sets
	// (https://docs.aws.amazon.com/vpc/latest/userguide/VPC_DHCP_Options.html) in the
	// Amazon Virtual Private Cloud User Guide.
	//
	CreateDhcpOptions(arg1 context.Context, arg2 *ec2.CreateDhcpOptionsInput, arg3 ...func(*ec2.Options)) (*ec2.CreateDhcpOptionsOutput, error)
	// [IPv6 only] Creates an egress-only internet gateway for your VPC. An egress-only
	// internet gateway is used to enable outbound communication over IPv6 from
	// instances in your VPC to the internet, and prevents hosts outside of your VPC
	// from initiating an IPv6 connection with your instance.
	//
	CreateEgressOnlyInternetGateway(arg1 context.Context, arg2 *ec2.CreateEgressOnlyInternetGatewayInput, arg3 ...func(*ec2.Options)) (*ec2.CreateEgressOnlyInternetGatewayOutput, error)
	// Launches an EC2 Fleet. You can create a single EC2 Fleet that includes multiple
	// launch specifications that vary by instance type, AMI, Availability Zone, or
	// subnet. For more information, see EC2 Fleet
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-fleet.html) in the
	// Amazon EC2 User Guide.
	//
	CreateFleet(arg1 context.Context, arg2 *ec2.CreateFleetInput, arg3 ...func(*ec2.Options)) (*ec2.CreateFleetOutput, error)
	// Creates one or more flow logs to capture information about IP traffic for a
	// specific network interface, subnet, or VPC. Flow log data for a monitored
	// network interface is recorded as flow log records, which are log events
	// consisting of fields that describe the traffic flow. For more information, see
	// Flow log records
	// (https://docs.aws.amazon.com/vpc/latest/userguide/flow-logs.html#flow-log-records)
	// in the Amazon Virtual Private Cloud User Guide. When publishing to CloudWatch
	// Logs, flow log records are published to a log group, and each network interface
	// has a unique log stream in the log group. When publishing to Amazon S3, flow log
	// records for all of the monitored network interfaces are published to a single
	// log file object that is stored in the specified bucket. For more information,
	// see VPC Flow Logs
	// (https://docs.aws.amazon.com/vpc/latest/userguide/flow-logs.html) in the Amazon
	// Virtual Private Cloud User Guide.
	//
	CreateFlowLogs(arg1 context.Context, arg2 *ec2.CreateFlowLogsInput, arg3 ...func(*ec2.Options)) (*ec2.CreateFlowLogsOutput, error)
	// Creates an Amazon FPGA Image (AFI) from the specified design checkpoint (DCP).
	// The create operation is asynchronous. To verify that the AFI is ready for use,
	// check the output logs. An AFI contains the FPGA bitstream that is ready to
	// download to an FPGA. You can securely deploy an AFI on multiple FPGA-accelerated
	// instances. For more information, see the Amazon Web Services FPGA Hardware
	// Development Kit (https://github.com/aws/aws-fpga/).
	//
	CreateFpgaImage(arg1 context.Context, arg2 *ec2.CreateFpgaImageInput, arg3 ...func(*ec2.Options)) (*ec2.CreateFpgaImageOutput, error)
	// Creates an Amazon EBS-backed AMI from an Amazon EBS-backed instance that is
	// either running or stopped. By default, when Amazon EC2 creates the new AMI, it
	// reboots the instance so that it can take snapshots of the attached volumes while
	// data is at rest, in order to ensure a consistent state. You can set the NoReboot
	// parameter to true in the API request, or use the --no-reboot option in the CLI
	// to prevent Amazon EC2 from shutting down and rebooting the instance. If you
	// choose to bypass the shutdown and reboot process by setting the NoReboot
	// parameter to true in the API request, or by using the --no-reboot option in the
	// CLI, we can't guarantee the file system integrity of the created image. If you
	// customized your instance with instance store volumes or Amazon EBS volumes in
	// addition to the root device volume, the new AMI contains block device mapping
	// information for those volumes. When you launch an instance from this new AMI,
	// the instance automatically launches with those additional volumes. For more
	// information, see Creating Amazon EBS-Backed Linux AMIs
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/creating-an-ami-ebs.html)
	// in the Amazon Elastic Compute Cloud User Guide.
	//
	CreateImage(arg1 context.Context, arg2 *ec2.CreateImageInput, arg3 ...func(*ec2.Options)) (*ec2.CreateImageOutput, error)
	// Creates an event window in which scheduled events for the associated Amazon EC2
	// instances can run. You can define either a set of time ranges or a cron
	// expression when creating the event window, but not both. All event window times
	// are in UTC. You can create up to 200 event windows per Amazon Web Services
	// Region. When you create the event window, targets (instance IDs, Dedicated Host
	// IDs, or tags) are not yet associated with it. To ensure that the event window
	// can be used, you must associate one or more targets with it by using the
	// AssociateInstanceEventWindow API. Event windows are applicable only for
	// scheduled events that stop, reboot, or terminate instances. Event windows are
	// not applicable for:
	//
	// * Expedited scheduled events and network maintenance
	// events.
	//
	// * Unscheduled maintenance such as AutoRecovery and unplanned
	// reboots.
	//
	// For more information, see Define event windows for scheduled events
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/event-windows.html) in the
	// Amazon EC2 User Guide.
	//
	CreateInstanceEventWindow(arg1 context.Context, arg2 *ec2.CreateInstanceEventWindowInput, arg3 ...func(*ec2.Options)) (*ec2.CreateInstanceEventWindowOutput, error)
	// Exports a running or stopped instance to an Amazon S3 bucket. For information
	// about the supported operating systems, image formats, and known limitations for
	// the types of instances you can export, see Exporting an instance as a VM Using
	// VM Import/Export
	// (https://docs.aws.amazon.com/vm-import/latest/userguide/vmexport.html) in the VM
	// Import/Export User Guide.
	//
	CreateInstanceExportTask(arg1 context.Context, arg2 *ec2.CreateInstanceExportTaskInput, arg3 ...func(*ec2.Options)) (*ec2.CreateInstanceExportTaskOutput, error)
	// Creates an internet gateway for use with a VPC. After creating the internet
	// gateway, you attach it to a VPC using AttachInternetGateway. For more
	// information about your VPC and internet gateway, see the Amazon Virtual Private
	// Cloud User Guide (https://docs.aws.amazon.com/vpc/latest/userguide/).
	//
	CreateInternetGateway(arg1 context.Context, arg2 *ec2.CreateInternetGatewayInput, arg3 ...func(*ec2.Options)) (*ec2.CreateInternetGatewayOutput, error)
	// Create an IPAM. Amazon VPC IP Address Manager (IPAM) is a VPC feature that you
	// can use to automate your IP address management workflows including assigning,
	// tracking, troubleshooting, and auditing IP addresses across Amazon Web Services
	// Regions and accounts throughout your Amazon Web Services Organization. For more
	// information, see Create an IPAM
	// (https://docs.aws.amazon.com/vpc/latest/ipam/create-ipam.html) in the Amazon VPC
	// IPAM User Guide.
	//
	CreateIpam(arg1 context.Context, arg2 *ec2.CreateIpamInput, arg3 ...func(*ec2.Options)) (*ec2.CreateIpamOutput, error)
	// Create an IP address pool for Amazon VPC IP Address Manager (IPAM). In IPAM, a
	// pool is a collection of contiguous IP addresses CIDRs. Pools enable you to
	// organize your IP addresses according to your routing and security needs. For
	// example, if you have separate routing and security needs for development and
	// production applications, you can create a pool for each. For more information,
	// see Create a top-level pool
	// (https://docs.aws.amazon.com/vpc/latest/ipam/create-top-ipam.html) in the Amazon
	// VPC IPAM User Guide.
	//
	CreateIpamPool(arg1 context.Context, arg2 *ec2.CreateIpamPoolInput, arg3 ...func(*ec2.Options)) (*ec2.CreateIpamPoolOutput, error)
	// Create an IPAM scope. In IPAM, a scope is the highest-level container within
	// IPAM. An IPAM contains two default scopes. Each scope represents the IP space
	// for a single network. The private scope is intended for all private IP address
	// space. The public scope is intended for all public IP address space. Scopes
	// enable you to reuse IP addresses across multiple unconnected networks without
	// causing IP address overlap or conflict. For more information, see Add a scope
	// (https://docs.aws.amazon.com/vpc/latest/ipam/add-scope-ipam.html) in the Amazon
	// VPC IPAM User Guide.
	//
	CreateIpamScope(arg1 context.Context, arg2 *ec2.CreateIpamScopeInput, arg3 ...func(*ec2.Options)) (*ec2.CreateIpamScopeOutput, error)
	// Creates an ED25519 or 2048-bit RSA key pair with the specified name and in the
	// specified PEM or PPK format. Amazon EC2 stores the public key and displays the
	// private key for you to save to a file. The private key is returned as an
	// unencrypted PEM encoded PKCS#1 private key or an unencrypted PPK formatted
	// private key for use with PuTTY. If a key with the specified name already exists,
	// Amazon EC2 returns an error. The key pair returned to you is available only in
	// the Amazon Web Services Region in which you create it. If you prefer, you can
	// create your own key pair using a third-party tool and upload it to any Region
	// using ImportKeyPair. You can have up to 5,000 key pairs per Amazon Web Services
	// Region. For more information, see Amazon EC2 key pairs
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-key-pairs.html) in the
	// Amazon Elastic Compute Cloud User Guide.
	//
	CreateKeyPair(arg1 context.Context, arg2 *ec2.CreateKeyPairInput, arg3 ...func(*ec2.Options)) (*ec2.CreateKeyPairOutput, error)
	// Creates a launch template. A launch template contains the parameters to launch
	// an instance. When you launch an instance using RunInstances, you can specify a
	// launch template instead of providing the launch parameters in the request. For
	// more information, see Launch an instance from a launch template
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-launch-templates.html)
	// in the Amazon Elastic Compute Cloud User Guide. If you want to clone an existing
	// launch template as the basis for creating a new launch template, you can use the
	// Amazon EC2 console. The API, SDKs, and CLI do not support cloning a template.
	// For more information, see Create a launch template from an existing launch
	// template
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-launch-templates.html#create-launch-template-from-existing-launch-template)
	// in the Amazon Elastic Compute Cloud User Guide.
	//
	CreateLaunchTemplate(arg1 context.Context, arg2 *ec2.CreateLaunchTemplateInput, arg3 ...func(*ec2.Options)) (*ec2.CreateLaunchTemplateOutput, error)
	// Creates a new version of a launch template. You can specify an existing version
	// of launch template from which to base the new version. Launch template versions
	// are numbered in the order in which they are created. You cannot specify, change,
	// or replace the numbering of launch template versions. Launch templates are
	// immutable; after you create a launch template, you can't modify it. Instead, you
	// can create a new version of the launch template that includes any changes you
	// require. For more information, see Modify a launch template (manage launch
	// template versions)
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-launch-templates.html#manage-launch-template-versions)
	// in the Amazon Elastic Compute Cloud User Guide.
	//
	CreateLaunchTemplateVersion(arg1 context.Context, arg2 *ec2.CreateLaunchTemplateVersionInput, arg3 ...func(*ec2.Options)) (*ec2.CreateLaunchTemplateVersionOutput, error)
	// Creates a static route for the specified local gateway route table.
	//
	CreateLocalGatewayRoute(arg1 context.Context, arg2 *ec2.CreateLocalGatewayRouteInput, arg3 ...func(*ec2.Options)) (*ec2.CreateLocalGatewayRouteOutput, error)
	// Associates the specified VPC with the specified local gateway route table.
	//
	CreateLocalGatewayRouteTableVpcAssociation(arg1 context.Context, arg2 *ec2.CreateLocalGatewayRouteTableVpcAssociationInput, arg3 ...func(*ec2.Options)) (*ec2.CreateLocalGatewayRouteTableVpcAssociationOutput, error)
	// Creates a managed prefix list. You can specify one or more entries for the
	// prefix list. Each entry consists of a CIDR block and an optional description.
	//
	CreateManagedPrefixList(arg1 context.Context, arg2 *ec2.CreateManagedPrefixListInput, arg3 ...func(*ec2.Options)) (*ec2.CreateManagedPrefixListOutput, error)
	// Creates a NAT gateway in the specified subnet. This action creates a network
	// interface in the specified subnet with a private IP address from the IP address
	// range of the subnet. You can create either a public NAT gateway or a private NAT
	// gateway. With a public NAT gateway, internet-bound traffic from a private subnet
	// can be routed to the NAT gateway, so that instances in a private subnet can
	// connect to the internet. With a private NAT gateway, private communication is
	// routed across VPCs and on-premises networks through a transit gateway or virtual
	// private gateway. Common use cases include running large workloads behind a small
	// pool of allowlisted IPv4 addresses, preserving private IPv4 addresses, and
	// communicating between overlapping networks. For more information, see NAT
	// gateways (https://docs.aws.amazon.com/vpc/latest/userguide/vpc-nat-gateway.html)
	// in the Amazon Virtual Private Cloud User Guide.
	//
	CreateNatGateway(arg1 context.Context, arg2 *ec2.CreateNatGatewayInput, arg3 ...func(*ec2.Options)) (*ec2.CreateNatGatewayOutput, error)
	// Creates a network ACL in a VPC. Network ACLs provide an optional layer of
	// security (in addition to security groups) for the instances in your VPC. For
	// more information, see Network ACLs
	// (https://docs.aws.amazon.com/vpc/latest/userguide/VPC_ACLs.html) in the Amazon
	// Virtual Private Cloud User Guide.
	//
	CreateNetworkAcl(arg1 context.Context, arg2 *ec2.CreateNetworkAclInput, arg3 ...func(*ec2.Options)) (*ec2.CreateNetworkAclOutput, error)
	// Creates an entry (a rule) in a network ACL with the specified rule number. Each
	// network ACL has a set of numbered ingress rules and a separate set of numbered
	// egress rules. When determining whether a packet should be allowed in or out of a
	// subnet associated with the ACL, we process the entries in the ACL according to
	// the rule numbers, in ascending order. Each network ACL has a set of ingress
	// rules and a separate set of egress rules. We recommend that you leave room
	// between the rule numbers (for example, 100, 110, 120, ...), and not number them
	// one right after the other (for example, 101, 102, 103, ...). This makes it
	// easier to add a rule between existing ones without having to renumber the rules.
	// After you add an entry, you can't modify it; you must either replace it, or
	// create an entry and delete the old one. For more information about network ACLs,
	// see Network ACLs
	// (https://docs.aws.amazon.com/vpc/latest/userguide/VPC_ACLs.html) in the Amazon
	// Virtual Private Cloud User Guide.
	//
	CreateNetworkAclEntry(arg1 context.Context, arg2 *ec2.CreateNetworkAclEntryInput, arg3 ...func(*ec2.Options)) (*ec2.CreateNetworkAclEntryOutput, error)
	// Creates a Network Access Scope. Amazon Web Services Network Access Analyzer
	// enables cloud networking and cloud operations teams to verify that their
	// networks on Amazon Web Services conform to their network security and governance
	// objectives. For more information, see the Amazon Web Services Network Access
	// Analyzer Guide
	// (https://docs.aws.amazon.com/vpc/latest/network-access-analyzer/).
	//
	CreateNetworkInsightsAccessScope(arg1 context.Context, arg2 *ec2.CreateNetworkInsightsAccessScopeInput, arg3 ...func(*ec2.Options)) (*ec2.CreateNetworkInsightsAccessScopeOutput, error)
	// Creates a path to analyze for reachability. Reachability Analyzer enables you to
	// analyze and debug network reachability between two resources in your virtual
	// private cloud (VPC). For more information, see What is Reachability Analyzer
	// (https://docs.aws.amazon.com/vpc/latest/reachability/).
	//
	CreateNetworkInsightsPath(arg1 context.Context, arg2 *ec2.CreateNetworkInsightsPathInput, arg3 ...func(*ec2.Options)) (*ec2.CreateNetworkInsightsPathOutput, error)
	// Creates a network interface in the specified subnet. For more information about
	// network interfaces, see Elastic Network Interfaces
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/using-eni.html) in the
	// Amazon Virtual Private Cloud User Guide.
	//
	CreateNetworkInterface(arg1 context.Context, arg2 *ec2.CreateNetworkInterfaceInput, arg3 ...func(*ec2.Options)) (*ec2.CreateNetworkInterfaceOutput, error)
	// Grants an Amazon Web Services-authorized account permission to attach the
	// specified network interface to an instance in their account. You can grant
	// permission to a single Amazon Web Services account only, and only one account at
	// a time.
	//
	CreateNetworkInterfacePermission(arg1 context.Context, arg2 *ec2.CreateNetworkInterfacePermissionInput, arg3 ...func(*ec2.Options)) (*ec2.CreateNetworkInterfacePermissionOutput, error)
	// Creates a placement group in which to launch instances. The strategy of the
	// placement group determines how the instances are organized within the group. A
	// cluster placement group is a logical grouping of instances within a single
	// Availability Zone that benefit from low network latency, high network
	// throughput. A spread placement group places instances on distinct hardware. A
	// partition placement group places groups of instances in different partitions,
	// where instances in one partition do not share the same hardware with instances
	// in another partition. For more information, see Placement groups
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/placement-groups.html) in
	// the Amazon EC2 User Guide.
	//
	CreatePlacementGroup(arg1 context.Context, arg2 *ec2.CreatePlacementGroupInput, arg3 ...func(*ec2.Options)) (*ec2.CreatePlacementGroupOutput, error)
	// Creates a public IPv4 address pool. A public IPv4 pool is an EC2 IP address pool
	// required for the public IPv4 CIDRs that you own and bring to Amazon Web Services
	// to manage with IPAM. IPv6 addresses you bring to Amazon Web Services, however,
	// use IPAM pools only. To monitor the status of pool creation, use
	// DescribePublicIpv4Pools
	// (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribePublicIpv4Pools.html).
	//
	CreatePublicIpv4Pool(arg1 context.Context, arg2 *ec2.CreatePublicIpv4PoolInput, arg3 ...func(*ec2.Options)) (*ec2.CreatePublicIpv4PoolOutput, error)
	// Creates a root volume replacement task for an Amazon EC2 instance. The root
	// volume can either be restored to its initial launch state, or it can be restored
	// using a specific snapshot. For more information, see Replace a root volume
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ebs-restoring-volume.html#replace-root)
	// in the Amazon Elastic Compute Cloud User Guide.
	//
	CreateReplaceRootVolumeTask(arg1 context.Context, arg2 *ec2.CreateReplaceRootVolumeTaskInput, arg3 ...func(*ec2.Options)) (*ec2.CreateReplaceRootVolumeTaskOutput, error)
	// Creates a listing for Amazon EC2 Standard Reserved Instances to be sold in the
	// Reserved Instance Marketplace. You can submit one Standard Reserved Instance
	// listing at a time. To get a list of your Standard Reserved Instances, you can
	// use the DescribeReservedInstances operation. Only Standard Reserved Instances
	// can be sold in the Reserved Instance Marketplace. Convertible Reserved Instances
	// cannot be sold. The Reserved Instance Marketplace matches sellers who want to
	// resell Standard Reserved Instance capacity that they no longer need with buyers
	// who want to purchase additional capacity. Reserved Instances bought and sold
	// through the Reserved Instance Marketplace work like any other Reserved
	// Instances. To sell your Standard Reserved Instances, you must first register as
	// a seller in the Reserved Instance Marketplace. After completing the registration
	// process, you can create a Reserved Instance Marketplace listing of some or all
	// of your Standard Reserved Instances, and specify the upfront price to receive
	// for them. Your Standard Reserved Instance listings then become available for
	// purchase. To view the details of your Standard Reserved Instance listing, you
	// can use the DescribeReservedInstancesListings operation. For more information,
	// see Reserved Instance Marketplace
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ri-market-general.html) in
	// the Amazon EC2 User Guide.
	//
	CreateReservedInstancesListing(arg1 context.Context, arg2 *ec2.CreateReservedInstancesListingInput, arg3 ...func(*ec2.Options)) (*ec2.CreateReservedInstancesListingOutput, error)
	// Starts a task that restores an AMI from an Amazon S3 object that was previously
	// created by using CreateStoreImageTask
	// (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_CreateStoreImageTask.html).
	// To use this API, you must have the required permissions. For more information,
	// see Permissions for storing and restoring AMIs using Amazon S3
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ami-store-restore.html#ami-s3-permissions)
	// in the Amazon Elastic Compute Cloud User Guide. For more information, see Store
	// and restore an AMI using Amazon S3
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ami-store-restore.html) in
	// the Amazon Elastic Compute Cloud User Guide.
	//
	CreateRestoreImageTask(arg1 context.Context, arg2 *ec2.CreateRestoreImageTaskInput, arg3 ...func(*ec2.Options)) (*ec2.CreateRestoreImageTaskOutput, error)
	// Creates a route in a route table within a VPC. You must specify either a
	// destination CIDR block or a prefix list ID. You must also specify exactly one of
	// the resources from the parameter list. When determining how to route traffic, we
	// use the route with the most specific match. For example, traffic is destined for
	// the IPv4 address 192.0.2.3, and the route table includes the following two IPv4
	// routes:
	//
	// * 192.0.2.0/24 (goes to some target A)
	//
	// * 192.0.2.0/28 (goes to some
	// target B)
	//
	// Both routes apply to the traffic destined for 192.0.2.3. However, the
	// second route in the list covers a smaller number of IP addresses and is
	// therefore more specific, so we use that route to determine where to target the
	// traffic. For more information about route tables, see Route tables
	// (https://docs.aws.amazon.com/vpc/latest/userguide/VPC_Route_Tables.html) in the
	// Amazon Virtual Private Cloud User Guide.
	//
	CreateRoute(arg1 context.Context, arg2 *ec2.CreateRouteInput, arg3 ...func(*ec2.Options)) (*ec2.CreateRouteOutput, error)
	// Creates a route table for the specified VPC. After you create a route table, you
	// can add routes and associate the table with a subnet. For more information, see
	// Route tables
	// (https://docs.aws.amazon.com/vpc/latest/userguide/VPC_Route_Tables.html) in the
	// Amazon Virtual Private Cloud User Guide.
	//
	CreateRouteTable(arg1 context.Context, arg2 *ec2.CreateRouteTableInput, arg3 ...func(*ec2.Options)) (*ec2.CreateRouteTableOutput, error)
	// Creates a security group. A security group acts as a virtual firewall for your
	// instance to control inbound and outbound traffic. For more information, see
	// Amazon EC2 security groups
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/using-network-security.html)
	// in the Amazon Elastic Compute Cloud User Guide and Security groups for your VPC
	// (https://docs.aws.amazon.com/AmazonVPC/latest/UserGuide/VPC_SecurityGroups.html)
	// in the Amazon Virtual Private Cloud User Guide. When you create a security
	// group, you specify a friendly name of your choice. You can have a security group
	// for use in EC2-Classic with the same name as a security group for use in a VPC.
	// However, you can't have two security groups for use in EC2-Classic with the same
	// name or two security groups for use in a VPC with the same name. You have a
	// default security group for use in EC2-Classic and a default security group for
	// use in your VPC. If you don't specify a security group when you launch an
	// instance, the instance is launched into the appropriate default security group.
	// A default security group includes a default rule that grants instances
	// unrestricted network access to each other. You can add or remove rules from your
	// security groups using AuthorizeSecurityGroupIngress,
	// AuthorizeSecurityGroupEgress, RevokeSecurityGroupIngress, and
	// RevokeSecurityGroupEgress. For more information about VPC security group limits,
	// see Amazon VPC Limits
	// (https://docs.aws.amazon.com/vpc/latest/userguide/amazon-vpc-limits.html).
	//
	CreateSecurityGroup(arg1 context.Context, arg2 *ec2.CreateSecurityGroupInput, arg3 ...func(*ec2.Options)) (*ec2.CreateSecurityGroupOutput, error)
	// Creates a snapshot of an EBS volume and stores it in Amazon S3. You can use
	// snapshots for backups, to make copies of EBS volumes, and to save data before
	// shutting down an instance. You can create snapshots of volumes in a Region and
	// volumes on an Outpost. If you create a snapshot of a volume in a Region, the
	// snapshot must be stored in the same Region as the volume. If you create a
	// snapshot of a volume on an Outpost, the snapshot can be stored on the same
	// Outpost as the volume, or in the Region for that Outpost. When a snapshot is
	// created, any Amazon Web Services Marketplace product codes that are associated
	// with the source volume are propagated to the snapshot. You can take a snapshot
	// of an attached volume that is in use. However, snapshots only capture data that
	// has been written to your Amazon EBS volume at the time the snapshot command is
	// issued; this might exclude any data that has been cached by any applications or
	// the operating system. If you can pause any file systems on the volume long
	// enough to take a snapshot, your snapshot should be complete. However, if you
	// cannot pause all file writes to the volume, you should unmount the volume from
	// within the instance, issue the snapshot command, and then remount the volume to
	// ensure a consistent and complete snapshot. You may remount and use your volume
	// while the snapshot status is pending. To create a snapshot for Amazon EBS
	// volumes that serve as root devices, you should stop the instance before taking
	// the snapshot. Snapshots that are taken from encrypted volumes are automatically
	// encrypted. Volumes that are created from encrypted snapshots are also
	// automatically encrypted. Your encrypted volumes and any associated snapshots
	// always remain protected. You can tag your snapshots during creation. For more
	// information, see Tag your Amazon EC2 resources
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/Using_Tags.html) in the
	// Amazon Elastic Compute Cloud User Guide. For more information, see Amazon
	// Elastic Block Store
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/AmazonEBS.html) and Amazon
	// EBS encryption
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSEncryption.html) in the
	// Amazon Elastic Compute Cloud User Guide.
	//
	CreateSnapshot(arg1 context.Context, arg2 *ec2.CreateSnapshotInput, arg3 ...func(*ec2.Options)) (*ec2.CreateSnapshotOutput, error)
	// Creates crash-consistent snapshots of multiple EBS volumes and stores the data
	// in S3. Volumes are chosen by specifying an instance. Any attached volumes will
	// produce one snapshot each that is crash-consistent across the instance. Boot
	// volumes can be excluded by changing the parameters. You can create multi-volume
	// snapshots of instances in a Region and instances on an Outpost. If you create
	// snapshots from an instance in a Region, the snapshots must be stored in the same
	// Region as the instance. If you create snapshots from an instance on an Outpost,
	// the snapshots can be stored on the same Outpost as the instance, or in the
	// Region for that Outpost.
	//
	CreateSnapshots(arg1 context.Context, arg2 *ec2.CreateSnapshotsInput, arg3 ...func(*ec2.Options)) (*ec2.CreateSnapshotsOutput, error)
	// Creates a data feed for Spot Instances, enabling you to view Spot Instance usage
	// logs. You can create one data feed per Amazon Web Services account. For more
	// information, see Spot Instance data feed
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/spot-data-feeds.html) in
	// the Amazon EC2 User Guide for Linux Instances.
	//
	CreateSpotDatafeedSubscription(arg1 context.Context, arg2 *ec2.CreateSpotDatafeedSubscriptionInput, arg3 ...func(*ec2.Options)) (*ec2.CreateSpotDatafeedSubscriptionOutput, error)
	// Stores an AMI as a single object in an Amazon S3 bucket. To use this API, you
	// must have the required permissions. For more information, see Permissions for
	// storing and restoring AMIs using Amazon S3
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ami-store-restore.html#ami-s3-permissions)
	// in the Amazon Elastic Compute Cloud User Guide. For more information, see Store
	// and restore an AMI using Amazon S3
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ami-store-restore.html) in
	// the Amazon Elastic Compute Cloud User Guide.
	//
	CreateStoreImageTask(arg1 context.Context, arg2 *ec2.CreateStoreImageTaskInput, arg3 ...func(*ec2.Options)) (*ec2.CreateStoreImageTaskOutput, error)
	// Creates a subnet in a specified VPC. You must specify an IPv4 CIDR block for the
	// subnet. After you create a subnet, you can't change its CIDR block. The allowed
	// block size is between a /16 netmask (65,536 IP addresses) and /28 netmask (16 IP
	// addresses). The CIDR block must not overlap with the CIDR block of an existing
	// subnet in the VPC. If you've associated an IPv6 CIDR block with your VPC, you
	// can create a subnet with an IPv6 CIDR block that uses a /64 prefix length.
	// Amazon Web Services reserves both the first four and the last IPv4 address in
	// each subnet's CIDR block. They're not available for use. If you add more than
	// one subnet to a VPC, they're set up in a star topology with a logical router in
	// the middle. When you stop an instance in a subnet, it retains its private IPv4
	// address. It's therefore possible to have a subnet with no running instances
	// (they're all stopped), but no remaining IP addresses available. For more
	// information about subnets, see Your VPC and subnets
	// (https://docs.aws.amazon.com/vpc/latest/userguide/VPC_Subnets.html) in the
	// Amazon Virtual Private Cloud User Guide.
	//
	CreateSubnet(arg1 context.Context, arg2 *ec2.CreateSubnetInput, arg3 ...func(*ec2.Options)) (*ec2.CreateSubnetOutput, error)
	// Creates a subnet CIDR reservation. For information about subnet CIDR
	// reservations, see Subnet CIDR reservations
	// (https://docs.aws.amazon.com/vpc/latest/userguide/subnet-cidr-reservation.html)
	// in the Amazon Virtual Private Cloud User Guide.
	//
	CreateSubnetCidrReservation(arg1 context.Context, arg2 *ec2.CreateSubnetCidrReservationInput, arg3 ...func(*ec2.Options)) (*ec2.CreateSubnetCidrReservationOutput, error)
	// Adds or overwrites only the specified tags for the specified Amazon EC2 resource
	// or resources. When you specify an existing tag key, the value is overwritten
	// with the new value. Each resource can have a maximum of 50 tags. Each tag
	// consists of a key and optional value. Tag keys must be unique per resource. For
	// more information about tags, see Tag your Amazon EC2 resources
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/Using_Tags.html) in the
	// Amazon Elastic Compute Cloud User Guide. For more information about creating IAM
	// policies that control users' access to resources based on tags, see Supported
	// resource-level permissions for Amazon EC2 API actions
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-supported-iam-actions-resources.html)
	// in the Amazon Elastic Compute Cloud User Guide.
	//
	CreateTags(arg1 context.Context, arg2 *ec2.CreateTagsInput, arg3 ...func(*ec2.Options)) (*ec2.CreateTagsOutput, error)
	// Creates a Traffic Mirror filter. A Traffic Mirror filter is a set of rules that
	// defines the traffic to mirror. By default, no traffic is mirrored. To mirror
	// traffic, use CreateTrafficMirrorFilterRule
	// (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_CreateTrafficMirrorFilterRule.htm)
	// to add Traffic Mirror rules to the filter. The rules you add define what traffic
	// gets mirrored. You can also use ModifyTrafficMirrorFilterNetworkServices
	// (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_ModifyTrafficMirrorFilterNetworkServices.html)
	// to mirror supported network services.
	//
	CreateTrafficMirrorFilter(arg1 context.Context, arg2 *ec2.CreateTrafficMirrorFilterInput, arg3 ...func(*ec2.Options)) (*ec2.CreateTrafficMirrorFilterOutput, error)
	// Creates a Traffic Mirror filter rule. A Traffic Mirror rule defines the Traffic
	// Mirror source traffic to mirror. You need the Traffic Mirror filter ID when you
	// create the rule.
	//
	CreateTrafficMirrorFilterRule(arg1 context.Context, arg2 *ec2.CreateTrafficMirrorFilterRuleInput, arg3 ...func(*ec2.Options)) (*ec2.CreateTrafficMirrorFilterRuleOutput, error)
	// Creates a Traffic Mirror session. A Traffic Mirror session actively copies
	// packets from a Traffic Mirror source to a Traffic Mirror target. Create a
	// filter, and then assign it to the session to define a subset of the traffic to
	// mirror, for example all TCP traffic. The Traffic Mirror source and the Traffic
	// Mirror target (monitoring appliances) can be in the same VPC, or in a different
	// VPC connected via VPC peering or a transit gateway. By default, no traffic is
	// mirrored. Use CreateTrafficMirrorFilter
	// (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_CreateTrafficMirrorFilter.htm)
	// to create filter rules that specify the traffic to mirror.
	//
	CreateTrafficMirrorSession(arg1 context.Context, arg2 *ec2.CreateTrafficMirrorSessionInput, arg3 ...func(*ec2.Options)) (*ec2.CreateTrafficMirrorSessionOutput, error)
	// Creates a target for your Traffic Mirror session. A Traffic Mirror target is the
	// destination for mirrored traffic. The Traffic Mirror source and the Traffic
	// Mirror target (monitoring appliances) can be in the same VPC, or in different
	// VPCs connected via VPC peering or a transit gateway. A Traffic Mirror target can
	// be a network interface, a Network Load Balancer, or a Gateway Load Balancer
	// endpoint. To use the target in a Traffic Mirror session, use
	// CreateTrafficMirrorSession
	// (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_CreateTrafficMirrorSession.htm).
	//
	CreateTrafficMirrorTarget(arg1 context.Context, arg2 *ec2.CreateTrafficMirrorTargetInput, arg3 ...func(*ec2.Options)) (*ec2.CreateTrafficMirrorTargetOutput, error)
	// Creates a transit gateway. You can use a transit gateway to interconnect your
	// virtual private clouds (VPC) and on-premises networks. After the transit gateway
	// enters the available state, you can attach your VPCs and VPN connections to the
	// transit gateway. To attach your VPCs, use CreateTransitGatewayVpcAttachment. To
	// attach a VPN connection, use CreateCustomerGateway to create a customer gateway
	// and specify the ID of the customer gateway and the ID of the transit gateway in
	// a call to CreateVpnConnection. When you create a transit gateway, we create a
	// default transit gateway route table and use it as the default association route
	// table and the default propagation route table. You can use
	// CreateTransitGatewayRouteTable to create additional transit gateway route
	// tables. If you disable automatic route propagation, we do not create a default
	// transit gateway route table. You can use
	// EnableTransitGatewayRouteTablePropagation to propagate routes from a resource
	// attachment to a transit gateway route table. If you disable automatic
	// associations, you can use AssociateTransitGatewayRouteTable to associate a
	// resource attachment with a transit gateway route table.
	//
	CreateTransitGateway(arg1 context.Context, arg2 *ec2.CreateTransitGatewayInput, arg3 ...func(*ec2.Options)) (*ec2.CreateTransitGatewayOutput, error)
	// Creates a Connect attachment from a specified transit gateway attachment. A
	// Connect attachment is a GRE-based tunnel attachment that you can use to
	// establish a connection between a transit gateway and an appliance. A Connect
	// attachment uses an existing VPC or Amazon Web Services Direct Connect attachment
	// as the underlying transport mechanism.
	//
	CreateTransitGatewayConnect(arg1 context.Context, arg2 *ec2.CreateTransitGatewayConnectInput, arg3 ...func(*ec2.Options)) (*ec2.CreateTransitGatewayConnectOutput, error)
	// Creates a Connect peer for a specified transit gateway Connect attachment
	// between a transit gateway and an appliance. The peer address and transit gateway
	// address must be the same IP address family (IPv4 or IPv6). For more information,
	// see Connect peers
	// (https://docs.aws.amazon.com/vpc/latest/tgw/tgw-connect.html#tgw-connect-peer)
	// in the Transit Gateways Guide.
	//
	CreateTransitGatewayConnectPeer(arg1 context.Context, arg2 *ec2.CreateTransitGatewayConnectPeerInput, arg3 ...func(*ec2.Options)) (*ec2.CreateTransitGatewayConnectPeerOutput, error)
	// Creates a multicast domain using the specified transit gateway. The transit
	// gateway must be in the available state before you create a domain. Use
	// DescribeTransitGateways
	// (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribeTransitGateways.html)
	// to see the state of transit gateway.
	//
	CreateTransitGatewayMulticastDomain(arg1 context.Context, arg2 *ec2.CreateTransitGatewayMulticastDomainInput, arg3 ...func(*ec2.Options)) (*ec2.CreateTransitGatewayMulticastDomainOutput, error)
	// Requests a transit gateway peering attachment between the specified transit
	// gateway (requester) and a peer transit gateway (accepter). The peer transit
	// gateway can be in your account or a different Amazon Web Services account. After
	// you create the peering attachment, the owner of the accepter transit gateway
	// must accept the attachment request.
	//
	CreateTransitGatewayPeeringAttachment(arg1 context.Context, arg2 *ec2.CreateTransitGatewayPeeringAttachmentInput, arg3 ...func(*ec2.Options)) (*ec2.CreateTransitGatewayPeeringAttachmentOutput, error)
	// Creates a transit gateway policy table.
	//
	CreateTransitGatewayPolicyTable(arg1 context.Context, arg2 *ec2.CreateTransitGatewayPolicyTableInput, arg3 ...func(*ec2.Options)) (*ec2.CreateTransitGatewayPolicyTableOutput, error)
	// Creates a reference (route) to a prefix list in a specified transit gateway
	// route table.
	//
	CreateTransitGatewayPrefixListReference(arg1 context.Context, arg2 *ec2.CreateTransitGatewayPrefixListReferenceInput, arg3 ...func(*ec2.Options)) (*ec2.CreateTransitGatewayPrefixListReferenceOutput, error)
	// Creates a static route for the specified transit gateway route table.
	//
	CreateTransitGatewayRoute(arg1 context.Context, arg2 *ec2.CreateTransitGatewayRouteInput, arg3 ...func(*ec2.Options)) (*ec2.CreateTransitGatewayRouteOutput, error)
	// Creates a route table for the specified transit gateway.
	//
	CreateTransitGatewayRouteTable(arg1 context.Context, arg2 *ec2.CreateTransitGatewayRouteTableInput, arg3 ...func(*ec2.Options)) (*ec2.CreateTransitGatewayRouteTableOutput, error)
	// Advertises a new transit gateway route table.
	//
	CreateTransitGatewayRouteTableAnnouncement(arg1 context.Context, arg2 *ec2.CreateTransitGatewayRouteTableAnnouncementInput, arg3 ...func(*ec2.Options)) (*ec2.CreateTransitGatewayRouteTableAnnouncementOutput, error)
	// Attaches the specified VPC to the specified transit gateway. If you attach a VPC
	// with a CIDR range that overlaps the CIDR range of a VPC that is already
	// attached, the new VPC CIDR range is not propagated to the default propagation
	// route table. To send VPC traffic to an attached transit gateway, add a route to
	// the VPC route table using CreateRoute.
	//
	CreateTransitGatewayVpcAttachment(arg1 context.Context, arg2 *ec2.CreateTransitGatewayVpcAttachmentInput, arg3 ...func(*ec2.Options)) (*ec2.CreateTransitGatewayVpcAttachmentOutput, error)
	// Creates an EBS volume that can be attached to an instance in the same
	// Availability Zone. You can create a new empty volume or restore a volume from an
	// EBS snapshot. Any Amazon Web Services Marketplace product codes from the
	// snapshot are propagated to the volume. You can create encrypted volumes.
	// Encrypted volumes must be attached to instances that support Amazon EBS
	// encryption. Volumes that are created from encrypted snapshots are also
	// automatically encrypted. For more information, see Amazon EBS encryption
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSEncryption.html) in the
	// Amazon Elastic Compute Cloud User Guide. You can tag your volumes during
	// creation. For more information, see Tag your Amazon EC2 resources
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/Using_Tags.html) in the
	// Amazon Elastic Compute Cloud User Guide. For more information, see Create an
	// Amazon EBS volume
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ebs-creating-volume.html)
	// in the Amazon Elastic Compute Cloud User Guide.
	//
	CreateVolume(arg1 context.Context, arg2 *ec2.CreateVolumeInput, arg3 ...func(*ec2.Options)) (*ec2.CreateVolumeOutput, error)
	// Creates a VPC with the specified IPv4 CIDR block. The smallest VPC you can
	// create uses a /28 netmask (16 IPv4 addresses), and the largest uses a /16
	// netmask (65,536 IPv4 addresses). For more information about how large to make
	// your VPC, see Your VPC and subnets
	// (https://docs.aws.amazon.com/vpc/latest/userguide/VPC_Subnets.html) in the
	// Amazon Virtual Private Cloud User Guide. You can optionally request an IPv6 CIDR
	// block for the VPC. You can request an Amazon-provided IPv6 CIDR block from
	// Amazon's pool of IPv6 addresses, or an IPv6 CIDR block from an IPv6 address pool
	// that you provisioned through bring your own IP addresses (BYOIP
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-byoip.html)). By
	// default, each instance you launch in the VPC has the default DHCP options, which
	// include only a default DNS server that we provide (AmazonProvidedDNS). For more
	// information, see DHCP options sets
	// (https://docs.aws.amazon.com/vpc/latest/userguide/VPC_DHCP_Options.html) in the
	// Amazon Virtual Private Cloud User Guide. You can specify the instance tenancy
	// value for the VPC when you create it. You can't change this value for the VPC
	// after you create it. For more information, see Dedicated Instances
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/dedicated-instance.html) in
	// the Amazon Elastic Compute Cloud User Guide.
	//
	CreateVpc(arg1 context.Context, arg2 *ec2.CreateVpcInput, arg3 ...func(*ec2.Options)) (*ec2.CreateVpcOutput, error)
	// Creates a VPC endpoint for a specified service. An endpoint enables you to
	// create a private connection between your VPC and the service. The service may be
	// provided by Amazon Web Services, an Amazon Web Services Marketplace Partner, or
	// another Amazon Web Services account. For more information, see the Amazon Web
	// Services PrivateLink Guide
	// (https://docs.aws.amazon.com/vpc/latest/privatelink/).
	//
	CreateVpcEndpoint(arg1 context.Context, arg2 *ec2.CreateVpcEndpointInput, arg3 ...func(*ec2.Options)) (*ec2.CreateVpcEndpointOutput, error)
	// Creates a connection notification for a specified VPC endpoint or VPC endpoint
	// service. A connection notification notifies you of specific endpoint events. You
	// must create an SNS topic to receive notifications. For more information, see
	// Create a Topic (https://docs.aws.amazon.com/sns/latest/dg/CreateTopic.html) in
	// the Amazon Simple Notification Service Developer Guide. You can create a
	// connection notification for interface endpoints only.
	//
	CreateVpcEndpointConnectionNotification(arg1 context.Context, arg2 *ec2.CreateVpcEndpointConnectionNotificationInput, arg3 ...func(*ec2.Options)) (*ec2.CreateVpcEndpointConnectionNotificationOutput, error)
	// Creates a VPC endpoint service to which service consumers (Amazon Web Services
	// accounts, IAM users, and IAM roles) can connect. Before you create an endpoint
	// service, you must create one of the following for your service:
	//
	// * A Network
	// Load Balancer
	// (https://docs.aws.amazon.com/elasticloadbalancing/latest/network/). Service
	// consumers connect to your service using an interface endpoint.
	//
	// * A Gateway Load
	// Balancer (https://docs.aws.amazon.com/elasticloadbalancing/latest/gateway/).
	// Service consumers connect to your service using a Gateway Load Balancer
	// endpoint.
	//
	// If you set the private DNS name, you must prove that you own the
	// private DNS domain name. For more information, see the Amazon Web Services
	// PrivateLink Guide (https://docs.aws.amazon.com/vpc/latest/privatelink/).
	//
	CreateVpcEndpointServiceConfiguration(arg1 context.Context, arg2 *ec2.CreateVpcEndpointServiceConfigurationInput, arg3 ...func(*ec2.Options)) (*ec2.CreateVpcEndpointServiceConfigurationOutput, error)
	// Requests a VPC peering connection between two VPCs: a requester VPC that you own
	// and an accepter VPC with which to create the connection. The accepter VPC can
	// belong to another Amazon Web Services account and can be in a different Region
	// to the requester VPC. The requester VPC and accepter VPC cannot have overlapping
	// CIDR blocks. Limitations and rules apply to a VPC peering connection. For more
	// information, see the limitations
	// (https://docs.aws.amazon.com/vpc/latest/peering/vpc-peering-basics.html#vpc-peering-limitations)
	// section in the VPC Peering Guide. The owner of the accepter VPC must accept the
	// peering request to activate the peering connection. The VPC peering connection
	// request expires after 7 days, after which it cannot be accepted or rejected. If
	// you create a VPC peering connection request between VPCs with overlapping CIDR
	// blocks, the VPC peering connection has a status of failed.
	//
	CreateVpcPeeringConnection(arg1 context.Context, arg2 *ec2.CreateVpcPeeringConnectionInput, arg3 ...func(*ec2.Options)) (*ec2.CreateVpcPeeringConnectionOutput, error)
	// Creates a VPN connection between an existing virtual private gateway or transit
	// gateway and a customer gateway. The supported connection type is ipsec.1. The
	// response includes information that you need to give to your network
	// administrator to configure your customer gateway. We strongly recommend that you
	// use HTTPS when calling this operation because the response contains sensitive
	// cryptographic information for configuring your customer gateway device. If you
	// decide to shut down your VPN connection for any reason and later create a new
	// VPN connection, you must reconfigure your customer gateway with the new
	// information returned from this call. This is an idempotent operation. If you
	// perform the operation more than once, Amazon EC2 doesn't return an error. For
	// more information, see Amazon Web Services Site-to-Site VPN
	// (https://docs.aws.amazon.com/vpn/latest/s2svpn/VPC_VPN.html) in the Amazon Web
	// Services Site-to-Site VPN User Guide.
	//
	CreateVpnConnection(arg1 context.Context, arg2 *ec2.CreateVpnConnectionInput, arg3 ...func(*ec2.Options)) (*ec2.CreateVpnConnectionOutput, error)
	// Creates a static route associated with a VPN connection between an existing
	// virtual private gateway and a VPN customer gateway. The static route allows
	// traffic to be routed from the virtual private gateway to the VPN customer
	// gateway. For more information, see Amazon Web Services Site-to-Site VPN
	// (https://docs.aws.amazon.com/vpn/latest/s2svpn/VPC_VPN.html) in the Amazon Web
	// Services Site-to-Site VPN User Guide.
	//
	CreateVpnConnectionRoute(arg1 context.Context, arg2 *ec2.CreateVpnConnectionRouteInput, arg3 ...func(*ec2.Options)) (*ec2.CreateVpnConnectionRouteOutput, error)
	// Creates a virtual private gateway. A virtual private gateway is the endpoint on
	// the VPC side of your VPN connection. You can create a virtual private gateway
	// before creating the VPC itself. For more information, see Amazon Web Services
	// Site-to-Site VPN (https://docs.aws.amazon.com/vpn/latest/s2svpn/VPC_VPN.html) in
	// the Amazon Web Services Site-to-Site VPN User Guide.
	//
	CreateVpnGateway(arg1 context.Context, arg2 *ec2.CreateVpnGatewayInput, arg3 ...func(*ec2.Options)) (*ec2.CreateVpnGatewayOutput, error)
	// Deletes a carrier gateway. If you do not delete the route that contains the
	// carrier gateway as the Target, the route is a blackhole route. For information
	// about how to delete a route, see DeleteRoute
	// (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DeleteRoute.html).
	//
	DeleteCarrierGateway(arg1 context.Context, arg2 *ec2.DeleteCarrierGatewayInput, arg3 ...func(*ec2.Options)) (*ec2.DeleteCarrierGatewayOutput, error)
	// Deletes the specified Client VPN endpoint. You must disassociate all target
	// networks before you can delete a Client VPN endpoint.
	//
	DeleteClientVpnEndpoint(arg1 context.Context, arg2 *ec2.DeleteClientVpnEndpointInput, arg3 ...func(*ec2.Options)) (*ec2.DeleteClientVpnEndpointOutput, error)
	// Deletes a route from a Client VPN endpoint. You can only delete routes that you
	// manually added using the CreateClientVpnRoute action. You cannot delete routes
	// that were automatically added when associating a subnet. To remove routes that
	// have been automatically added, disassociate the target subnet from the Client
	// VPN endpoint.
	//
	DeleteClientVpnRoute(arg1 context.Context, arg2 *ec2.DeleteClientVpnRouteInput, arg3 ...func(*ec2.Options)) (*ec2.DeleteClientVpnRouteOutput, error)
	// Deletes the specified customer gateway. You must delete the VPN connection
	// before you can delete the customer gateway.
	//
	DeleteCustomerGateway(arg1 context.Context, arg2 *ec2.DeleteCustomerGatewayInput, arg3 ...func(*ec2.Options)) (*ec2.DeleteCustomerGatewayOutput, error)
	// Deletes the specified set of DHCP options. You must disassociate the set of DHCP
	// options before you can delete it. You can disassociate the set of DHCP options
	// by associating either a new set of options or the default set of options with
	// the VPC.
	//
	DeleteDhcpOptions(arg1 context.Context, arg2 *ec2.DeleteDhcpOptionsInput, arg3 ...func(*ec2.Options)) (*ec2.DeleteDhcpOptionsOutput, error)
	// Deletes an egress-only internet gateway.
	//
	DeleteEgressOnlyInternetGateway(arg1 context.Context, arg2 *ec2.DeleteEgressOnlyInternetGatewayInput, arg3 ...func(*ec2.Options)) (*ec2.DeleteEgressOnlyInternetGatewayOutput, error)
	// Deletes the specified EC2 Fleet. After you delete an EC2 Fleet, it launches no
	// new instances. You must specify whether a deleted EC2 Fleet should also
	// terminate its instances. If you choose to terminate the instances, the EC2 Fleet
	// enters the deleted_terminating state. Otherwise, the EC2 Fleet enters the
	// deleted_running state, and the instances continue to run until they are
	// interrupted or you terminate them manually. For instant fleets, EC2 Fleet must
	// terminate the instances when the fleet is deleted. A deleted instant fleet with
	// running instances is not supported. Restrictions
	//
	// * You can delete up to 25
	// instant fleets in a single request. If you exceed this number, no instant fleets
	// are deleted and an error is returned. There is no restriction on the number of
	// fleets of type maintain or request that can be deleted in a single request.
	//
	// *
	// Up to 1000 instances can be terminated in a single request to delete instant
	// fleets.
	//
	// For more information, see Delete an EC2 Fleet
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/manage-ec2-fleet.html#delete-fleet)
	// in the Amazon EC2 User Guide.
	//
	DeleteFleets(arg1 context.Context, arg2 *ec2.DeleteFleetsInput, arg3 ...func(*ec2.Options)) (*ec2.DeleteFleetsOutput, error)
	// Deletes one or more flow logs.
	//
	DeleteFlowLogs(arg1 context.Context, arg2 *ec2.DeleteFlowLogsInput, arg3 ...func(*ec2.Options)) (*ec2.DeleteFlowLogsOutput, error)
	// Deletes the specified Amazon FPGA Image (AFI).
	//
	DeleteFpgaImage(arg1 context.Context, arg2 *ec2.DeleteFpgaImageInput, arg3 ...func(*ec2.Options)) (*ec2.DeleteFpgaImageOutput, error)
	// Deletes the specified event window. For more information, see Define event
	// windows for scheduled events
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/event-windows.html) in the
	// Amazon EC2 User Guide.
	//
	DeleteInstanceEventWindow(arg1 context.Context, arg2 *ec2.DeleteInstanceEventWindowInput, arg3 ...func(*ec2.Options)) (*ec2.DeleteInstanceEventWindowOutput, error)
	// Deletes the specified internet gateway. You must detach the internet gateway
	// from the VPC before you can delete it.
	//
	DeleteInternetGateway(arg1 context.Context, arg2 *ec2.DeleteInternetGatewayInput, arg3 ...func(*ec2.Options)) (*ec2.DeleteInternetGatewayOutput, error)
	// Delete an IPAM. Deleting an IPAM removes all monitored data associated with the
	// IPAM including the historical data for CIDRs. For more information, see Delete
	// an IPAM (https://docs.aws.amazon.com/vpc/latest/ipam/delete-ipam.html) in the
	// Amazon VPC IPAM User Guide.
	//
	DeleteIpam(arg1 context.Context, arg2 *ec2.DeleteIpamInput, arg3 ...func(*ec2.Options)) (*ec2.DeleteIpamOutput, error)
	// Delete an IPAM pool. You cannot delete an IPAM pool if there are allocations in
	// it or CIDRs provisioned to it. To release allocations, see
	// ReleaseIpamPoolAllocation
	// (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_ReleaseIpamPoolAllocation.html).
	// To deprovision pool CIDRs, see DeprovisionIpamPoolCidr
	// (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DeprovisionIpamPoolCidr.html).
	// For more information, see Delete a pool
	// (https://docs.aws.amazon.com/vpc/latest/ipam/delete-pool-ipam.html) in the
	// Amazon VPC IPAM User Guide.
	//
	DeleteIpamPool(arg1 context.Context, arg2 *ec2.DeleteIpamPoolInput, arg3 ...func(*ec2.Options)) (*ec2.DeleteIpamPoolOutput, error)
	// Delete the scope for an IPAM. You cannot delete the default scopes. For more
	// information, see Delete a scope
	// (https://docs.aws.amazon.com/vpc/latest/ipam/delete-scope-ipam.html) in the
	// Amazon VPC IPAM User Guide.
	//
	DeleteIpamScope(arg1 context.Context, arg2 *ec2.DeleteIpamScopeInput, arg3 ...func(*ec2.Options)) (*ec2.DeleteIpamScopeOutput, error)
	// Deletes the specified key pair, by removing the public key from Amazon EC2.
	//
	DeleteKeyPair(arg1 context.Context, arg2 *ec2.DeleteKeyPairInput, arg3 ...func(*ec2.Options)) (*ec2.DeleteKeyPairOutput, error)
	// Deletes a launch template. Deleting a launch template deletes all of its
	// versions.
	//
	DeleteLaunchTemplate(arg1 context.Context, arg2 *ec2.DeleteLaunchTemplateInput, arg3 ...func(*ec2.Options)) (*ec2.DeleteLaunchTemplateOutput, error)
	// Deletes one or more versions of a launch template. You cannot delete the default
	// version of a launch template; you must first assign a different version as the
	// default. If the default version is the only version for the launch template, you
	// must delete the entire launch template using DeleteLaunchTemplate.
	//
	DeleteLaunchTemplateVersions(arg1 context.Context, arg2 *ec2.DeleteLaunchTemplateVersionsInput, arg3 ...func(*ec2.Options)) (*ec2.DeleteLaunchTemplateVersionsOutput, error)
	// Deletes the specified route from the specified local gateway route table.
	//
	DeleteLocalGatewayRoute(arg1 context.Context, arg2 *ec2.DeleteLocalGatewayRouteInput, arg3 ...func(*ec2.Options)) (*ec2.DeleteLocalGatewayRouteOutput, error)
	// Deletes the specified association between a VPC and local gateway route table.
	//
	DeleteLocalGatewayRouteTableVpcAssociation(arg1 context.Context, arg2 *ec2.DeleteLocalGatewayRouteTableVpcAssociationInput, arg3 ...func(*ec2.Options)) (*ec2.DeleteLocalGatewayRouteTableVpcAssociationOutput, error)
	// Deletes the specified managed prefix list. You must first remove all references
	// to the prefix list in your resources.
	//
	DeleteManagedPrefixList(arg1 context.Context, arg2 *ec2.DeleteManagedPrefixListInput, arg3 ...func(*ec2.Options)) (*ec2.DeleteManagedPrefixListOutput, error)
	// Deletes the specified NAT gateway. Deleting a public NAT gateway disassociates
	// its Elastic IP address, but does not release the address from your account.
	// Deleting a NAT gateway does not delete any NAT gateway routes in your route
	// tables.
	//
	DeleteNatGateway(arg1 context.Context, arg2 *ec2.DeleteNatGatewayInput, arg3 ...func(*ec2.Options)) (*ec2.DeleteNatGatewayOutput, error)
	// Deletes the specified network ACL. You can't delete the ACL if it's associated
	// with any subnets. You can't delete the default network ACL.
	//
	DeleteNetworkAcl(arg1 context.Context, arg2 *ec2.DeleteNetworkAclInput, arg3 ...func(*ec2.Options)) (*ec2.DeleteNetworkAclOutput, error)
	// Deletes the specified ingress or egress entry (rule) from the specified network
	// ACL.
	//
	DeleteNetworkAclEntry(arg1 context.Context, arg2 *ec2.DeleteNetworkAclEntryInput, arg3 ...func(*ec2.Options)) (*ec2.DeleteNetworkAclEntryOutput, error)
	// Deletes the specified Network Access Scope.
	//
	DeleteNetworkInsightsAccessScope(arg1 context.Context, arg2 *ec2.DeleteNetworkInsightsAccessScopeInput, arg3 ...func(*ec2.Options)) (*ec2.DeleteNetworkInsightsAccessScopeOutput, error)
	// Deletes the specified Network Access Scope analysis.
	//
	DeleteNetworkInsightsAccessScopeAnalysis(arg1 context.Context, arg2 *ec2.DeleteNetworkInsightsAccessScopeAnalysisInput, arg3 ...func(*ec2.Options)) (*ec2.DeleteNetworkInsightsAccessScopeAnalysisOutput, error)
	// Deletes the specified network insights analysis.
	//
	DeleteNetworkInsightsAnalysis(arg1 context.Context, arg2 *ec2.DeleteNetworkInsightsAnalysisInput, arg3 ...func(*ec2.Options)) (*ec2.DeleteNetworkInsightsAnalysisOutput, error)
	// Deletes the specified path.
	//
	DeleteNetworkInsightsPath(arg1 context.Context, arg2 *ec2.DeleteNetworkInsightsPathInput, arg3 ...func(*ec2.Options)) (*ec2.DeleteNetworkInsightsPathOutput, error)
	// Deletes the specified network interface. You must detach the network interface
	// before you can delete it.
	//
	DeleteNetworkInterface(arg1 context.Context, arg2 *ec2.DeleteNetworkInterfaceInput, arg3 ...func(*ec2.Options)) (*ec2.DeleteNetworkInterfaceOutput, error)
	// Deletes a permission for a network interface. By default, you cannot delete the
	// permission if the account for which you're removing the permission has attached
	// the network interface to an instance. However, you can force delete the
	// permission, regardless of any attachment.
	//
	DeleteNetworkInterfacePermission(arg1 context.Context, arg2 *ec2.DeleteNetworkInterfacePermissionInput, arg3 ...func(*ec2.Options)) (*ec2.DeleteNetworkInterfacePermissionOutput, error)
	// Deletes the specified placement group. You must terminate all instances in the
	// placement group before you can delete the placement group. For more information,
	// see Placement groups
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/placement-groups.html) in
	// the Amazon EC2 User Guide.
	//
	DeletePlacementGroup(arg1 context.Context, arg2 *ec2.DeletePlacementGroupInput, arg3 ...func(*ec2.Options)) (*ec2.DeletePlacementGroupOutput, error)
	// Delete a public IPv4 pool. A public IPv4 pool is an EC2 IP address pool required
	// for the public IPv4 CIDRs that you own and bring to Amazon Web Services to
	// manage with IPAM. IPv6 addresses you bring to Amazon Web Services, however, use
	// IPAM pools only.
	//
	DeletePublicIpv4Pool(arg1 context.Context, arg2 *ec2.DeletePublicIpv4PoolInput, arg3 ...func(*ec2.Options)) (*ec2.DeletePublicIpv4PoolOutput, error)
	// Deletes the queued purchases for the specified Reserved Instances.
	//
	DeleteQueuedReservedInstances(arg1 context.Context, arg2 *ec2.DeleteQueuedReservedInstancesInput, arg3 ...func(*ec2.Options)) (*ec2.DeleteQueuedReservedInstancesOutput, error)
	// Deletes the specified route from the specified route table.
	//
	DeleteRoute(arg1 context.Context, arg2 *ec2.DeleteRouteInput, arg3 ...func(*ec2.Options)) (*ec2.DeleteRouteOutput, error)
	// Deletes the specified route table. You must disassociate the route table from
	// any subnets before you can delete it. You can't delete the main route table.
	//
	DeleteRouteTable(arg1 context.Context, arg2 *ec2.DeleteRouteTableInput, arg3 ...func(*ec2.Options)) (*ec2.DeleteRouteTableOutput, error)
	// Deletes a security group. If you attempt to delete a security group that is
	// associated with an instance, or is referenced by another security group, the
	// operation fails with InvalidGroup.InUse in EC2-Classic or DependencyViolation in
	// EC2-VPC.
	//
	DeleteSecurityGroup(arg1 context.Context, arg2 *ec2.DeleteSecurityGroupInput, arg3 ...func(*ec2.Options)) (*ec2.DeleteSecurityGroupOutput, error)
	// Deletes the specified snapshot. When you make periodic snapshots of a volume,
	// the snapshots are incremental, and only the blocks on the device that have
	// changed since your last snapshot are saved in the new snapshot. When you delete
	// a snapshot, only the data not needed for any other snapshot is removed. So
	// regardless of which prior snapshots have been deleted, all active snapshots will
	// have access to all the information needed to restore the volume. You cannot
	// delete a snapshot of the root device of an EBS volume used by a registered AMI.
	// You must first de-register the AMI before you can delete the snapshot. For more
	// information, see Delete an Amazon EBS snapshot
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ebs-deleting-snapshot.html)
	// in the Amazon Elastic Compute Cloud User Guide.
	//
	DeleteSnapshot(arg1 context.Context, arg2 *ec2.DeleteSnapshotInput, arg3 ...func(*ec2.Options)) (*ec2.DeleteSnapshotOutput, error)
	// Deletes the data feed for Spot Instances.
	//
	DeleteSpotDatafeedSubscription(arg1 context.Context, arg2 *ec2.DeleteSpotDatafeedSubscriptionInput, arg3 ...func(*ec2.Options)) (*ec2.DeleteSpotDatafeedSubscriptionOutput, error)
	// Deletes the specified subnet. You must terminate all running instances in the
	// subnet before you can delete the subnet.
	//
	DeleteSubnet(arg1 context.Context, arg2 *ec2.DeleteSubnetInput, arg3 ...func(*ec2.Options)) (*ec2.DeleteSubnetOutput, error)
	// Deletes a subnet CIDR reservation.
	//
	DeleteSubnetCidrReservation(arg1 context.Context, arg2 *ec2.DeleteSubnetCidrReservationInput, arg3 ...func(*ec2.Options)) (*ec2.DeleteSubnetCidrReservationOutput, error)
	// Deletes the specified set of tags from the specified set of resources. To list
	// the current tags, use DescribeTags. For more information about tags, see Tag
	// your Amazon EC2 resources
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/Using_Tags.html) in the
	// Amazon Elastic Compute Cloud User Guide.
	//
	DeleteTags(arg1 context.Context, arg2 *ec2.DeleteTagsInput, arg3 ...func(*ec2.Options)) (*ec2.DeleteTagsOutput, error)
	// Deletes the specified Traffic Mirror filter. You cannot delete a Traffic Mirror
	// filter that is in use by a Traffic Mirror session.
	//
	DeleteTrafficMirrorFilter(arg1 context.Context, arg2 *ec2.DeleteTrafficMirrorFilterInput, arg3 ...func(*ec2.Options)) (*ec2.DeleteTrafficMirrorFilterOutput, error)
	// Deletes the specified Traffic Mirror rule.
	//
	DeleteTrafficMirrorFilterRule(arg1 context.Context, arg2 *ec2.DeleteTrafficMirrorFilterRuleInput, arg3 ...func(*ec2.Options)) (*ec2.DeleteTrafficMirrorFilterRuleOutput, error)
	// Deletes the specified Traffic Mirror session.
	//
	DeleteTrafficMirrorSession(arg1 context.Context, arg2 *ec2.DeleteTrafficMirrorSessionInput, arg3 ...func(*ec2.Options)) (*ec2.DeleteTrafficMirrorSessionOutput, error)
	// Deletes the specified Traffic Mirror target. You cannot delete a Traffic Mirror
	// target that is in use by a Traffic Mirror session.
	//
	DeleteTrafficMirrorTarget(arg1 context.Context, arg2 *ec2.DeleteTrafficMirrorTargetInput, arg3 ...func(*ec2.Options)) (*ec2.DeleteTrafficMirrorTargetOutput, error)
	// Deletes the specified transit gateway.
	//
	DeleteTransitGateway(arg1 context.Context, arg2 *ec2.DeleteTransitGatewayInput, arg3 ...func(*ec2.Options)) (*ec2.DeleteTransitGatewayOutput, error)
	// Deletes the specified Connect attachment. You must first delete any Connect
	// peers for the attachment.
	//
	DeleteTransitGatewayConnect(arg1 context.Context, arg2 *ec2.DeleteTransitGatewayConnectInput, arg3 ...func(*ec2.Options)) (*ec2.DeleteTransitGatewayConnectOutput, error)
	// Deletes the specified Connect peer.
	//
	DeleteTransitGatewayConnectPeer(arg1 context.Context, arg2 *ec2.DeleteTransitGatewayConnectPeerInput, arg3 ...func(*ec2.Options)) (*ec2.DeleteTransitGatewayConnectPeerOutput, error)
	// Deletes the specified transit gateway multicast domain.
	//
	DeleteTransitGatewayMulticastDomain(arg1 context.Context, arg2 *ec2.DeleteTransitGatewayMulticastDomainInput, arg3 ...func(*ec2.Options)) (*ec2.DeleteTransitGatewayMulticastDomainOutput, error)
	// Deletes a transit gateway peering attachment.
	//
	DeleteTransitGatewayPeeringAttachment(arg1 context.Context, arg2 *ec2.DeleteTransitGatewayPeeringAttachmentInput, arg3 ...func(*ec2.Options)) (*ec2.DeleteTransitGatewayPeeringAttachmentOutput, error)
	// Deletes the specified transit gateway policy table.
	//
	DeleteTransitGatewayPolicyTable(arg1 context.Context, arg2 *ec2.DeleteTransitGatewayPolicyTableInput, arg3 ...func(*ec2.Options)) (*ec2.DeleteTransitGatewayPolicyTableOutput, error)
	// Deletes a reference (route) to a prefix list in a specified transit gateway
	// route table.
	//
	DeleteTransitGatewayPrefixListReference(arg1 context.Context, arg2 *ec2.DeleteTransitGatewayPrefixListReferenceInput, arg3 ...func(*ec2.Options)) (*ec2.DeleteTransitGatewayPrefixListReferenceOutput, error)
	// Deletes the specified route from the specified transit gateway route table.
	//
	DeleteTransitGatewayRoute(arg1 context.Context, arg2 *ec2.DeleteTransitGatewayRouteInput, arg3 ...func(*ec2.Options)) (*ec2.DeleteTransitGatewayRouteOutput, error)
	// Deletes the specified transit gateway route table. You must disassociate the
	// route table from any transit gateway route tables before you can delete it.
	//
	DeleteTransitGatewayRouteTable(arg1 context.Context, arg2 *ec2.DeleteTransitGatewayRouteTableInput, arg3 ...func(*ec2.Options)) (*ec2.DeleteTransitGatewayRouteTableOutput, error)
	// Advertises to the transit gateway that a transit gateway route table is deleted.
	//
	DeleteTransitGatewayRouteTableAnnouncement(arg1 context.Context, arg2 *ec2.DeleteTransitGatewayRouteTableAnnouncementInput, arg3 ...func(*ec2.Options)) (*ec2.DeleteTransitGatewayRouteTableAnnouncementOutput, error)
	// Deletes the specified VPC attachment.
	//
	DeleteTransitGatewayVpcAttachment(arg1 context.Context, arg2 *ec2.DeleteTransitGatewayVpcAttachmentInput, arg3 ...func(*ec2.Options)) (*ec2.DeleteTransitGatewayVpcAttachmentOutput, error)
	// Deletes the specified EBS volume. The volume must be in the available state (not
	// attached to an instance). The volume can remain in the deleting state for
	// several minutes. For more information, see Delete an Amazon EBS volume
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ebs-deleting-volume.html)
	// in the Amazon Elastic Compute Cloud User Guide.
	//
	DeleteVolume(arg1 context.Context, arg2 *ec2.DeleteVolumeInput, arg3 ...func(*ec2.Options)) (*ec2.DeleteVolumeOutput, error)
	// Deletes the specified VPC. You must detach or delete all gateways and resources
	// that are associated with the VPC before you can delete it. For example, you must
	// terminate all instances running in the VPC, delete all security groups
	// associated with the VPC (except the default one), delete all route tables
	// associated with the VPC (except the default one), and so on.
	//
	DeleteVpc(arg1 context.Context, arg2 *ec2.DeleteVpcInput, arg3 ...func(*ec2.Options)) (*ec2.DeleteVpcOutput, error)
	// Deletes one or more VPC endpoint connection notifications.
	//
	DeleteVpcEndpointConnectionNotifications(arg1 context.Context, arg2 *ec2.DeleteVpcEndpointConnectionNotificationsInput, arg3 ...func(*ec2.Options)) (*ec2.DeleteVpcEndpointConnectionNotificationsOutput, error)
	// Deletes one or more VPC endpoint service configurations in your account. Before
	// you delete the endpoint service configuration, you must reject any Available or
	// PendingAcceptance interface endpoint connections that are attached to the
	// service.
	//
	DeleteVpcEndpointServiceConfigurations(arg1 context.Context, arg2 *ec2.DeleteVpcEndpointServiceConfigurationsInput, arg3 ...func(*ec2.Options)) (*ec2.DeleteVpcEndpointServiceConfigurationsOutput, error)
	// Deletes one or more specified VPC endpoints. You can delete any of the following
	// types of VPC endpoints.
	//
	// * Gateway endpoint,
	//
	// * Gateway Load Balancer
	// endpoint,
	//
	// * Interface endpoint
	//
	// The following rules apply when you delete a VPC
	// endpoint:
	//
	// * When you delete a gateway endpoint, we delete the endpoint routes
	// in the route tables that are associated with the endpoint.
	//
	// * When you delete a
	// Gateway Load Balancer endpoint, we delete the endpoint network interfaces. You
	// can only delete Gateway Load Balancer endpoints when the routes that are
	// associated with the endpoint are deleted.
	//
	// * When you delete an interface
	// endpoint, we delete the endpoint network interfaces.
	//
	DeleteVpcEndpoints(arg1 context.Context, arg2 *ec2.DeleteVpcEndpointsInput, arg3 ...func(*ec2.Options)) (*ec2.DeleteVpcEndpointsOutput, error)
	// Deletes a VPC peering connection. Either the owner of the requester VPC or the
	// owner of the accepter VPC can delete the VPC peering connection if it's in the
	// active state. The owner of the requester VPC can delete a VPC peering connection
	// in the pending-acceptance state. You cannot delete a VPC peering connection
	// that's in the failed state.
	//
	DeleteVpcPeeringConnection(arg1 context.Context, arg2 *ec2.DeleteVpcPeeringConnectionInput, arg3 ...func(*ec2.Options)) (*ec2.DeleteVpcPeeringConnectionOutput, error)
	// Deletes the specified VPN connection. If you're deleting the VPC and its
	// associated components, we recommend that you detach the virtual private gateway
	// from the VPC and delete the VPC before deleting the VPN connection. If you
	// believe that the tunnel credentials for your VPN connection have been
	// compromised, you can delete the VPN connection and create a new one that has new
	// keys, without needing to delete the VPC or virtual private gateway. If you
	// create a new VPN connection, you must reconfigure the customer gateway device
	// using the new configuration information returned with the new VPN connection ID.
	// For certificate-based authentication, delete all Certificate Manager (ACM)
	// private certificates used for the Amazon Web Services-side tunnel endpoints for
	// the VPN connection before deleting the VPN connection.
	//
	DeleteVpnConnection(arg1 context.Context, arg2 *ec2.DeleteVpnConnectionInput, arg3 ...func(*ec2.Options)) (*ec2.DeleteVpnConnectionOutput, error)
	// Deletes the specified static route associated with a VPN connection between an
	// existing virtual private gateway and a VPN customer gateway. The static route
	// allows traffic to be routed from the virtual private gateway to the VPN customer
	// gateway.
	//
	DeleteVpnConnectionRoute(arg1 context.Context, arg2 *ec2.DeleteVpnConnectionRouteInput, arg3 ...func(*ec2.Options)) (*ec2.DeleteVpnConnectionRouteOutput, error)
	// Deletes the specified virtual private gateway. You must first detach the virtual
	// private gateway from the VPC. Note that you don't need to delete the virtual
	// private gateway if you plan to delete and recreate the VPN connection between
	// your VPC and your network.
	//
	DeleteVpnGateway(arg1 context.Context, arg2 *ec2.DeleteVpnGatewayInput, arg3 ...func(*ec2.Options)) (*ec2.DeleteVpnGatewayOutput, error)
	// Releases the specified address range that you provisioned for use with your
	// Amazon Web Services resources through bring your own IP addresses (BYOIP) and
	// deletes the corresponding address pool. Before you can release an address range,
	// you must stop advertising it using WithdrawByoipCidr and you must not have any
	// IP addresses allocated from its address range.
	//
	DeprovisionByoipCidr(arg1 context.Context, arg2 *ec2.DeprovisionByoipCidrInput, arg3 ...func(*ec2.Options)) (*ec2.DeprovisionByoipCidrOutput, error)
	// Deprovision a CIDR provisioned from an IPAM pool. If you deprovision a CIDR from
	// a pool that has a source pool, the CIDR is recycled back into the source pool.
	// For more information, see Deprovision pool CIDRs
	// (https://docs.aws.amazon.com/vpc/latest/ipam/depro-pool-cidr-ipam.html) in the
	// Amazon VPC IPAM User Guide.
	//
	DeprovisionIpamPoolCidr(arg1 context.Context, arg2 *ec2.DeprovisionIpamPoolCidrInput, arg3 ...func(*ec2.Options)) (*ec2.DeprovisionIpamPoolCidrOutput, error)
	// Deprovision a CIDR from a public IPv4 pool.
	//
	DeprovisionPublicIpv4PoolCidr(arg1 context.Context, arg2 *ec2.DeprovisionPublicIpv4PoolCidrInput, arg3 ...func(*ec2.Options)) (*ec2.DeprovisionPublicIpv4PoolCidrOutput, error)
	// Deregisters the specified AMI. After you deregister an AMI, it can't be used to
	// launch new instances. If you deregister an AMI that matches a Recycle Bin
	// retention rule, the AMI is retained in the Recycle Bin for the specified
	// retention period. For more information, see Recycle Bin
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/recycle-bin.html) in the
	// Amazon Elastic Compute Cloud User Guide. When you deregister an AMI, it doesn't
	// affect any instances that you've already launched from the AMI. You'll continue
	// to incur usage costs for those instances until you terminate them. When you
	// deregister an Amazon EBS-backed AMI, it doesn't affect the snapshot that was
	// created for the root volume of the instance during the AMI creation process.
	// When you deregister an instance store-backed AMI, it doesn't affect the files
	// that you uploaded to Amazon S3 when you created the AMI.
	//
	DeregisterImage(arg1 context.Context, arg2 *ec2.DeregisterImageInput, arg3 ...func(*ec2.Options)) (*ec2.DeregisterImageOutput, error)
	// Deregisters tag keys to prevent tags that have the specified tag keys from being
	// included in scheduled event notifications for resources in the Region.
	//
	DeregisterInstanceEventNotificationAttributes(arg1 context.Context, arg2 *ec2.DeregisterInstanceEventNotificationAttributesInput, arg3 ...func(*ec2.Options)) (*ec2.DeregisterInstanceEventNotificationAttributesOutput, error)
	// Deregisters the specified members (network interfaces) from the transit gateway
	// multicast group.
	//
	DeregisterTransitGatewayMulticastGroupMembers(arg1 context.Context, arg2 *ec2.DeregisterTransitGatewayMulticastGroupMembersInput, arg3 ...func(*ec2.Options)) (*ec2.DeregisterTransitGatewayMulticastGroupMembersOutput, error)
	// Deregisters the specified sources (network interfaces) from the transit gateway
	// multicast group.
	//
	DeregisterTransitGatewayMulticastGroupSources(arg1 context.Context, arg2 *ec2.DeregisterTransitGatewayMulticastGroupSourcesInput, arg3 ...func(*ec2.Options)) (*ec2.DeregisterTransitGatewayMulticastGroupSourcesOutput, error)
	// Describes attributes of your Amazon Web Services account. The following are the
	// supported account attributes:
	//
	// * supported-platforms: Indicates whether your
	// account can launch instances into EC2-Classic and EC2-VPC, or only into
	// EC2-VPC.
	//
	// * default-vpc: The ID of the default VPC for your account, or none.
	//
	// *
	// max-instances: This attribute is no longer supported. The returned value does
	// not reflect your actual vCPU limit for running On-Demand Instances. For more
	// information, see On-Demand Instance Limits
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-on-demand-instances.html#ec2-on-demand-instances-limits)
	// in the Amazon Elastic Compute Cloud User Guide.
	//
	// *
	// vpc-max-security-groups-per-interface: The maximum number of security groups
	// that you can assign to a network interface.
	//
	// * max-elastic-ips: The maximum
	// number of Elastic IP addresses that you can allocate for use with
	// EC2-Classic.
	//
	// * vpc-max-elastic-ips: The maximum number of Elastic IP addresses
	// that you can allocate for use with EC2-VPC.
	//
	DescribeAccountAttributes(arg1 context.Context, arg2 *ec2.DescribeAccountAttributesInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeAccountAttributesOutput, error)
	// Describes the specified Elastic IP addresses or all of your Elastic IP
	// addresses. An Elastic IP address is for use in either the EC2-Classic platform
	// or in a VPC. For more information, see Elastic IP Addresses
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/elastic-ip-addresses-eip.html)
	// in the Amazon Elastic Compute Cloud User Guide.
	//
	DescribeAddresses(arg1 context.Context, arg2 *ec2.DescribeAddressesInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeAddressesOutput, error)
	// Describes the attributes of the specified Elastic IP addresses. For
	// requirements, see Using reverse DNS for email applications
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/elastic-ip-addresses-eip.html#Using_Elastic_Addressing_Reverse_DNS).
	//
	DescribeAddressesAttribute(arg1 context.Context, arg2 *ec2.DescribeAddressesAttributeInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeAddressesAttributeOutput, error)
	// Describes the longer ID format settings for all resource types in a specific
	// Region. This request is useful for performing a quick audit to determine whether
	// a specific Region is fully opted in for longer IDs (17-character IDs). This
	// request only returns information about resource types that support longer IDs.
	// The following resource types support longer IDs: bundle | conversion-task |
	// customer-gateway | dhcp-options | elastic-ip-allocation | elastic-ip-association
	// | export-task | flow-log | image | import-task | instance | internet-gateway |
	// network-acl | network-acl-association | network-interface |
	// network-interface-attachment | prefix-list | reservation | route-table |
	// route-table-association | security-group | snapshot | subnet |
	// subnet-cidr-block-association | volume | vpc | vpc-cidr-block-association |
	// vpc-endpoint | vpc-peering-connection | vpn-connection | vpn-gateway.
	//
	DescribeAggregateIdFormat(arg1 context.Context, arg2 *ec2.DescribeAggregateIdFormatInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeAggregateIdFormatOutput, error)
	// Describes the Availability Zones, Local Zones, and Wavelength Zones that are
	// available to you. If there is an event impacting a zone, you can use this
	// request to view the state and any provided messages for that zone. For more
	// information about Availability Zones, Local Zones, and Wavelength Zones, see
	// Regions and zones
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/using-regions-availability-zones.html)
	// in the Amazon Elastic Compute Cloud User Guide.
	//
	DescribeAvailabilityZones(arg1 context.Context, arg2 *ec2.DescribeAvailabilityZonesInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeAvailabilityZonesOutput, error)
	// Describes the specified bundle tasks or all of your bundle tasks. Completed
	// bundle tasks are listed for only a limited time. If your bundle task is no
	// longer in the list, you can still register an AMI from it. Just use
	// RegisterImage with the Amazon S3 bucket name and image manifest name you
	// provided to the bundle task.
	//
	DescribeBundleTasks(arg1 context.Context, arg2 *ec2.DescribeBundleTasksInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeBundleTasksOutput, error)
	// Describes the IP address ranges that were specified in calls to
	// ProvisionByoipCidr. To describe the address pools that were created when you
	// provisioned the address ranges, use DescribePublicIpv4Pools or
	// DescribeIpv6Pools.
	//
	DescribeByoipCidrs(arg1 context.Context, arg2 *ec2.DescribeByoipCidrsInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeByoipCidrsOutput, error)
	// Describes one or more Capacity Reservation Fleets.
	//
	DescribeCapacityReservationFleets(arg1 context.Context, arg2 *ec2.DescribeCapacityReservationFleetsInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeCapacityReservationFleetsOutput, error)
	// Describes one or more of your Capacity Reservations. The results describe only
	// the Capacity Reservations in the Amazon Web Services Region that you're
	// currently using.
	//
	DescribeCapacityReservations(arg1 context.Context, arg2 *ec2.DescribeCapacityReservationsInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeCapacityReservationsOutput, error)
	// Describes one or more of your carrier gateways.
	//
	DescribeCarrierGateways(arg1 context.Context, arg2 *ec2.DescribeCarrierGatewaysInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeCarrierGatewaysOutput, error)
	// Describes one or more of your linked EC2-Classic instances. This request only
	// returns information about EC2-Classic instances linked to a VPC through
	// ClassicLink. You cannot use this request to return information about other
	// instances.
	//
	DescribeClassicLinkInstances(arg1 context.Context, arg2 *ec2.DescribeClassicLinkInstancesInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeClassicLinkInstancesOutput, error)
	// Describes the authorization rules for a specified Client VPN endpoint.
	//
	DescribeClientVpnAuthorizationRules(arg1 context.Context, arg2 *ec2.DescribeClientVpnAuthorizationRulesInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeClientVpnAuthorizationRulesOutput, error)
	// Describes active client connections and connections that have been terminated
	// within the last 60 minutes for the specified Client VPN endpoint.
	//
	DescribeClientVpnConnections(arg1 context.Context, arg2 *ec2.DescribeClientVpnConnectionsInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeClientVpnConnectionsOutput, error)
	// Describes one or more Client VPN endpoints in the account.
	//
	DescribeClientVpnEndpoints(arg1 context.Context, arg2 *ec2.DescribeClientVpnEndpointsInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeClientVpnEndpointsOutput, error)
	// Describes the routes for the specified Client VPN endpoint.
	//
	DescribeClientVpnRoutes(arg1 context.Context, arg2 *ec2.DescribeClientVpnRoutesInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeClientVpnRoutesOutput, error)
	// Describes the target networks associated with the specified Client VPN endpoint.
	//
	DescribeClientVpnTargetNetworks(arg1 context.Context, arg2 *ec2.DescribeClientVpnTargetNetworksInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeClientVpnTargetNetworksOutput, error)
	// Describes the specified customer-owned address pools or all of your
	// customer-owned address pools.
	//
	DescribeCoipPools(arg1 context.Context, arg2 *ec2.DescribeCoipPoolsInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeCoipPoolsOutput, error)
	// Describes the specified conversion tasks or all your conversion tasks. For more
	// information, see the VM Import/Export User Guide
	// (https://docs.aws.amazon.com/vm-import/latest/userguide/). For information about
	// the import manifest referenced by this API action, see VM Import Manifest
	// (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/manifest.html).
	//
	DescribeConversionTasks(arg1 context.Context, arg2 *ec2.DescribeConversionTasksInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeConversionTasksOutput, error)
	// Describes one or more of your VPN customer gateways. For more information, see
	// Amazon Web Services Site-to-Site VPN
	// (https://docs.aws.amazon.com/vpn/latest/s2svpn/VPC_VPN.html) in the Amazon Web
	// Services Site-to-Site VPN User Guide.
	//
	DescribeCustomerGateways(arg1 context.Context, arg2 *ec2.DescribeCustomerGatewaysInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeCustomerGatewaysOutput, error)
	// Describes one or more of your DHCP options sets. For more information, see DHCP
	// options sets
	// (https://docs.aws.amazon.com/vpc/latest/userguide/VPC_DHCP_Options.html) in the
	// Amazon Virtual Private Cloud User Guide.
	//
	DescribeDhcpOptions(arg1 context.Context, arg2 *ec2.DescribeDhcpOptionsInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeDhcpOptionsOutput, error)
	// Describes one or more of your egress-only internet gateways.
	//
	DescribeEgressOnlyInternetGateways(arg1 context.Context, arg2 *ec2.DescribeEgressOnlyInternetGatewaysInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeEgressOnlyInternetGatewaysOutput, error)
	// Describes the Elastic Graphics accelerator associated with your instances. For
	// more information about Elastic Graphics, see Amazon Elastic Graphics
	// (https://docs.aws.amazon.com/AWSEC2/latest/WindowsGuide/elastic-graphics.html).
	//
	DescribeElasticGpus(arg1 context.Context, arg2 *ec2.DescribeElasticGpusInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeElasticGpusOutput, error)
	// Describes the specified export image tasks or all of your export image tasks.
	//
	DescribeExportImageTasks(arg1 context.Context, arg2 *ec2.DescribeExportImageTasksInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeExportImageTasksOutput, error)
	// Describes the specified export instance tasks or all of your export instance
	// tasks.
	//
	DescribeExportTasks(arg1 context.Context, arg2 *ec2.DescribeExportTasksInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeExportTasksOutput, error)
	// Describe details for Windows AMIs that are configured for faster launching.
	//
	DescribeFastLaunchImages(arg1 context.Context, arg2 *ec2.DescribeFastLaunchImagesInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeFastLaunchImagesOutput, error)
	// Describes the state of fast snapshot restores for your snapshots.
	//
	DescribeFastSnapshotRestores(arg1 context.Context, arg2 *ec2.DescribeFastSnapshotRestoresInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeFastSnapshotRestoresOutput, error)
	// Describes the events for the specified EC2 Fleet during the specified time. EC2
	// Fleet events are delayed by up to 30 seconds before they can be described. This
	// ensures that you can query by the last evaluated time and not miss a recorded
	// event. EC2 Fleet events are available for 48 hours. For more information, see
	// Monitor fleet events using Amazon EventBridge
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/fleet-monitor.html) in the
	// Amazon EC2 User Guide.
	//
	DescribeFleetHistory(arg1 context.Context, arg2 *ec2.DescribeFleetHistoryInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeFleetHistoryOutput, error)
	// Describes the running instances for the specified EC2 Fleet. For more
	// information, see Monitor your EC2 Fleet
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/manage-ec2-fleet.html#monitor-ec2-fleet)
	// in the Amazon EC2 User Guide.
	//
	DescribeFleetInstances(arg1 context.Context, arg2 *ec2.DescribeFleetInstancesInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeFleetInstancesOutput, error)
	// Describes the specified EC2 Fleets or all of your EC2 Fleets. For more
	// information, see Monitor your EC2 Fleet
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/manage-ec2-fleet.html#monitor-ec2-fleet)
	// in the Amazon EC2 User Guide.
	//
	DescribeFleets(arg1 context.Context, arg2 *ec2.DescribeFleetsInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeFleetsOutput, error)
	// Describes one or more flow logs. To view the information in your flow logs (the
	// log streams for the network interfaces), you must use the CloudWatch Logs
	// console or the CloudWatch Logs API.
	//
	DescribeFlowLogs(arg1 context.Context, arg2 *ec2.DescribeFlowLogsInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeFlowLogsOutput, error)
	// Describes the specified attribute of the specified Amazon FPGA Image (AFI).
	//
	DescribeFpgaImageAttribute(arg1 context.Context, arg2 *ec2.DescribeFpgaImageAttributeInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeFpgaImageAttributeOutput, error)
	// Describes the Amazon FPGA Images (AFIs) available to you. These include public
	// AFIs, private AFIs that you own, and AFIs owned by other Amazon Web Services
	// accounts for which you have load permissions.
	//
	DescribeFpgaImages(arg1 context.Context, arg2 *ec2.DescribeFpgaImagesInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeFpgaImagesOutput, error)
	// Describes the Dedicated Host reservations that are available to purchase. The
	// results describe all of the Dedicated Host reservation offerings, including
	// offerings that might not match the instance family and Region of your Dedicated
	// Hosts. When purchasing an offering, ensure that the instance family and Region
	// of the offering matches that of the Dedicated Hosts with which it is to be
	// associated. For more information about supported instance types, see Dedicated
	// Hosts
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/dedicated-hosts-overview.html)
	// in the Amazon EC2 User Guide.
	//
	DescribeHostReservationOfferings(arg1 context.Context, arg2 *ec2.DescribeHostReservationOfferingsInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeHostReservationOfferingsOutput, error)
	// Describes reservations that are associated with Dedicated Hosts in your account.
	//
	DescribeHostReservations(arg1 context.Context, arg2 *ec2.DescribeHostReservationsInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeHostReservationsOutput, error)
	// Describes the specified Dedicated Hosts or all your Dedicated Hosts. The results
	// describe only the Dedicated Hosts in the Region you're currently using. All
	// listed instances consume capacity on your Dedicated Host. Dedicated Hosts that
	// have recently been released are listed with the state released.
	//
	DescribeHosts(arg1 context.Context, arg2 *ec2.DescribeHostsInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeHostsOutput, error)
	// Describes your IAM instance profile associations.
	//
	DescribeIamInstanceProfileAssociations(arg1 context.Context, arg2 *ec2.DescribeIamInstanceProfileAssociationsInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeIamInstanceProfileAssociationsOutput, error)
	// Describes the ID format settings for your resources on a per-Region basis, for
	// example, to view which resource types are enabled for longer IDs. This request
	// only returns information about resource types whose ID formats can be modified;
	// it does not return information about other resource types. The following
	// resource types support longer IDs: bundle | conversion-task | customer-gateway |
	// dhcp-options | elastic-ip-allocation | elastic-ip-association | export-task |
	// flow-log | image | import-task | instance | internet-gateway | network-acl |
	// network-acl-association | network-interface | network-interface-attachment |
	// prefix-list | reservation | route-table | route-table-association |
	// security-group | snapshot | subnet | subnet-cidr-block-association | volume |
	// vpc | vpc-cidr-block-association | vpc-endpoint | vpc-peering-connection |
	// vpn-connection | vpn-gateway. These settings apply to the IAM user who makes the
	// request; they do not apply to the entire Amazon Web Services account. By
	// default, an IAM user defaults to the same settings as the root user, unless they
	// explicitly override the settings by running the ModifyIdFormat command.
	// Resources created with longer IDs are visible to all IAM users, regardless of
	// these settings and provided that they have permission to use the relevant
	// Describe command for the resource type.
	//
	DescribeIdFormat(arg1 context.Context, arg2 *ec2.DescribeIdFormatInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeIdFormatOutput, error)
	// Describes the ID format settings for resources for the specified IAM user, IAM
	// role, or root user. For example, you can view the resource types that are
	// enabled for longer IDs. This request only returns information about resource
	// types whose ID formats can be modified; it does not return information about
	// other resource types. For more information, see Resource IDs
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/resource-ids.html) in the
	// Amazon Elastic Compute Cloud User Guide. The following resource types support
	// longer IDs: bundle | conversion-task | customer-gateway | dhcp-options |
	// elastic-ip-allocation | elastic-ip-association | export-task | flow-log | image
	// | import-task | instance | internet-gateway | network-acl |
	// network-acl-association | network-interface | network-interface-attachment |
	// prefix-list | reservation | route-table | route-table-association |
	// security-group | snapshot | subnet | subnet-cidr-block-association | volume |
	// vpc | vpc-cidr-block-association | vpc-endpoint | vpc-peering-connection |
	// vpn-connection | vpn-gateway. These settings apply to the principal specified in
	// the request. They do not apply to the principal that makes the request.
	//
	DescribeIdentityIdFormat(arg1 context.Context, arg2 *ec2.DescribeIdentityIdFormatInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeIdentityIdFormatOutput, error)
	// Describes the specified attribute of the specified AMI. You can specify only one
	// attribute at a time.
	//
	DescribeImageAttribute(arg1 context.Context, arg2 *ec2.DescribeImageAttributeInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeImageAttributeOutput, error)
	// Describes the specified images (AMIs, AKIs, and ARIs) available to you or all of
	// the images available to you. The images available to you include public images,
	// private images that you own, and private images owned by other Amazon Web
	// Services accounts for which you have explicit launch permissions. Recently
	// deregistered images appear in the returned results for a short interval and then
	// return empty results. After all instances that reference a deregistered AMI are
	// terminated, specifying the ID of the image will eventually return an error
	// indicating that the AMI ID cannot be found.
	//
	DescribeImages(arg1 context.Context, arg2 *ec2.DescribeImagesInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeImagesOutput, error)
	// Displays details about an import virtual machine or import snapshot tasks that
	// are already created.
	//
	DescribeImportImageTasks(arg1 context.Context, arg2 *ec2.DescribeImportImageTasksInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeImportImageTasksOutput, error)
	// Describes your import snapshot tasks.
	//
	DescribeImportSnapshotTasks(arg1 context.Context, arg2 *ec2.DescribeImportSnapshotTasksInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeImportSnapshotTasksOutput, error)
	// Describes the specified attribute of the specified instance. You can specify
	// only one attribute at a time. Valid attribute values are: instanceType | kernel
	// | ramdisk | userData | disableApiTermination | instanceInitiatedShutdownBehavior
	// | rootDeviceName | blockDeviceMapping | productCodes | sourceDestCheck |
	// groupSet | ebsOptimized | sriovNetSupport
	//
	DescribeInstanceAttribute(arg1 context.Context, arg2 *ec2.DescribeInstanceAttributeInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeInstanceAttributeOutput, error)
	// Describes the credit option for CPU usage of the specified burstable performance
	// instances. The credit options are standard and unlimited. If you do not specify
	// an instance ID, Amazon EC2 returns burstable performance instances with the
	// unlimited credit option, as well as instances that were previously configured as
	// T2, T3, and T3a with the unlimited credit option. For example, if you resize a
	// T2 instance, while it is configured as unlimited, to an M4 instance, Amazon EC2
	// returns the M4 instance. If you specify one or more instance IDs, Amazon EC2
	// returns the credit option (standard or unlimited) of those instances. If you
	// specify an instance ID that is not valid, such as an instance that is not a
	// burstable performance instance, an error is returned. Recently terminated
	// instances might appear in the returned results. This interval is usually less
	// than one hour. If an Availability Zone is experiencing a service disruption and
	// you specify instance IDs in the affected zone, or do not specify any instance
	// IDs at all, the call fails. If you specify only instance IDs in an unaffected
	// zone, the call works normally. For more information, see Burstable performance
	// instances
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/burstable-performance-instances.html)
	// in the Amazon EC2 User Guide.
	//
	DescribeInstanceCreditSpecifications(arg1 context.Context, arg2 *ec2.DescribeInstanceCreditSpecificationsInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeInstanceCreditSpecificationsOutput, error)
	// Describes the tag keys that are registered to appear in scheduled event
	// notifications for resources in the current Region.
	//
	DescribeInstanceEventNotificationAttributes(arg1 context.Context, arg2 *ec2.DescribeInstanceEventNotificationAttributesInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeInstanceEventNotificationAttributesOutput, error)
	// Describes the specified event windows or all event windows. If you specify event
	// window IDs, the output includes information for only the specified event
	// windows. If you specify filters, the output includes information for only those
	// event windows that meet the filter criteria. If you do not specify event windows
	// IDs or filters, the output includes information for all event windows, which can
	// affect performance. We recommend that you use pagination to ensure that the
	// operation returns quickly and successfully. For more information, see Define
	// event windows for scheduled events
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/event-windows.html) in the
	// Amazon EC2 User Guide.
	//
	DescribeInstanceEventWindows(arg1 context.Context, arg2 *ec2.DescribeInstanceEventWindowsInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeInstanceEventWindowsOutput, error)
	// Describes the status of the specified instances or all of your instances. By
	// default, only running instances are described, unless you specifically indicate
	// to return the status of all instances. Instance status includes the following
	// components:
	//
	// * Status checks - Amazon EC2 performs status checks on running EC2
	// instances to identify hardware and software issues. For more information, see
	// Status checks for your instances
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/monitoring-system-instance-status-check.html)
	// and Troubleshoot instances with failed status checks
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/TroubleshootingInstances.html)
	// in the Amazon EC2 User Guide.
	//
	// * Scheduled events - Amazon EC2 can schedule
	// events (such as reboot, stop, or terminate) for your instances related to
	// hardware issues, software updates, or system maintenance. For more information,
	// see Scheduled events for your instances
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/monitoring-instances-status-check_sched.html)
	// in the Amazon EC2 User Guide.
	//
	// * Instance state - You can manage your instances
	// from the moment you launch them through their termination. For more information,
	// see Instance lifecycle
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-instance-lifecycle.html)
	// in the Amazon EC2 User Guide.
	//
	DescribeInstanceStatus(arg1 context.Context, arg2 *ec2.DescribeInstanceStatusInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeInstanceStatusOutput, error)
	// Returns a list of all instance types offered. The results can be filtered by
	// location (Region or Availability Zone). If no location is specified, the
	// instance types offered in the current Region are returned.
	//
	DescribeInstanceTypeOfferings(arg1 context.Context, arg2 *ec2.DescribeInstanceTypeOfferingsInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeInstanceTypeOfferingsOutput, error)
	// Describes the details of the instance types that are offered in a location. The
	// results can be filtered by the attributes of the instance types.
	//
	DescribeInstanceTypes(arg1 context.Context, arg2 *ec2.DescribeInstanceTypesInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeInstanceTypesOutput, error)
	// Describes the specified instances or all instances. If you specify instance IDs,
	// the output includes information for only the specified instances. If you specify
	// filters, the output includes information for only those instances that meet the
	// filter criteria. If you do not specify instance IDs or filters, the output
	// includes information for all instances, which can affect performance. We
	// recommend that you use pagination to ensure that the operation returns quickly
	// and successfully. If you specify an instance ID that is not valid, an error is
	// returned. If you specify an instance that you do not own, it is not included in
	// the output. Recently terminated instances might appear in the returned results.
	// This interval is usually less than one hour. If you describe instances in the
	// rare case where an Availability Zone is experiencing a service disruption and
	// you specify instance IDs that are in the affected zone, or do not specify any
	// instance IDs at all, the call fails. If you describe instances and specify only
	// instance IDs that are in an unaffected zone, the call works normally.
	//
	DescribeInstances(arg1 context.Context, arg2 *ec2.DescribeInstancesInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeInstancesOutput, error)
	// Describes one or more of your internet gateways.
	//
	DescribeInternetGateways(arg1 context.Context, arg2 *ec2.DescribeInternetGatewaysInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeInternetGatewaysOutput, error)
	// Get information about your IPAM pools.
	//
	DescribeIpamPools(arg1 context.Context, arg2 *ec2.DescribeIpamPoolsInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeIpamPoolsOutput, error)
	// Get information about your IPAM scopes.
	//
	DescribeIpamScopes(arg1 context.Context, arg2 *ec2.DescribeIpamScopesInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeIpamScopesOutput, error)
	// Get information about your IPAM pools. For more information, see What is IPAM?
	// (https://docs.aws.amazon.com/vpc/latest/ipam/what-is-it-ipam.html) in the Amazon
	// VPC IPAM User Guide.
	//
	DescribeIpams(arg1 context.Context, arg2 *ec2.DescribeIpamsInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeIpamsOutput, error)
	// Describes your IPv6 address pools.
	//
	DescribeIpv6Pools(arg1 context.Context, arg2 *ec2.DescribeIpv6PoolsInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeIpv6PoolsOutput, error)
	// Describes the specified key pairs or all of your key pairs. For more information
	// about key pairs, see Amazon EC2 key pairs
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-key-pairs.html) in the
	// Amazon Elastic Compute Cloud User Guide.
	//
	DescribeKeyPairs(arg1 context.Context, arg2 *ec2.DescribeKeyPairsInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeKeyPairsOutput, error)
	// Describes one or more versions of a specified launch template. You can describe
	// all versions, individual versions, or a range of versions. You can also describe
	// all the latest versions or all the default versions of all the launch templates
	// in your account.
	//
	DescribeLaunchTemplateVersions(arg1 context.Context, arg2 *ec2.DescribeLaunchTemplateVersionsInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeLaunchTemplateVersionsOutput, error)
	// Describes one or more launch templates.
	//
	DescribeLaunchTemplates(arg1 context.Context, arg2 *ec2.DescribeLaunchTemplatesInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeLaunchTemplatesOutput, error)
	// Describes the associations between virtual interface groups and local gateway
	// route tables.
	//
	DescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociations(arg1 context.Context, arg2 *ec2.DescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociationsInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociationsOutput, error)
	// Describes the specified associations between VPCs and local gateway route
	// tables.
	//
	DescribeLocalGatewayRouteTableVpcAssociations(arg1 context.Context, arg2 *ec2.DescribeLocalGatewayRouteTableVpcAssociationsInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeLocalGatewayRouteTableVpcAssociationsOutput, error)
	// Describes one or more local gateway route tables. By default, all local gateway
	// route tables are described. Alternatively, you can filter the results.
	//
	DescribeLocalGatewayRouteTables(arg1 context.Context, arg2 *ec2.DescribeLocalGatewayRouteTablesInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeLocalGatewayRouteTablesOutput, error)
	// Describes the specified local gateway virtual interface groups.
	//
	DescribeLocalGatewayVirtualInterfaceGroups(arg1 context.Context, arg2 *ec2.DescribeLocalGatewayVirtualInterfaceGroupsInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeLocalGatewayVirtualInterfaceGroupsOutput, error)
	// Describes the specified local gateway virtual interfaces.
	//
	DescribeLocalGatewayVirtualInterfaces(arg1 context.Context, arg2 *ec2.DescribeLocalGatewayVirtualInterfacesInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeLocalGatewayVirtualInterfacesOutput, error)
	// Describes one or more local gateways. By default, all local gateways are
	// described. Alternatively, you can filter the results.
	//
	DescribeLocalGateways(arg1 context.Context, arg2 *ec2.DescribeLocalGatewaysInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeLocalGatewaysOutput, error)
	// Describes your managed prefix lists and any Amazon Web Services-managed prefix
	// lists. To view the entries for your prefix list, use
	// GetManagedPrefixListEntries.
	//
	DescribeManagedPrefixLists(arg1 context.Context, arg2 *ec2.DescribeManagedPrefixListsInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeManagedPrefixListsOutput, error)
	// Describes your Elastic IP addresses that are being moved to the EC2-VPC
	// platform, or that are being restored to the EC2-Classic platform. This request
	// does not return information about any other Elastic IP addresses in your
	// account.
	//
	DescribeMovingAddresses(arg1 context.Context, arg2 *ec2.DescribeMovingAddressesInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeMovingAddressesOutput, error)
	// Describes one or more of your NAT gateways.
	//
	DescribeNatGateways(arg1 context.Context, arg2 *ec2.DescribeNatGatewaysInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeNatGatewaysOutput, error)
	// Describes one or more of your network ACLs. For more information, see Network
	// ACLs (https://docs.aws.amazon.com/vpc/latest/userguide/VPC_ACLs.html) in the
	// Amazon Virtual Private Cloud User Guide.
	//
	DescribeNetworkAcls(arg1 context.Context, arg2 *ec2.DescribeNetworkAclsInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeNetworkAclsOutput, error)
	// Describes the specified Network Access Scope analyses.
	//
	DescribeNetworkInsightsAccessScopeAnalyses(arg1 context.Context, arg2 *ec2.DescribeNetworkInsightsAccessScopeAnalysesInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeNetworkInsightsAccessScopeAnalysesOutput, error)
	// Describes the specified Network Access Scopes.
	//
	DescribeNetworkInsightsAccessScopes(arg1 context.Context, arg2 *ec2.DescribeNetworkInsightsAccessScopesInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeNetworkInsightsAccessScopesOutput, error)
	// Describes one or more of your network insights analyses.
	//
	DescribeNetworkInsightsAnalyses(arg1 context.Context, arg2 *ec2.DescribeNetworkInsightsAnalysesInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeNetworkInsightsAnalysesOutput, error)
	// Describes one or more of your paths.
	//
	DescribeNetworkInsightsPaths(arg1 context.Context, arg2 *ec2.DescribeNetworkInsightsPathsInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeNetworkInsightsPathsOutput, error)
	// Describes a network interface attribute. You can specify only one attribute at a
	// time.
	//
	DescribeNetworkInterfaceAttribute(arg1 context.Context, arg2 *ec2.DescribeNetworkInterfaceAttributeInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeNetworkInterfaceAttributeOutput, error)
	// Describes the permissions for your network interfaces.
	//
	DescribeNetworkInterfacePermissions(arg1 context.Context, arg2 *ec2.DescribeNetworkInterfacePermissionsInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeNetworkInterfacePermissionsOutput, error)
	// Describes one or more of your network interfaces.
	//
	DescribeNetworkInterfaces(arg1 context.Context, arg2 *ec2.DescribeNetworkInterfacesInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeNetworkInterfacesOutput, error)
	// Describes the specified placement groups or all of your placement groups. For
	// more information, see Placement groups
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/placement-groups.html) in
	// the Amazon EC2 User Guide.
	//
	DescribePlacementGroups(arg1 context.Context, arg2 *ec2.DescribePlacementGroupsInput, arg3 ...func(*ec2.Options)) (*ec2.DescribePlacementGroupsOutput, error)
	// Describes available Amazon Web Services services in a prefix list format, which
	// includes the prefix list name and prefix list ID of the service and the IP
	// address range for the service. We recommend that you use
	// DescribeManagedPrefixLists instead.
	//
	DescribePrefixLists(arg1 context.Context, arg2 *ec2.DescribePrefixListsInput, arg3 ...func(*ec2.Options)) (*ec2.DescribePrefixListsOutput, error)
	// Describes the ID format settings for the root user and all IAM roles and IAM
	// users that have explicitly specified a longer ID (17-character ID) preference.
	// By default, all IAM roles and IAM users default to the same ID settings as the
	// root user, unless they explicitly override the settings. This request is useful
	// for identifying those IAM users and IAM roles that have overridden the default
	// ID settings. The following resource types support longer IDs: bundle |
	// conversion-task | customer-gateway | dhcp-options | elastic-ip-allocation |
	// elastic-ip-association | export-task | flow-log | image | import-task | instance
	// | internet-gateway | network-acl | network-acl-association | network-interface |
	// network-interface-attachment | prefix-list | reservation | route-table |
	// route-table-association | security-group | snapshot | subnet |
	// subnet-cidr-block-association | volume | vpc | vpc-cidr-block-association |
	// vpc-endpoint | vpc-peering-connection | vpn-connection | vpn-gateway.
	//
	DescribePrincipalIdFormat(arg1 context.Context, arg2 *ec2.DescribePrincipalIdFormatInput, arg3 ...func(*ec2.Options)) (*ec2.DescribePrincipalIdFormatOutput, error)
	// Describes the specified IPv4 address pools.
	//
	DescribePublicIpv4Pools(arg1 context.Context, arg2 *ec2.DescribePublicIpv4PoolsInput, arg3 ...func(*ec2.Options)) (*ec2.DescribePublicIpv4PoolsOutput, error)
	// Describes the Regions that are enabled for your account, or all Regions. For a
	// list of the Regions supported by Amazon EC2, see  Amazon Elastic Compute Cloud
	// endpoints and quotas
	// (https://docs.aws.amazon.com/general/latest/gr/ec2-service.html). For
	// information about enabling and disabling Regions for your account, see Managing
	// Amazon Web Services Regions
	// (https://docs.aws.amazon.com/general/latest/gr/rande-manage.html) in the Amazon
	// Web Services General Reference.
	//
	DescribeRegions(arg1 context.Context, arg2 *ec2.DescribeRegionsInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeRegionsOutput, error)
	// Describes a root volume replacement task. For more information, see Replace a
	// root volume
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ebs-restoring-volume.html#replace-root)
	// in the Amazon Elastic Compute Cloud User Guide.
	//
	DescribeReplaceRootVolumeTasks(arg1 context.Context, arg2 *ec2.DescribeReplaceRootVolumeTasksInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeReplaceRootVolumeTasksOutput, error)
	// Describes one or more of the Reserved Instances that you purchased. For more
	// information about Reserved Instances, see Reserved Instances
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/concepts-on-demand-reserved-instances.html)
	// in the Amazon EC2 User Guide.
	//
	DescribeReservedInstances(arg1 context.Context, arg2 *ec2.DescribeReservedInstancesInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeReservedInstancesOutput, error)
	// Describes your account's Reserved Instance listings in the Reserved Instance
	// Marketplace. The Reserved Instance Marketplace matches sellers who want to
	// resell Reserved Instance capacity that they no longer need with buyers who want
	// to purchase additional capacity. Reserved Instances bought and sold through the
	// Reserved Instance Marketplace work like any other Reserved Instances. As a
	// seller, you choose to list some or all of your Reserved Instances, and you
	// specify the upfront price to receive for them. Your Reserved Instances are then
	// listed in the Reserved Instance Marketplace and are available for purchase. As a
	// buyer, you specify the configuration of the Reserved Instance to purchase, and
	// the Marketplace matches what you're searching for with what's available. The
	// Marketplace first sells the lowest priced Reserved Instances to you, and
	// continues to sell available Reserved Instance listings to you until your demand
	// is met. You are charged based on the total price of all of the listings that you
	// purchase. For more information, see Reserved Instance Marketplace
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ri-market-general.html) in
	// the Amazon EC2 User Guide.
	//
	DescribeReservedInstancesListings(arg1 context.Context, arg2 *ec2.DescribeReservedInstancesListingsInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeReservedInstancesListingsOutput, error)
	// Describes the modifications made to your Reserved Instances. If no parameter is
	// specified, information about all your Reserved Instances modification requests
	// is returned. If a modification ID is specified, only information about the
	// specific modification is returned. For more information, see Modifying Reserved
	// Instances
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ri-modifying.html) in the
	// Amazon EC2 User Guide.
	//
	DescribeReservedInstancesModifications(arg1 context.Context, arg2 *ec2.DescribeReservedInstancesModificationsInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeReservedInstancesModificationsOutput, error)
	// Describes Reserved Instance offerings that are available for purchase. With
	// Reserved Instances, you purchase the right to launch instances for a period of
	// time. During that time period, you do not receive insufficient capacity errors,
	// and you pay a lower usage rate than the rate charged for On-Demand instances for
	// the actual time used. If you have listed your own Reserved Instances for sale in
	// the Reserved Instance Marketplace, they will be excluded from these results.
	// This is to ensure that you do not purchase your own Reserved Instances. For more
	// information, see Reserved Instance Marketplace
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ri-market-general.html) in
	// the Amazon EC2 User Guide.
	//
	DescribeReservedInstancesOfferings(arg1 context.Context, arg2 *ec2.DescribeReservedInstancesOfferingsInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeReservedInstancesOfferingsOutput, error)
	// Describes one or more of your route tables. Each subnet in your VPC must be
	// associated with a route table. If a subnet is not explicitly associated with any
	// route table, it is implicitly associated with the main route table. This command
	// does not return the subnet ID for implicit associations. For more information,
	// see Route tables
	// (https://docs.aws.amazon.com/vpc/latest/userguide/VPC_Route_Tables.html) in the
	// Amazon Virtual Private Cloud User Guide.
	//
	DescribeRouteTables(arg1 context.Context, arg2 *ec2.DescribeRouteTablesInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeRouteTablesOutput, error)
	// Finds available schedules that meet the specified criteria. You can search for
	// an available schedule no more than 3 months in advance. You must meet the
	// minimum required duration of 1,200 hours per year. For example, the minimum
	// daily schedule is 4 hours, the minimum weekly schedule is 24 hours, and the
	// minimum monthly schedule is 100 hours. After you find a schedule that meets your
	// needs, call PurchaseScheduledInstances to purchase Scheduled Instances with that
	// schedule.
	//
	DescribeScheduledInstanceAvailability(arg1 context.Context, arg2 *ec2.DescribeScheduledInstanceAvailabilityInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeScheduledInstanceAvailabilityOutput, error)
	// Describes the specified Scheduled Instances or all your Scheduled Instances.
	//
	DescribeScheduledInstances(arg1 context.Context, arg2 *ec2.DescribeScheduledInstancesInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeScheduledInstancesOutput, error)
	// [VPC only] Describes the VPCs on the other side of a VPC peering connection that
	// are referencing the security groups you've specified in this request.
	//
	DescribeSecurityGroupReferences(arg1 context.Context, arg2 *ec2.DescribeSecurityGroupReferencesInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeSecurityGroupReferencesOutput, error)
	// Describes one or more of your security group rules.
	//
	DescribeSecurityGroupRules(arg1 context.Context, arg2 *ec2.DescribeSecurityGroupRulesInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeSecurityGroupRulesOutput, error)
	// Describes the specified security groups or all of your security groups. A
	// security group is for use with instances either in the EC2-Classic platform or
	// in a specific VPC. For more information, see Amazon EC2 security groups
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/using-network-security.html)
	// in the Amazon Elastic Compute Cloud User Guide and Security groups for your VPC
	// (https://docs.aws.amazon.com/AmazonVPC/latest/UserGuide/VPC_SecurityGroups.html)
	// in the Amazon Virtual Private Cloud User Guide.
	//
	DescribeSecurityGroups(arg1 context.Context, arg2 *ec2.DescribeSecurityGroupsInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeSecurityGroupsOutput, error)
	// Describes the specified attribute of the specified snapshot. You can specify
	// only one attribute at a time. For more information about EBS snapshots, see
	// Amazon EBS snapshots
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSSnapshots.html) in the
	// Amazon Elastic Compute Cloud User Guide.
	//
	DescribeSnapshotAttribute(arg1 context.Context, arg2 *ec2.DescribeSnapshotAttributeInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeSnapshotAttributeOutput, error)
	// Describes the storage tier status of one or more Amazon EBS snapshots.
	//
	DescribeSnapshotTierStatus(arg1 context.Context, arg2 *ec2.DescribeSnapshotTierStatusInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeSnapshotTierStatusOutput, error)
	// Describes the specified EBS snapshots available to you or all of the EBS
	// snapshots available to you. The snapshots available to you include public
	// snapshots, private snapshots that you own, and private snapshots owned by other
	// Amazon Web Services accounts for which you have explicit create volume
	// permissions. The create volume permissions fall into the following
	// categories:
	//
	// * public: The owner of the snapshot granted create volume
	// permissions for the snapshot to the all group. All Amazon Web Services accounts
	// have create volume permissions for these snapshots.
	//
	// * explicit: The owner of
	// the snapshot granted create volume permissions to a specific Amazon Web Services
	// account.
	//
	// * implicit: An Amazon Web Services account has implicit create volume
	// permissions for all snapshots it owns.
	//
	// The list of snapshots returned can be
	// filtered by specifying snapshot IDs, snapshot owners, or Amazon Web Services
	// accounts with create volume permissions. If no options are specified, Amazon EC2
	// returns all snapshots for which you have create volume permissions. If you
	// specify one or more snapshot IDs, only snapshots that have the specified IDs are
	// returned. If you specify an invalid snapshot ID, an error is returned. If you
	// specify a snapshot ID for which you do not have access, it is not included in
	// the returned results. If you specify one or more snapshot owners using the
	// OwnerIds option, only snapshots from the specified owners and for which you have
	// access are returned. The results can include the Amazon Web Services account IDs
	// of the specified owners, amazon for snapshots owned by Amazon, or self for
	// snapshots that you own. If you specify a list of restorable users, only
	// snapshots with create snapshot permissions for those users are returned. You can
	// specify Amazon Web Services account IDs (if you own the snapshots), self for
	// snapshots for which you own or have explicit permissions, or all for public
	// snapshots. If you are describing a long list of snapshots, we recommend that you
	// paginate the output to make the list more manageable. The MaxResults parameter
	// sets the maximum number of results returned in a single page. If the list of
	// results exceeds your MaxResults value, then that number of results is returned
	// along with a NextToken value that can be passed to a subsequent
	// DescribeSnapshots request to retrieve the remaining results. To get the state of
	// fast snapshot restores for a snapshot, use DescribeFastSnapshotRestores. For
	// more information about EBS snapshots, see Amazon EBS snapshots
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSSnapshots.html) in the
	// Amazon Elastic Compute Cloud User Guide.
	//
	DescribeSnapshots(arg1 context.Context, arg2 *ec2.DescribeSnapshotsInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeSnapshotsOutput, error)
	// Describes the data feed for Spot Instances. For more information, see Spot
	// Instance data feed
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/spot-data-feeds.html) in
	// the Amazon EC2 User Guide for Linux Instances.
	//
	DescribeSpotDatafeedSubscription(arg1 context.Context, arg2 *ec2.DescribeSpotDatafeedSubscriptionInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeSpotDatafeedSubscriptionOutput, error)
	// Describes the running instances for the specified Spot Fleet.
	//
	DescribeSpotFleetInstances(arg1 context.Context, arg2 *ec2.DescribeSpotFleetInstancesInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeSpotFleetInstancesOutput, error)
	// Describes the events for the specified Spot Fleet request during the specified
	// time. Spot Fleet events are delayed by up to 30 seconds before they can be
	// described. This ensures that you can query by the last evaluated time and not
	// miss a recorded event. Spot Fleet events are available for 48 hours. For more
	// information, see Monitor fleet events using Amazon EventBridge
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/fleet-monitor.html) in the
	// Amazon EC2 User Guide for Linux Instances.
	//
	DescribeSpotFleetRequestHistory(arg1 context.Context, arg2 *ec2.DescribeSpotFleetRequestHistoryInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeSpotFleetRequestHistoryOutput, error)
	// Describes your Spot Fleet requests. Spot Fleet requests are deleted 48 hours
	// after they are canceled and their instances are terminated.
	//
	DescribeSpotFleetRequests(arg1 context.Context, arg2 *ec2.DescribeSpotFleetRequestsInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeSpotFleetRequestsOutput, error)
	// Describes the specified Spot Instance requests. You can use
	// DescribeSpotInstanceRequests to find a running Spot Instance by examining the
	// response. If the status of the Spot Instance is fulfilled, the instance ID
	// appears in the response and contains the identifier of the instance.
	// Alternatively, you can use DescribeInstances
	// (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribeInstances)
	// with a filter to look for instances where the instance lifecycle is spot. We
	// recommend that you set MaxResults to a value between 5 and 1000 to limit the
	// number of results returned. This paginates the output, which makes the list more
	// manageable and returns the results faster. If the list of results exceeds your
	// MaxResults value, then that number of results is returned along with a NextToken
	// value that can be passed to a subsequent DescribeSpotInstanceRequests request to
	// retrieve the remaining results. Spot Instance requests are deleted four hours
	// after they are canceled and their instances are terminated.
	//
	DescribeSpotInstanceRequests(arg1 context.Context, arg2 *ec2.DescribeSpotInstanceRequestsInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeSpotInstanceRequestsOutput, error)
	// Describes the Spot price history. For more information, see Spot Instance
	// pricing history
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/using-spot-instances-history.html)
	// in the Amazon EC2 User Guide for Linux Instances. When you specify a start and
	// end time, the operation returns the prices of the instance types within that
	// time range. It also returns the last price change before the start time, which
	// is the effective price as of the start time.
	//
	DescribeSpotPriceHistory(arg1 context.Context, arg2 *ec2.DescribeSpotPriceHistoryInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeSpotPriceHistoryOutput, error)
	// [VPC only] Describes the stale security group rules for security groups in a
	// specified VPC. Rules are stale when they reference a deleted security group in
	// the same VPC or in a peer VPC, or if they reference a security group in a peer
	// VPC for which the VPC peering connection has been deleted.
	//
	DescribeStaleSecurityGroups(arg1 context.Context, arg2 *ec2.DescribeStaleSecurityGroupsInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeStaleSecurityGroupsOutput, error)
	// Describes the progress of the AMI store tasks. You can describe the store tasks
	// for specified AMIs. If you don't specify the AMIs, you get a paginated list of
	// store tasks from the last 31 days. For each AMI task, the response indicates if
	// the task is InProgress, Completed, or Failed. For tasks InProgress, the response
	// shows the estimated progress as a percentage. Tasks are listed in reverse
	// chronological order. Currently, only tasks from the past 31 days can be viewed.
	// To use this API, you must have the required permissions. For more information,
	// see Permissions for storing and restoring AMIs using Amazon S3
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ami-store-restore.html#ami-s3-permissions)
	// in the Amazon Elastic Compute Cloud User Guide. For more information, see Store
	// and restore an AMI using Amazon S3
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ami-store-restore.html) in
	// the Amazon Elastic Compute Cloud User Guide.
	//
	DescribeStoreImageTasks(arg1 context.Context, arg2 *ec2.DescribeStoreImageTasksInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeStoreImageTasksOutput, error)
	// Describes one or more of your subnets. For more information, see Your VPC and
	// subnets (https://docs.aws.amazon.com/vpc/latest/userguide/VPC_Subnets.html) in
	// the Amazon Virtual Private Cloud User Guide.
	//
	DescribeSubnets(arg1 context.Context, arg2 *ec2.DescribeSubnetsInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeSubnetsOutput, error)
	// Describes the specified tags for your EC2 resources. For more information about
	// tags, see Tag your Amazon EC2 resources
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/Using_Tags.html) in the
	// Amazon Elastic Compute Cloud User Guide.
	//
	DescribeTags(arg1 context.Context, arg2 *ec2.DescribeTagsInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeTagsOutput, error)
	// Describes one or more Traffic Mirror filters.
	//
	DescribeTrafficMirrorFilters(arg1 context.Context, arg2 *ec2.DescribeTrafficMirrorFiltersInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeTrafficMirrorFiltersOutput, error)
	// Describes one or more Traffic Mirror sessions. By default, all Traffic Mirror
	// sessions are described. Alternatively, you can filter the results.
	//
	DescribeTrafficMirrorSessions(arg1 context.Context, arg2 *ec2.DescribeTrafficMirrorSessionsInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeTrafficMirrorSessionsOutput, error)
	// Information about one or more Traffic Mirror targets.
	//
	DescribeTrafficMirrorTargets(arg1 context.Context, arg2 *ec2.DescribeTrafficMirrorTargetsInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeTrafficMirrorTargetsOutput, error)
	// Describes one or more attachments between resources and transit gateways. By
	// default, all attachments are described. Alternatively, you can filter the
	// results by attachment ID, attachment state, resource ID, or resource owner.
	//
	DescribeTransitGatewayAttachments(arg1 context.Context, arg2 *ec2.DescribeTransitGatewayAttachmentsInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeTransitGatewayAttachmentsOutput, error)
	// Describes one or more Connect peers.
	//
	DescribeTransitGatewayConnectPeers(arg1 context.Context, arg2 *ec2.DescribeTransitGatewayConnectPeersInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeTransitGatewayConnectPeersOutput, error)
	// Describes one or more Connect attachments.
	//
	DescribeTransitGatewayConnects(arg1 context.Context, arg2 *ec2.DescribeTransitGatewayConnectsInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeTransitGatewayConnectsOutput, error)
	// Describes one or more transit gateway multicast domains.
	//
	DescribeTransitGatewayMulticastDomains(arg1 context.Context, arg2 *ec2.DescribeTransitGatewayMulticastDomainsInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeTransitGatewayMulticastDomainsOutput, error)
	// Describes your transit gateway peering attachments.
	//
	DescribeTransitGatewayPeeringAttachments(arg1 context.Context, arg2 *ec2.DescribeTransitGatewayPeeringAttachmentsInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeTransitGatewayPeeringAttachmentsOutput, error)
	// Describes one or more transit gateway route policy tables.
	//
	DescribeTransitGatewayPolicyTables(arg1 context.Context, arg2 *ec2.DescribeTransitGatewayPolicyTablesInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeTransitGatewayPolicyTablesOutput, error)
	// Describes one or more transit gateway route table advertisements.
	//
	DescribeTransitGatewayRouteTableAnnouncements(arg1 context.Context, arg2 *ec2.DescribeTransitGatewayRouteTableAnnouncementsInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeTransitGatewayRouteTableAnnouncementsOutput, error)
	// Describes one or more transit gateway route tables. By default, all transit
	// gateway route tables are described. Alternatively, you can filter the results.
	//
	DescribeTransitGatewayRouteTables(arg1 context.Context, arg2 *ec2.DescribeTransitGatewayRouteTablesInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeTransitGatewayRouteTablesOutput, error)
	// Describes one or more VPC attachments. By default, all VPC attachments are
	// described. Alternatively, you can filter the results.
	//
	DescribeTransitGatewayVpcAttachments(arg1 context.Context, arg2 *ec2.DescribeTransitGatewayVpcAttachmentsInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeTransitGatewayVpcAttachmentsOutput, error)
	// Describes one or more transit gateways. By default, all transit gateways are
	// described. Alternatively, you can filter the results.
	//
	DescribeTransitGateways(arg1 context.Context, arg2 *ec2.DescribeTransitGatewaysInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeTransitGatewaysOutput, error)
	// This API action is currently in limited preview only. If you are interested in
	// using this feature, contact your account manager. Describes one or more network
	// interface trunk associations.
	//
	DescribeTrunkInterfaceAssociations(arg1 context.Context, arg2 *ec2.DescribeTrunkInterfaceAssociationsInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeTrunkInterfaceAssociationsOutput, error)
	// Describes the specified attribute of the specified volume. You can specify only
	// one attribute at a time. For more information about EBS volumes, see Amazon EBS
	// volumes (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSVolumes.html) in
	// the Amazon Elastic Compute Cloud User Guide.
	//
	DescribeVolumeAttribute(arg1 context.Context, arg2 *ec2.DescribeVolumeAttributeInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeVolumeAttributeOutput, error)
	// Describes the status of the specified volumes. Volume status provides the result
	// of the checks performed on your volumes to determine events that can impair the
	// performance of your volumes. The performance of a volume can be affected if an
	// issue occurs on the volume's underlying host. If the volume's underlying host
	// experiences a power outage or system issue, after the system is restored, there
	// could be data inconsistencies on the volume. Volume events notify you if this
	// occurs. Volume actions notify you if any action needs to be taken in response to
	// the event. The DescribeVolumeStatus operation provides the following information
	// about the specified volumes: Status: Reflects the current status of the volume.
	// The possible values are ok, impaired , warning, or insufficient-data. If all
	// checks pass, the overall status of the volume is ok. If the check fails, the
	// overall status is impaired. If the status is insufficient-data, then the checks
	// might still be taking place on your volume at the time. We recommend that you
	// retry the request. For more information about volume status, see Monitor the
	// status of your volumes
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/monitoring-volume-status.html)
	// in the Amazon Elastic Compute Cloud User Guide. Events: Reflect the cause of a
	// volume status and might require you to take action. For example, if your volume
	// returns an impaired status, then the volume event might be
	// potential-data-inconsistency. This means that your volume has been affected by
	// an issue with the underlying host, has all I/O operations disabled, and might
	// have inconsistent data. Actions: Reflect the actions you might have to take in
	// response to an event. For example, if the status of the volume is impaired and
	// the volume event shows potential-data-inconsistency, then the action shows
	// enable-volume-io. This means that you may want to enable the I/O operations for
	// the volume by calling the EnableVolumeIO action and then check the volume for
	// data consistency. Volume status is based on the volume status checks, and does
	// not reflect the volume state. Therefore, volume status does not indicate volumes
	// in the error state (for example, when a volume is incapable of accepting I/O.)
	//
	DescribeVolumeStatus(arg1 context.Context, arg2 *ec2.DescribeVolumeStatusInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeVolumeStatusOutput, error)
	// Describes the specified EBS volumes or all of your EBS volumes. If you are
	// describing a long list of volumes, we recommend that you paginate the output to
	// make the list more manageable. The MaxResults parameter sets the maximum number
	// of results returned in a single page. If the list of results exceeds your
	// MaxResults value, then that number of results is returned along with a NextToken
	// value that can be passed to a subsequent DescribeVolumes request to retrieve the
	// remaining results. For more information about EBS volumes, see Amazon EBS
	// volumes (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSVolumes.html) in
	// the Amazon Elastic Compute Cloud User Guide.
	//
	DescribeVolumes(arg1 context.Context, arg2 *ec2.DescribeVolumesInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeVolumesOutput, error)
	// Describes the most recent volume modification request for the specified EBS
	// volumes. If a volume has never been modified, some information in the output
	// will be null. If a volume has been modified more than once, the output includes
	// only the most recent modification request. You can also use CloudWatch Events to
	// check the status of a modification to an EBS volume. For information about
	// CloudWatch Events, see the Amazon CloudWatch Events User Guide
	// (https://docs.aws.amazon.com/AmazonCloudWatch/latest/events/). For more
	// information, see Monitor the progress of volume modifications
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/monitoring-volume-modifications.html)
	// in the Amazon Elastic Compute Cloud User Guide.
	//
	DescribeVolumesModifications(arg1 context.Context, arg2 *ec2.DescribeVolumesModificationsInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeVolumesModificationsOutput, error)
	// Describes the specified attribute of the specified VPC. You can specify only one
	// attribute at a time.
	//
	DescribeVpcAttribute(arg1 context.Context, arg2 *ec2.DescribeVpcAttributeInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeVpcAttributeOutput, error)
	// Describes the ClassicLink status of one or more VPCs.
	//
	DescribeVpcClassicLink(arg1 context.Context, arg2 *ec2.DescribeVpcClassicLinkInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeVpcClassicLinkOutput, error)
	// Describes the ClassicLink DNS support status of one or more VPCs. If enabled,
	// the DNS hostname of a linked EC2-Classic instance resolves to its private IP
	// address when addressed from an instance in the VPC to which it's linked.
	// Similarly, the DNS hostname of an instance in a VPC resolves to its private IP
	// address when addressed from a linked EC2-Classic instance. For more information,
	// see ClassicLink
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/vpc-classiclink.html) in
	// the Amazon Elastic Compute Cloud User Guide.
	//
	DescribeVpcClassicLinkDnsSupport(arg1 context.Context, arg2 *ec2.DescribeVpcClassicLinkDnsSupportInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeVpcClassicLinkDnsSupportOutput, error)
	// Describes the connection notifications for VPC endpoints and VPC endpoint
	// services.
	//
	DescribeVpcEndpointConnectionNotifications(arg1 context.Context, arg2 *ec2.DescribeVpcEndpointConnectionNotificationsInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeVpcEndpointConnectionNotificationsOutput, error)
	// Describes the VPC endpoint connections to your VPC endpoint services, including
	// any endpoints that are pending your acceptance.
	//
	DescribeVpcEndpointConnections(arg1 context.Context, arg2 *ec2.DescribeVpcEndpointConnectionsInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeVpcEndpointConnectionsOutput, error)
	// Describes the VPC endpoint service configurations in your account (your
	// services).
	//
	DescribeVpcEndpointServiceConfigurations(arg1 context.Context, arg2 *ec2.DescribeVpcEndpointServiceConfigurationsInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeVpcEndpointServiceConfigurationsOutput, error)
	// Describes the principals (service consumers) that are permitted to discover your
	// VPC endpoint service.
	//
	DescribeVpcEndpointServicePermissions(arg1 context.Context, arg2 *ec2.DescribeVpcEndpointServicePermissionsInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeVpcEndpointServicePermissionsOutput, error)
	// Describes available services to which you can create a VPC endpoint. When the
	// service provider and the consumer have different accounts in multiple
	// Availability Zones, and the consumer views the VPC endpoint service information,
	// the response only includes the common Availability Zones. For example, when the
	// service provider account uses us-east-1a and us-east-1c and the consumer uses
	// us-east-1a and us-east-1b, the response includes the VPC endpoint services in
	// the common Availability Zone, us-east-1a.
	//
	DescribeVpcEndpointServices(arg1 context.Context, arg2 *ec2.DescribeVpcEndpointServicesInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeVpcEndpointServicesOutput, error)
	// Describes one or more of your VPC endpoints.
	//
	DescribeVpcEndpoints(arg1 context.Context, arg2 *ec2.DescribeVpcEndpointsInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeVpcEndpointsOutput, error)
	// Describes one or more of your VPC peering connections.
	//
	DescribeVpcPeeringConnections(arg1 context.Context, arg2 *ec2.DescribeVpcPeeringConnectionsInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeVpcPeeringConnectionsOutput, error)
	// Describes one or more of your VPCs.
	//
	DescribeVpcs(arg1 context.Context, arg2 *ec2.DescribeVpcsInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeVpcsOutput, error)
	// Describes one or more of your VPN connections. For more information, see Amazon
	// Web Services Site-to-Site VPN
	// (https://docs.aws.amazon.com/vpn/latest/s2svpn/VPC_VPN.html) in the Amazon Web
	// Services Site-to-Site VPN User Guide.
	//
	DescribeVpnConnections(arg1 context.Context, arg2 *ec2.DescribeVpnConnectionsInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeVpnConnectionsOutput, error)
	// Describes one or more of your virtual private gateways. For more information,
	// see Amazon Web Services Site-to-Site VPN
	// (https://docs.aws.amazon.com/vpn/latest/s2svpn/VPC_VPN.html) in the Amazon Web
	// Services Site-to-Site VPN User Guide.
	//
	DescribeVpnGateways(arg1 context.Context, arg2 *ec2.DescribeVpnGatewaysInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeVpnGatewaysOutput, error)
	// Unlinks (detaches) a linked EC2-Classic instance from a VPC. After the instance
	// has been unlinked, the VPC security groups are no longer associated with it. An
	// instance is automatically unlinked from a VPC when it's stopped.
	//
	DetachClassicLinkVpc(arg1 context.Context, arg2 *ec2.DetachClassicLinkVpcInput, arg3 ...func(*ec2.Options)) (*ec2.DetachClassicLinkVpcOutput, error)
	// Detaches an internet gateway from a VPC, disabling connectivity between the
	// internet and the VPC. The VPC must not contain any running instances with
	// Elastic IP addresses or public IPv4 addresses.
	//
	DetachInternetGateway(arg1 context.Context, arg2 *ec2.DetachInternetGatewayInput, arg3 ...func(*ec2.Options)) (*ec2.DetachInternetGatewayOutput, error)
	// Detaches a network interface from an instance.
	//
	DetachNetworkInterface(arg1 context.Context, arg2 *ec2.DetachNetworkInterfaceInput, arg3 ...func(*ec2.Options)) (*ec2.DetachNetworkInterfaceOutput, error)
	// Detaches an EBS volume from an instance. Make sure to unmount any file systems
	// on the device within your operating system before detaching the volume. Failure
	// to do so can result in the volume becoming stuck in the busy state while
	// detaching. If this happens, detachment can be delayed indefinitely until you
	// unmount the volume, force detachment, reboot the instance, or all three. If an
	// EBS volume is the root device of an instance, it can't be detached while the
	// instance is running. To detach the root volume, stop the instance first. When a
	// volume with an Amazon Web Services Marketplace product code is detached from an
	// instance, the product code is no longer associated with the instance. For more
	// information, see Detach an Amazon EBS volume
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ebs-detaching-volume.html)
	// in the Amazon Elastic Compute Cloud User Guide.
	//
	DetachVolume(arg1 context.Context, arg2 *ec2.DetachVolumeInput, arg3 ...func(*ec2.Options)) (*ec2.DetachVolumeOutput, error)
	// Detaches a virtual private gateway from a VPC. You do this if you're planning to
	// turn off the VPC and not use it anymore. You can confirm a virtual private
	// gateway has been completely detached from a VPC by describing the virtual
	// private gateway (any attachments to the virtual private gateway are also
	// described). You must wait for the attachment's state to switch to detached
	// before you can delete the VPC or attach a different VPC to the virtual private
	// gateway.
	//
	DetachVpnGateway(arg1 context.Context, arg2 *ec2.DetachVpnGatewayInput, arg3 ...func(*ec2.Options)) (*ec2.DetachVpnGatewayOutput, error)
	// Disables EBS encryption by default for your account in the current Region. After
	// you disable encryption by default, you can still create encrypted volumes by
	// enabling encryption when you create each volume. Disabling encryption by default
	// does not change the encryption status of your existing volumes. For more
	// information, see Amazon EBS encryption
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSEncryption.html) in the
	// Amazon Elastic Compute Cloud User Guide.
	//
	DisableEbsEncryptionByDefault(arg1 context.Context, arg2 *ec2.DisableEbsEncryptionByDefaultInput, arg3 ...func(*ec2.Options)) (*ec2.DisableEbsEncryptionByDefaultOutput, error)
	// Discontinue faster launching for a Windows AMI, and clean up existing
	// pre-provisioned snapshots. When you disable faster launching, the AMI uses the
	// standard launch process for each instance. All pre-provisioned snapshots must be
	// removed before you can enable faster launching again. To change these settings,
	// you must own the AMI.
	//
	DisableFastLaunch(arg1 context.Context, arg2 *ec2.DisableFastLaunchInput, arg3 ...func(*ec2.Options)) (*ec2.DisableFastLaunchOutput, error)
	// Disables fast snapshot restores for the specified snapshots in the specified
	// Availability Zones.
	//
	DisableFastSnapshotRestores(arg1 context.Context, arg2 *ec2.DisableFastSnapshotRestoresInput, arg3 ...func(*ec2.Options)) (*ec2.DisableFastSnapshotRestoresOutput, error)
	// Cancels the deprecation of the specified AMI. For more information, see
	// Deprecate an AMI
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ami-deprecate.html) in the
	// Amazon Elastic Compute Cloud User Guide.
	//
	DisableImageDeprecation(arg1 context.Context, arg2 *ec2.DisableImageDeprecationInput, arg3 ...func(*ec2.Options)) (*ec2.DisableImageDeprecationOutput, error)
	// Disable the IPAM account. For more information, see Enable integration with
	// Organizations
	// (https://docs.aws.amazon.com/vpc/latest/ipam/enable-integ-ipam.html) in the
	// Amazon VPC IPAM User Guide.
	//
	DisableIpamOrganizationAdminAccount(arg1 context.Context, arg2 *ec2.DisableIpamOrganizationAdminAccountInput, arg3 ...func(*ec2.Options)) (*ec2.DisableIpamOrganizationAdminAccountOutput, error)
	// Disables access to the EC2 serial console of all instances for your account. By
	// default, access to the EC2 serial console is disabled for your account. For more
	// information, see Manage account access to the EC2 serial console
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/configure-access-to-serial-console.html#serial-console-account-access)
	// in the Amazon EC2 User Guide.
	//
	DisableSerialConsoleAccess(arg1 context.Context, arg2 *ec2.DisableSerialConsoleAccessInput, arg3 ...func(*ec2.Options)) (*ec2.DisableSerialConsoleAccessOutput, error)
	// Disables the specified resource attachment from propagating routes to the
	// specified propagation route table.
	//
	DisableTransitGatewayRouteTablePropagation(arg1 context.Context, arg2 *ec2.DisableTransitGatewayRouteTablePropagationInput, arg3 ...func(*ec2.Options)) (*ec2.DisableTransitGatewayRouteTablePropagationOutput, error)
	// Disables a virtual private gateway (VGW) from propagating routes to a specified
	// route table of a VPC.
	//
	DisableVgwRoutePropagation(arg1 context.Context, arg2 *ec2.DisableVgwRoutePropagationInput, arg3 ...func(*ec2.Options)) (*ec2.DisableVgwRoutePropagationOutput, error)
	// Disables ClassicLink for a VPC. You cannot disable ClassicLink for a VPC that
	// has EC2-Classic instances linked to it.
	//
	DisableVpcClassicLink(arg1 context.Context, arg2 *ec2.DisableVpcClassicLinkInput, arg3 ...func(*ec2.Options)) (*ec2.DisableVpcClassicLinkOutput, error)
	// Disables ClassicLink DNS support for a VPC. If disabled, DNS hostnames resolve
	// to public IP addresses when addressed between a linked EC2-Classic instance and
	// instances in the VPC to which it's linked. For more information, see ClassicLink
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/vpc-classiclink.html) in
	// the Amazon Elastic Compute Cloud User Guide. You must specify a VPC ID in the
	// request.
	//
	DisableVpcClassicLinkDnsSupport(arg1 context.Context, arg2 *ec2.DisableVpcClassicLinkDnsSupportInput, arg3 ...func(*ec2.Options)) (*ec2.DisableVpcClassicLinkDnsSupportOutput, error)
	// Disassociates an Elastic IP address from the instance or network interface it's
	// associated with. An Elastic IP address is for use in either the EC2-Classic
	// platform or in a VPC. For more information, see Elastic IP Addresses
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/elastic-ip-addresses-eip.html)
	// in the Amazon Elastic Compute Cloud User Guide. This is an idempotent operation.
	// If you perform the operation more than once, Amazon EC2 doesn't return an error.
	//
	DisassociateAddress(arg1 context.Context, arg2 *ec2.DisassociateAddressInput, arg3 ...func(*ec2.Options)) (*ec2.DisassociateAddressOutput, error)
	// Disassociates a target network from the specified Client VPN endpoint. When you
	// disassociate the last target network from a Client VPN, the following
	// happens:
	//
	// * The route that was automatically added for the VPC is deleted
	//
	// * All
	// active client connections are terminated
	//
	// * New client connections are
	// disallowed
	//
	// * The Client VPN endpoint's status changes to pending-associate
	//
	DisassociateClientVpnTargetNetwork(arg1 context.Context, arg2 *ec2.DisassociateClientVpnTargetNetworkInput, arg3 ...func(*ec2.Options)) (*ec2.DisassociateClientVpnTargetNetworkOutput, error)
	// Disassociates an IAM role from an Certificate Manager (ACM) certificate.
	// Disassociating an IAM role from an ACM certificate removes the Amazon S3 object
	// that contains the certificate, certificate chain, and encrypted private key from
	// the Amazon S3 bucket. It also revokes the IAM role's permission to use the KMS
	// key used to encrypt the private key. This effectively revokes the role's
	// permission to use the certificate.
	//
	DisassociateEnclaveCertificateIamRole(arg1 context.Context, arg2 *ec2.DisassociateEnclaveCertificateIamRoleInput, arg3 ...func(*ec2.Options)) (*ec2.DisassociateEnclaveCertificateIamRoleOutput, error)
	// Disassociates an IAM instance profile from a running or stopped instance. Use
	// DescribeIamInstanceProfileAssociations to get the association ID.
	//
	DisassociateIamInstanceProfile(arg1 context.Context, arg2 *ec2.DisassociateIamInstanceProfileInput, arg3 ...func(*ec2.Options)) (*ec2.DisassociateIamInstanceProfileOutput, error)
	// Disassociates one or more targets from an event window. For more information,
	// see Define event windows for scheduled events
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/event-windows.html) in the
	// Amazon EC2 User Guide.
	//
	DisassociateInstanceEventWindow(arg1 context.Context, arg2 *ec2.DisassociateInstanceEventWindowInput, arg3 ...func(*ec2.Options)) (*ec2.DisassociateInstanceEventWindowOutput, error)
	// Disassociates a subnet or gateway from a route table. After you perform this
	// action, the subnet no longer uses the routes in the route table. Instead, it
	// uses the routes in the VPC's main route table. For more information about route
	// tables, see Route tables
	// (https://docs.aws.amazon.com/vpc/latest/userguide/VPC_Route_Tables.html) in the
	// Amazon Virtual Private Cloud User Guide.
	//
	DisassociateRouteTable(arg1 context.Context, arg2 *ec2.DisassociateRouteTableInput, arg3 ...func(*ec2.Options)) (*ec2.DisassociateRouteTableOutput, error)
	// Disassociates a CIDR block from a subnet. Currently, you can disassociate an
	// IPv6 CIDR block only. You must detach or delete all gateways and resources that
	// are associated with the CIDR block before you can disassociate it.
	//
	DisassociateSubnetCidrBlock(arg1 context.Context, arg2 *ec2.DisassociateSubnetCidrBlockInput, arg3 ...func(*ec2.Options)) (*ec2.DisassociateSubnetCidrBlockOutput, error)
	// Disassociates the specified subnets from the transit gateway multicast domain.
	//
	DisassociateTransitGatewayMulticastDomain(arg1 context.Context, arg2 *ec2.DisassociateTransitGatewayMulticastDomainInput, arg3 ...func(*ec2.Options)) (*ec2.DisassociateTransitGatewayMulticastDomainOutput, error)
	// Removes the association between an an attachment and a policy table.
	//
	DisassociateTransitGatewayPolicyTable(arg1 context.Context, arg2 *ec2.DisassociateTransitGatewayPolicyTableInput, arg3 ...func(*ec2.Options)) (*ec2.DisassociateTransitGatewayPolicyTableOutput, error)
	// Disassociates a resource attachment from a transit gateway route table.
	//
	DisassociateTransitGatewayRouteTable(arg1 context.Context, arg2 *ec2.DisassociateTransitGatewayRouteTableInput, arg3 ...func(*ec2.Options)) (*ec2.DisassociateTransitGatewayRouteTableOutput, error)
	// This API action is currently in limited preview only. If you are interested in
	// using this feature, contact your account manager. Removes an association between
	// a branch network interface with a trunk network interface.
	//
	DisassociateTrunkInterface(arg1 context.Context, arg2 *ec2.DisassociateTrunkInterfaceInput, arg3 ...func(*ec2.Options)) (*ec2.DisassociateTrunkInterfaceOutput, error)
	// Disassociates a CIDR block from a VPC. To disassociate the CIDR block, you must
	// specify its association ID. You can get the association ID by using
	// DescribeVpcs. You must detach or delete all gateways and resources that are
	// associated with the CIDR block before you can disassociate it. You cannot
	// disassociate the CIDR block with which you originally created the VPC (the
	// primary CIDR block).
	//
	DisassociateVpcCidrBlock(arg1 context.Context, arg2 *ec2.DisassociateVpcCidrBlockInput, arg3 ...func(*ec2.Options)) (*ec2.DisassociateVpcCidrBlockOutput, error)
	// Enables EBS encryption by default for your account in the current Region. After
	// you enable encryption by default, the EBS volumes that you create are always
	// encrypted, either using the default KMS key or the KMS key that you specified
	// when you created each volume. For more information, see Amazon EBS encryption
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSEncryption.html) in the
	// Amazon Elastic Compute Cloud User Guide. You can specify the default KMS key for
	// encryption by default using ModifyEbsDefaultKmsKeyId or ResetEbsDefaultKmsKeyId.
	// Enabling encryption by default has no effect on the encryption status of your
	// existing volumes. After you enable encryption by default, you can no longer
	// launch instances using instance types that do not support encryption. For more
	// information, see Supported instance types
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSEncryption.html#EBSEncryption_supported_instances).
	//
	EnableEbsEncryptionByDefault(arg1 context.Context, arg2 *ec2.EnableEbsEncryptionByDefaultInput, arg3 ...func(*ec2.Options)) (*ec2.EnableEbsEncryptionByDefaultOutput, error)
	// When you enable faster launching for a Windows AMI, images are pre-provisioned,
	// using snapshots to launch instances up to 65%!f(MISSING)aster. To create the optimized
	// Windows image, Amazon EC2 launches an instance and runs through Sysprep steps,
	// rebooting as required. Then it creates a set of reserved snapshots that are used
	// for subsequent launches. The reserved snapshots are automatically replenished as
	// they are used, depending on your settings for launch frequency. To change these
	// settings, you must own the AMI.
	//
	EnableFastLaunch(arg1 context.Context, arg2 *ec2.EnableFastLaunchInput, arg3 ...func(*ec2.Options)) (*ec2.EnableFastLaunchOutput, error)
	// Enables fast snapshot restores for the specified snapshots in the specified
	// Availability Zones. You get the full benefit of fast snapshot restores after
	// they enter the enabled state. To get the current state of fast snapshot
	// restores, use DescribeFastSnapshotRestores. To disable fast snapshot restores,
	// use DisableFastSnapshotRestores. For more information, see Amazon EBS fast
	// snapshot restore
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ebs-fast-snapshot-restore.html)
	// in the Amazon Elastic Compute Cloud User Guide.
	//
	EnableFastSnapshotRestores(arg1 context.Context, arg2 *ec2.EnableFastSnapshotRestoresInput, arg3 ...func(*ec2.Options)) (*ec2.EnableFastSnapshotRestoresOutput, error)
	// Enables deprecation of the specified AMI at the specified date and time. For
	// more information, see Deprecate an AMI
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ami-deprecate.html) in the
	// Amazon Elastic Compute Cloud User Guide.
	//
	EnableImageDeprecation(arg1 context.Context, arg2 *ec2.EnableImageDeprecationInput, arg3 ...func(*ec2.Options)) (*ec2.EnableImageDeprecationOutput, error)
	// Enable an Organizations member account as the IPAM admin account. You cannot
	// select the Organizations management account as the IPAM admin account. For more
	// information, see Enable integration with Organizations
	// (https://docs.aws.amazon.com/vpc/latest/ipam/enable-integ-ipam.html) in the
	// Amazon VPC IPAM User Guide.
	//
	EnableIpamOrganizationAdminAccount(arg1 context.Context, arg2 *ec2.EnableIpamOrganizationAdminAccountInput, arg3 ...func(*ec2.Options)) (*ec2.EnableIpamOrganizationAdminAccountOutput, error)
	// Enables access to the EC2 serial console of all instances for your account. By
	// default, access to the EC2 serial console is disabled for your account. For more
	// information, see Manage account access to the EC2 serial console
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/configure-access-to-serial-console.html#serial-console-account-access)
	// in the Amazon EC2 User Guide.
	//
	EnableSerialConsoleAccess(arg1 context.Context, arg2 *ec2.EnableSerialConsoleAccessInput, arg3 ...func(*ec2.Options)) (*ec2.EnableSerialConsoleAccessOutput, error)
	// Enables the specified attachment to propagate routes to the specified
	// propagation route table.
	//
	EnableTransitGatewayRouteTablePropagation(arg1 context.Context, arg2 *ec2.EnableTransitGatewayRouteTablePropagationInput, arg3 ...func(*ec2.Options)) (*ec2.EnableTransitGatewayRouteTablePropagationOutput, error)
	// Enables a virtual private gateway (VGW) to propagate routes to the specified
	// route table of a VPC.
	//
	EnableVgwRoutePropagation(arg1 context.Context, arg2 *ec2.EnableVgwRoutePropagationInput, arg3 ...func(*ec2.Options)) (*ec2.EnableVgwRoutePropagationOutput, error)
	// Enables I/O operations for a volume that had I/O operations disabled because the
	// data on the volume was potentially inconsistent.
	//
	EnableVolumeIO(arg1 context.Context, arg2 *ec2.EnableVolumeIOInput, arg3 ...func(*ec2.Options)) (*ec2.EnableVolumeIOOutput, error)
	// Enables a VPC for ClassicLink. You can then link EC2-Classic instances to your
	// ClassicLink-enabled VPC to allow communication over private IP addresses. You
	// cannot enable your VPC for ClassicLink if any of your VPC route tables have
	// existing routes for address ranges within the 10.0.0.0/8 IP address range,
	// excluding local routes for VPCs in the 10.0.0.0/16 and 10.1.0.0/16 IP address
	// ranges. For more information, see ClassicLink
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/vpc-classiclink.html) in
	// the Amazon Elastic Compute Cloud User Guide.
	//
	EnableVpcClassicLink(arg1 context.Context, arg2 *ec2.EnableVpcClassicLinkInput, arg3 ...func(*ec2.Options)) (*ec2.EnableVpcClassicLinkOutput, error)
	// Enables a VPC to support DNS hostname resolution for ClassicLink. If enabled,
	// the DNS hostname of a linked EC2-Classic instance resolves to its private IP
	// address when addressed from an instance in the VPC to which it's linked.
	// Similarly, the DNS hostname of an instance in a VPC resolves to its private IP
	// address when addressed from a linked EC2-Classic instance. For more information,
	// see ClassicLink
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/vpc-classiclink.html) in
	// the Amazon Elastic Compute Cloud User Guide. You must specify a VPC ID in the
	// request.
	//
	EnableVpcClassicLinkDnsSupport(arg1 context.Context, arg2 *ec2.EnableVpcClassicLinkDnsSupportInput, arg3 ...func(*ec2.Options)) (*ec2.EnableVpcClassicLinkDnsSupportOutput, error)
	// Downloads the client certificate revocation list for the specified Client VPN
	// endpoint.
	//
	ExportClientVpnClientCertificateRevocationList(arg1 context.Context, arg2 *ec2.ExportClientVpnClientCertificateRevocationListInput, arg3 ...func(*ec2.Options)) (*ec2.ExportClientVpnClientCertificateRevocationListOutput, error)
	// Downloads the contents of the Client VPN endpoint configuration file for the
	// specified Client VPN endpoint. The Client VPN endpoint configuration file
	// includes the Client VPN endpoint and certificate information clients need to
	// establish a connection with the Client VPN endpoint.
	//
	ExportClientVpnClientConfiguration(arg1 context.Context, arg2 *ec2.ExportClientVpnClientConfigurationInput, arg3 ...func(*ec2.Options)) (*ec2.ExportClientVpnClientConfigurationOutput, error)
	// Exports an Amazon Machine Image (AMI) to a VM file. For more information, see
	// Exporting a VM directly from an Amazon Machine Image (AMI)
	// (https://docs.aws.amazon.com/vm-import/latest/userguide/vmexport_image.html) in
	// the VM Import/Export User Guide.
	//
	ExportImage(arg1 context.Context, arg2 *ec2.ExportImageInput, arg3 ...func(*ec2.Options)) (*ec2.ExportImageOutput, error)
	// Exports routes from the specified transit gateway route table to the specified
	// S3 bucket. By default, all routes are exported. Alternatively, you can filter by
	// CIDR range. The routes are saved to the specified bucket in a JSON file. For
	// more information, see Export Route Tables to Amazon S3
	// (https://docs.aws.amazon.com/vpc/latest/tgw/tgw-route-tables.html#tgw-export-route-tables)
	// in Transit Gateways.
	//
	ExportTransitGatewayRoutes(arg1 context.Context, arg2 *ec2.ExportTransitGatewayRoutesInput, arg3 ...func(*ec2.Options)) (*ec2.ExportTransitGatewayRoutesOutput, error)
	// Returns the IAM roles that are associated with the specified ACM (ACM)
	// certificate. It also returns the name of the Amazon S3 bucket and the Amazon S3
	// object key where the certificate, certificate chain, and encrypted private key
	// bundle are stored, and the ARN of the KMS key that's used to encrypt the private
	// key.
	//
	GetAssociatedEnclaveCertificateIamRoles(arg1 context.Context, arg2 *ec2.GetAssociatedEnclaveCertificateIamRolesInput, arg3 ...func(*ec2.Options)) (*ec2.GetAssociatedEnclaveCertificateIamRolesOutput, error)
	// Gets information about the IPv6 CIDR block associations for a specified IPv6
	// address pool.
	//
	GetAssociatedIpv6PoolCidrs(arg1 context.Context, arg2 *ec2.GetAssociatedIpv6PoolCidrsInput, arg3 ...func(*ec2.Options)) (*ec2.GetAssociatedIpv6PoolCidrsOutput, error)
	// Gets usage information about a Capacity Reservation. If the Capacity Reservation
	// is shared, it shows usage information for the Capacity Reservation owner and
	// each Amazon Web Services account that is currently using the shared capacity. If
	// the Capacity Reservation is not shared, it shows only the Capacity Reservation
	// owner's usage.
	//
	GetCapacityReservationUsage(arg1 context.Context, arg2 *ec2.GetCapacityReservationUsageInput, arg3 ...func(*ec2.Options)) (*ec2.GetCapacityReservationUsageOutput, error)
	// Describes the allocations from the specified customer-owned address pool.
	//
	GetCoipPoolUsage(arg1 context.Context, arg2 *ec2.GetCoipPoolUsageInput, arg3 ...func(*ec2.Options)) (*ec2.GetCoipPoolUsageOutput, error)
	// Gets the console output for the specified instance. For Linux instances, the
	// instance console output displays the exact console output that would normally be
	// displayed on a physical monitor attached to a computer. For Windows instances,
	// the instance console output includes the last three system event log errors. By
	// default, the console output returns buffered information that was posted shortly
	// after an instance transition state (start, stop, reboot, or terminate). This
	// information is available for at least one hour after the most recent post. Only
	// the most recent 64 KB of console output is available. You can optionally
	// retrieve the latest serial console output at any time during the instance
	// lifecycle. This option is supported on instance types that use the Nitro
	// hypervisor. For more information, see Instance console output
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instance-console.html#instance-console-console-output)
	// in the Amazon EC2 User Guide.
	//
	GetConsoleOutput(arg1 context.Context, arg2 *ec2.GetConsoleOutputInput, arg3 ...func(*ec2.Options)) (*ec2.GetConsoleOutputOutput, error)
	// Retrieve a JPG-format screenshot of a running instance to help with
	// troubleshooting. The returned content is Base64-encoded.
	//
	GetConsoleScreenshot(arg1 context.Context, arg2 *ec2.GetConsoleScreenshotInput, arg3 ...func(*ec2.Options)) (*ec2.GetConsoleScreenshotOutput, error)
	// Describes the default credit option for CPU usage of a burstable performance
	// instance family. For more information, see Burstable performance instances
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/burstable-performance-instances.html)
	// in the Amazon EC2 User Guide.
	//
	GetDefaultCreditSpecification(arg1 context.Context, arg2 *ec2.GetDefaultCreditSpecificationInput, arg3 ...func(*ec2.Options)) (*ec2.GetDefaultCreditSpecificationOutput, error)
	// Describes the default KMS key for EBS encryption by default for your account in
	// this Region. You can change the default KMS key for encryption by default using
	// ModifyEbsDefaultKmsKeyId or ResetEbsDefaultKmsKeyId. For more information, see
	// Amazon EBS encryption
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSEncryption.html) in the
	// Amazon Elastic Compute Cloud User Guide.
	//
	GetEbsDefaultKmsKeyId(arg1 context.Context, arg2 *ec2.GetEbsDefaultKmsKeyIdInput, arg3 ...func(*ec2.Options)) (*ec2.GetEbsDefaultKmsKeyIdOutput, error)
	// Describes whether EBS encryption by default is enabled for your account in the
	// current Region. For more information, see Amazon EBS encryption
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSEncryption.html) in the
	// Amazon Elastic Compute Cloud User Guide.
	//
	GetEbsEncryptionByDefault(arg1 context.Context, arg2 *ec2.GetEbsEncryptionByDefaultInput, arg3 ...func(*ec2.Options)) (*ec2.GetEbsEncryptionByDefaultOutput, error)
	// Generates a CloudFormation template that streamlines and automates the
	// integration of VPC flow logs with Amazon Athena. This make it easier for you to
	// query and gain insights from VPC flow logs data. Based on the information that
	// you provide, we configure resources in the template to do the following:
	//
	// *
	// Create a table in Athena that maps fields to a custom log format
	//
	// * Create a
	// Lambda function that updates the table with new partitions on a daily, weekly,
	// or monthly basis
	//
	// * Create a table partitioned between two timestamps in the
	// past
	//
	// * Create a set of named queries in Athena that you can use to get started
	// quickly
	//
	GetFlowLogsIntegrationTemplate(arg1 context.Context, arg2 *ec2.GetFlowLogsIntegrationTemplateInput, arg3 ...func(*ec2.Options)) (*ec2.GetFlowLogsIntegrationTemplateOutput, error)
	// Lists the resource groups to which a Capacity Reservation has been added.
	//
	GetGroupsForCapacityReservation(arg1 context.Context, arg2 *ec2.GetGroupsForCapacityReservationInput, arg3 ...func(*ec2.Options)) (*ec2.GetGroupsForCapacityReservationOutput, error)
	// Preview a reservation purchase with configurations that match those of your
	// Dedicated Host. You must have active Dedicated Hosts in your account before you
	// purchase a reservation. This is a preview of the PurchaseHostReservation action
	// and does not result in the offering being purchased.
	//
	GetHostReservationPurchasePreview(arg1 context.Context, arg2 *ec2.GetHostReservationPurchasePreviewInput, arg3 ...func(*ec2.Options)) (*ec2.GetHostReservationPurchasePreviewOutput, error)
	// Returns a list of instance types with the specified instance attributes. You can
	// use the response to preview the instance types without launching instances. Note
	// that the response does not consider capacity. When you specify multiple
	// parameters, you get instance types that satisfy all of the specified parameters.
	// If you specify multiple values for a parameter, you get instance types that
	// satisfy any of the specified values. For more information, see Preview instance
	// types with specified attributes
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/spot-fleet-attribute-based-instance-type-selection.html#spotfleet-get-instance-types-from-instance-requirements),
	// Attribute-based instance type selection for EC2 Fleet
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-fleet-attribute-based-instance-type-selection.html),
	// Attribute-based instance type selection for Spot Fleet
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/spot-fleet-attribute-based-instance-type-selection.html),
	// and Spot placement score
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/spot-placement-score.html)
	// in the Amazon EC2 User Guide, and Creating an Auto Scaling group using
	// attribute-based instance type selection
	// (https://docs.aws.amazon.com/autoscaling/ec2/userguide/create-asg-instance-type-requirements.html)
	// in the Amazon EC2 Auto Scaling User Guide.
	//
	GetInstanceTypesFromInstanceRequirements(arg1 context.Context, arg2 *ec2.GetInstanceTypesFromInstanceRequirementsInput, arg3 ...func(*ec2.Options)) (*ec2.GetInstanceTypesFromInstanceRequirementsOutput, error)
	// A binary representation of the UEFI variable store. Only non-volatile variables
	// are stored. This is a base64 encoded and zlib compressed binary value that must
	// be properly encoded. When you use register-image
	// (https://docs.aws.amazon.com/cli/latest/reference/ec2/register-image.html) to
	// create an AMI, you can create an exact copy of your variable store by passing
	// the UEFI data in the UefiData parameter. You can modify the UEFI data by using
	// the python-uefivars tool (https://github.com/awslabs/python-uefivars) on GitHub.
	// You can use the tool to convert the UEFI data into a human-readable format
	// (JSON), which you can inspect and modify, and then convert back into the binary
	// format to use with register-image. For more information, see UEFI Secure Boot
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/uefi-secure-boot.html) in
	// the Amazon EC2 User Guide.
	//
	GetInstanceUefiData(arg1 context.Context, arg2 *ec2.GetInstanceUefiDataInput, arg3 ...func(*ec2.Options)) (*ec2.GetInstanceUefiDataOutput, error)
	// Retrieve historical information about a CIDR within an IPAM scope. For more
	// information, see View the history of IP addresses
	// (https://docs.aws.amazon.com/vpc/latest/ipam/view-history-cidr-ipam.html) in the
	// Amazon VPC IPAM User Guide.
	//
	GetIpamAddressHistory(arg1 context.Context, arg2 *ec2.GetIpamAddressHistoryInput, arg3 ...func(*ec2.Options)) (*ec2.GetIpamAddressHistoryOutput, error)
	// Get a list of all the CIDR allocations in an IPAM pool.
	//
	GetIpamPoolAllocations(arg1 context.Context, arg2 *ec2.GetIpamPoolAllocationsInput, arg3 ...func(*ec2.Options)) (*ec2.GetIpamPoolAllocationsOutput, error)
	// Get the CIDRs provisioned to an IPAM pool.
	//
	GetIpamPoolCidrs(arg1 context.Context, arg2 *ec2.GetIpamPoolCidrsInput, arg3 ...func(*ec2.Options)) (*ec2.GetIpamPoolCidrsOutput, error)
	// Get information about the resources in a scope.
	//
	GetIpamResourceCidrs(arg1 context.Context, arg2 *ec2.GetIpamResourceCidrsInput, arg3 ...func(*ec2.Options)) (*ec2.GetIpamResourceCidrsOutput, error)
	// Retrieves the configuration data of the specified instance. You can use this
	// data to create a launch template. This action calls on other describe actions to
	// get instance information. Depending on your instance configuration, you may need
	// to allow the following actions in your IAM policy: DescribeSpotInstanceRequests,
	// DescribeInstanceCreditSpecifications, DescribeVolumes,
	// DescribeInstanceAttribute, and DescribeElasticGpus. Or, you can allow describe*
	// depending on your instance requirements.
	//
	GetLaunchTemplateData(arg1 context.Context, arg2 *ec2.GetLaunchTemplateDataInput, arg3 ...func(*ec2.Options)) (*ec2.GetLaunchTemplateDataOutput, error)
	// Gets information about the resources that are associated with the specified
	// managed prefix list.
	//
	GetManagedPrefixListAssociations(arg1 context.Context, arg2 *ec2.GetManagedPrefixListAssociationsInput, arg3 ...func(*ec2.Options)) (*ec2.GetManagedPrefixListAssociationsOutput, error)
	// Gets information about the entries for a specified managed prefix list.
	//
	GetManagedPrefixListEntries(arg1 context.Context, arg2 *ec2.GetManagedPrefixListEntriesInput, arg3 ...func(*ec2.Options)) (*ec2.GetManagedPrefixListEntriesOutput, error)
	// Gets the findings for the specified Network Access Scope analysis.
	//
	GetNetworkInsightsAccessScopeAnalysisFindings(arg1 context.Context, arg2 *ec2.GetNetworkInsightsAccessScopeAnalysisFindingsInput, arg3 ...func(*ec2.Options)) (*ec2.GetNetworkInsightsAccessScopeAnalysisFindingsOutput, error)
	// Gets the content for the specified Network Access Scope.
	//
	GetNetworkInsightsAccessScopeContent(arg1 context.Context, arg2 *ec2.GetNetworkInsightsAccessScopeContentInput, arg3 ...func(*ec2.Options)) (*ec2.GetNetworkInsightsAccessScopeContentOutput, error)
	// Retrieves the encrypted administrator password for a running Windows instance.
	// The Windows password is generated at boot by the EC2Config service or EC2Launch
	// scripts (Windows Server 2016 and later). This usually only happens the first
	// time an instance is launched. For more information, see EC2Config
	// (https://docs.aws.amazon.com/AWSEC2/latest/WindowsGuide/UsingConfig_WinAMI.html)
	// and EC2Launch
	// (https://docs.aws.amazon.com/AWSEC2/latest/WindowsGuide/ec2launch.html) in the
	// Amazon EC2 User Guide. For the EC2Config service, the password is not generated
	// for rebundled AMIs unless Ec2SetPassword is enabled before bundling. The
	// password is encrypted using the key pair that you specified when you launched
	// the instance. You must provide the corresponding key pair file. When you launch
	// an instance, password generation and encryption may take a few minutes. If you
	// try to retrieve the password before it's available, the output returns an empty
	// string. We recommend that you wait up to 15 minutes after launching an instance
	// before trying to retrieve the generated password.
	//
	GetPasswordData(arg1 context.Context, arg2 *ec2.GetPasswordDataInput, arg3 ...func(*ec2.Options)) (*ec2.GetPasswordDataOutput, error)
	// Returns a quote and exchange information for exchanging one or more specified
	// Convertible Reserved Instances for a new Convertible Reserved Instance. If the
	// exchange cannot be performed, the reason is returned in the response. Use
	// AcceptReservedInstancesExchangeQuote to perform the exchange.
	//
	GetReservedInstancesExchangeQuote(arg1 context.Context, arg2 *ec2.GetReservedInstancesExchangeQuoteInput, arg3 ...func(*ec2.Options)) (*ec2.GetReservedInstancesExchangeQuoteOutput, error)
	// Retrieves the access status of your account to the EC2 serial console of all
	// instances. By default, access to the EC2 serial console is disabled for your
	// account. For more information, see Manage account access to the EC2 serial
	// console
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/configure-access-to-serial-console.html#serial-console-account-access)
	// in the Amazon EC2 User Guide.
	//
	GetSerialConsoleAccessStatus(arg1 context.Context, arg2 *ec2.GetSerialConsoleAccessStatusInput, arg3 ...func(*ec2.Options)) (*ec2.GetSerialConsoleAccessStatusOutput, error)
	// Calculates the Spot placement score for a Region or Availability Zone based on
	// the specified target capacity and compute requirements. You can specify your
	// compute requirements either by using InstanceRequirementsWithMetadata and
	// letting Amazon EC2 choose the optimal instance types to fulfill your Spot
	// request, or you can specify the instance types by using InstanceTypes. For more
	// information, see Spot placement score
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/spot-placement-score.html)
	// in the Amazon EC2 User Guide.
	//
	GetSpotPlacementScores(arg1 context.Context, arg2 *ec2.GetSpotPlacementScoresInput, arg3 ...func(*ec2.Options)) (*ec2.GetSpotPlacementScoresOutput, error)
	// Gets information about the subnet CIDR reservations.
	//
	GetSubnetCidrReservations(arg1 context.Context, arg2 *ec2.GetSubnetCidrReservationsInput, arg3 ...func(*ec2.Options)) (*ec2.GetSubnetCidrReservationsOutput, error)
	// Lists the route tables to which the specified resource attachment propagates
	// routes.
	//
	GetTransitGatewayAttachmentPropagations(arg1 context.Context, arg2 *ec2.GetTransitGatewayAttachmentPropagationsInput, arg3 ...func(*ec2.Options)) (*ec2.GetTransitGatewayAttachmentPropagationsOutput, error)
	// Gets information about the associations for the transit gateway multicast
	// domain.
	//
	GetTransitGatewayMulticastDomainAssociations(arg1 context.Context, arg2 *ec2.GetTransitGatewayMulticastDomainAssociationsInput, arg3 ...func(*ec2.Options)) (*ec2.GetTransitGatewayMulticastDomainAssociationsOutput, error)
	// Gets a list of the transit gateway policy table associations.
	//
	GetTransitGatewayPolicyTableAssociations(arg1 context.Context, arg2 *ec2.GetTransitGatewayPolicyTableAssociationsInput, arg3 ...func(*ec2.Options)) (*ec2.GetTransitGatewayPolicyTableAssociationsOutput, error)
	// Returns a list of transit gateway policy table entries.
	//
	GetTransitGatewayPolicyTableEntries(arg1 context.Context, arg2 *ec2.GetTransitGatewayPolicyTableEntriesInput, arg3 ...func(*ec2.Options)) (*ec2.GetTransitGatewayPolicyTableEntriesOutput, error)
	// Gets information about the prefix list references in a specified transit gateway
	// route table.
	//
	GetTransitGatewayPrefixListReferences(arg1 context.Context, arg2 *ec2.GetTransitGatewayPrefixListReferencesInput, arg3 ...func(*ec2.Options)) (*ec2.GetTransitGatewayPrefixListReferencesOutput, error)
	// Gets information about the associations for the specified transit gateway route
	// table.
	//
	GetTransitGatewayRouteTableAssociations(arg1 context.Context, arg2 *ec2.GetTransitGatewayRouteTableAssociationsInput, arg3 ...func(*ec2.Options)) (*ec2.GetTransitGatewayRouteTableAssociationsOutput, error)
	// Gets information about the route table propagations for the specified transit
	// gateway route table.
	//
	GetTransitGatewayRouteTablePropagations(arg1 context.Context, arg2 *ec2.GetTransitGatewayRouteTablePropagationsInput, arg3 ...func(*ec2.Options)) (*ec2.GetTransitGatewayRouteTablePropagationsOutput, error)
	// Download an Amazon Web Services-provided sample configuration file to be used
	// with the customer gateway device specified for your Site-to-Site VPN connection.
	//
	GetVpnConnectionDeviceSampleConfiguration(arg1 context.Context, arg2 *ec2.GetVpnConnectionDeviceSampleConfigurationInput, arg3 ...func(*ec2.Options)) (*ec2.GetVpnConnectionDeviceSampleConfigurationOutput, error)
	// Obtain a list of customer gateway devices for which sample configuration files
	// can be provided. The request has no additional parameters. You can also see the
	// list of device types with sample configuration files available under Your
	// customer gateway device
	// (https://docs.aws.amazon.com/vpn/latest/s2svpn/your-cgw.html) in the Amazon Web
	// Services Site-to-Site VPN User Guide.
	//
	GetVpnConnectionDeviceTypes(arg1 context.Context, arg2 *ec2.GetVpnConnectionDeviceTypesInput, arg3 ...func(*ec2.Options)) (*ec2.GetVpnConnectionDeviceTypesOutput, error)
	// Uploads a client certificate revocation list to the specified Client VPN
	// endpoint. Uploading a client certificate revocation list overwrites the existing
	// client certificate revocation list. Uploading a client certificate revocation
	// list resets existing client connections.
	//
	ImportClientVpnClientCertificateRevocationList(arg1 context.Context, arg2 *ec2.ImportClientVpnClientCertificateRevocationListInput, arg3 ...func(*ec2.Options)) (*ec2.ImportClientVpnClientCertificateRevocationListOutput, error)
	// Import single or multi-volume disk images or EBS snapshots into an Amazon
	// Machine Image (AMI). For more information, see Importing a VM as an image using
	// VM Import/Export
	// (https://docs.aws.amazon.com/vm-import/latest/userguide/vmimport-image-import.html)
	// in the VM Import/Export User Guide.
	//
	ImportImage(arg1 context.Context, arg2 *ec2.ImportImageInput, arg3 ...func(*ec2.Options)) (*ec2.ImportImageOutput, error)
	// Creates an import instance task using metadata from the specified disk image.
	// This API action supports only single-volume VMs. To import multi-volume VMs, use
	// ImportImage instead. This API action is not supported by the Command Line
	// Interface (CLI). For information about using the Amazon EC2 CLI, which is
	// deprecated, see Importing a VM to Amazon EC2
	// (https://awsdocs.s3.amazonaws.com/EC2/ec2-clt.pdf#UsingVirtualMachinesinAmazonEC2)
	// in the Amazon EC2 CLI Reference PDF file. For information about the import
	// manifest referenced by this API action, see VM Import Manifest
	// (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/manifest.html).
	//
	ImportInstance(arg1 context.Context, arg2 *ec2.ImportInstanceInput, arg3 ...func(*ec2.Options)) (*ec2.ImportInstanceOutput, error)
	// Imports the public key from an RSA or ED25519 key pair that you created with a
	// third-party tool. Compare this with CreateKeyPair, in which Amazon Web Services
	// creates the key pair and gives the keys to you (Amazon Web Services keeps a copy
	// of the public key). With ImportKeyPair, you create the key pair and give Amazon
	// Web Services just the public key. The private key is never transferred between
	// you and Amazon Web Services. For more information about key pairs, see Amazon
	// EC2 key pairs
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-key-pairs.html) in the
	// Amazon Elastic Compute Cloud User Guide.
	//
	ImportKeyPair(arg1 context.Context, arg2 *ec2.ImportKeyPairInput, arg3 ...func(*ec2.Options)) (*ec2.ImportKeyPairOutput, error)
	// Imports a disk into an EBS snapshot. For more information, see Importing a disk
	// as a snapshot using VM Import/Export
	// (https://docs.aws.amazon.com/vm-import/latest/userguide/vmimport-import-snapshot.html)
	// in the VM Import/Export User Guide.
	//
	ImportSnapshot(arg1 context.Context, arg2 *ec2.ImportSnapshotInput, arg3 ...func(*ec2.Options)) (*ec2.ImportSnapshotOutput, error)
	// Creates an import volume task using metadata from the specified disk image. This
	// API action supports only single-volume VMs. To import multi-volume VMs, use
	// ImportImage instead. To import a disk to a snapshot, use ImportSnapshot instead.
	// This API action is not supported by the Command Line Interface (CLI). For
	// information about using the Amazon EC2 CLI, which is deprecated, see Importing
	// Disks to Amazon EBS
	// (https://awsdocs.s3.amazonaws.com/EC2/ec2-clt.pdf#importing-your-volumes-into-amazon-ebs)
	// in the Amazon EC2 CLI Reference PDF file. For information about the import
	// manifest referenced by this API action, see VM Import Manifest
	// (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/manifest.html).
	//
	ImportVolume(arg1 context.Context, arg2 *ec2.ImportVolumeInput, arg3 ...func(*ec2.Options)) (*ec2.ImportVolumeOutput, error)
	// Lists one or more AMIs that are currently in the Recycle Bin. For more
	// information, see Recycle Bin
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/recycle-bin.html) in the
	// Amazon Elastic Compute Cloud User Guide.
	//
	ListImagesInRecycleBin(arg1 context.Context, arg2 *ec2.ListImagesInRecycleBinInput, arg3 ...func(*ec2.Options)) (*ec2.ListImagesInRecycleBinOutput, error)
	// Lists one or more snapshots that are currently in the Recycle Bin.
	//
	ListSnapshotsInRecycleBin(arg1 context.Context, arg2 *ec2.ListSnapshotsInRecycleBinInput, arg3 ...func(*ec2.Options)) (*ec2.ListSnapshotsInRecycleBinOutput, error)
	// Modifies an attribute of the specified Elastic IP address. For requirements, see
	// Using reverse DNS for email applications
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/elastic-ip-addresses-eip.html#Using_Elastic_Addressing_Reverse_DNS).
	//
	ModifyAddressAttribute(arg1 context.Context, arg2 *ec2.ModifyAddressAttributeInput, arg3 ...func(*ec2.Options)) (*ec2.ModifyAddressAttributeOutput, error)
	// Changes the opt-in status of the Local Zone and Wavelength Zone group for your
	// account. Use  DescribeAvailabilityZones
	// (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribeAvailabilityZones.html)
	// to view the value for GroupName.
	//
	ModifyAvailabilityZoneGroup(arg1 context.Context, arg2 *ec2.ModifyAvailabilityZoneGroupInput, arg3 ...func(*ec2.Options)) (*ec2.ModifyAvailabilityZoneGroupOutput, error)
	// Modifies a Capacity Reservation's capacity and the conditions under which it is
	// to be released. You cannot change a Capacity Reservation's instance type, EBS
	// optimization, instance store settings, platform, Availability Zone, or instance
	// eligibility. If you need to modify any of these attributes, we recommend that
	// you cancel the Capacity Reservation, and then create a new one with the required
	// attributes.
	//
	ModifyCapacityReservation(arg1 context.Context, arg2 *ec2.ModifyCapacityReservationInput, arg3 ...func(*ec2.Options)) (*ec2.ModifyCapacityReservationOutput, error)
	// Modifies a Capacity Reservation Fleet. When you modify the total target capacity
	// of a Capacity Reservation Fleet, the Fleet automatically creates new Capacity
	// Reservations, or modifies or cancels existing Capacity Reservations in the Fleet
	// to meet the new total target capacity. When you modify the end date for the
	// Fleet, the end dates for all of the individual Capacity Reservations in the
	// Fleet are updated accordingly.
	//
	ModifyCapacityReservationFleet(arg1 context.Context, arg2 *ec2.ModifyCapacityReservationFleetInput, arg3 ...func(*ec2.Options)) (*ec2.ModifyCapacityReservationFleetOutput, error)
	// Modifies the specified Client VPN endpoint. Modifying the DNS server resets
	// existing client connections.
	//
	ModifyClientVpnEndpoint(arg1 context.Context, arg2 *ec2.ModifyClientVpnEndpointInput, arg3 ...func(*ec2.Options)) (*ec2.ModifyClientVpnEndpointOutput, error)
	// Modifies the default credit option for CPU usage of burstable performance
	// instances. The default credit option is set at the account level per Amazon Web
	// Services Region, and is specified per instance family. All new burstable
	// performance instances in the account launch using the default credit option.
	// ModifyDefaultCreditSpecification is an asynchronous operation, which works at an
	// Amazon Web Services Region level and modifies the credit option for each
	// Availability Zone. All zones in a Region are updated within five minutes. But if
	// instances are launched during this operation, they might not get the new credit
	// option until the zone is updated. To verify whether the update has occurred, you
	// can call GetDefaultCreditSpecification and check DefaultCreditSpecification for
	// updates. For more information, see Burstable performance instances
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/burstable-performance-instances.html)
	// in the Amazon EC2 User Guide.
	//
	ModifyDefaultCreditSpecification(arg1 context.Context, arg2 *ec2.ModifyDefaultCreditSpecificationInput, arg3 ...func(*ec2.Options)) (*ec2.ModifyDefaultCreditSpecificationOutput, error)
	// Changes the default KMS key for EBS encryption by default for your account in
	// this Region. Amazon Web Services creates a unique Amazon Web Services managed
	// KMS key in each Region for use with encryption by default. If you change the
	// default KMS key to a symmetric customer managed KMS key, it is used instead of
	// the Amazon Web Services managed KMS key. To reset the default KMS key to the
	// Amazon Web Services managed KMS key for EBS, use ResetEbsDefaultKmsKeyId. Amazon
	// EBS does not support asymmetric KMS keys. If you delete or disable the customer
	// managed KMS key that you specified for use with encryption by default, your
	// instances will fail to launch. For more information, see Amazon EBS encryption
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSEncryption.html) in the
	// Amazon Elastic Compute Cloud User Guide.
	//
	ModifyEbsDefaultKmsKeyId(arg1 context.Context, arg2 *ec2.ModifyEbsDefaultKmsKeyIdInput, arg3 ...func(*ec2.Options)) (*ec2.ModifyEbsDefaultKmsKeyIdOutput, error)
	// Modifies the specified EC2 Fleet. You can only modify an EC2 Fleet request of
	// type maintain. While the EC2 Fleet is being modified, it is in the modifying
	// state. To scale up your EC2 Fleet, increase its target capacity. The EC2 Fleet
	// launches the additional Spot Instances according to the allocation strategy for
	// the EC2 Fleet request. If the allocation strategy is lowest-price, the EC2 Fleet
	// launches instances using the Spot Instance pool with the lowest price. If the
	// allocation strategy is diversified, the EC2 Fleet distributes the instances
	// across the Spot Instance pools. If the allocation strategy is
	// capacity-optimized, EC2 Fleet launches instances from Spot Instance pools with
	// optimal capacity for the number of instances that are launching. To scale down
	// your EC2 Fleet, decrease its target capacity. First, the EC2 Fleet cancels any
	// open requests that exceed the new target capacity. You can request that the EC2
	// Fleet terminate Spot Instances until the size of the fleet no longer exceeds the
	// new target capacity. If the allocation strategy is lowest-price, the EC2 Fleet
	// terminates the instances with the highest price per unit. If the allocation
	// strategy is capacity-optimized, the EC2 Fleet terminates the instances in the
	// Spot Instance pools that have the least available Spot Instance capacity. If the
	// allocation strategy is diversified, the EC2 Fleet terminates instances across
	// the Spot Instance pools. Alternatively, you can request that the EC2 Fleet keep
	// the fleet at its current size, but not replace any Spot Instances that are
	// interrupted or that you terminate manually. If you are finished with your EC2
	// Fleet for now, but will use it again later, you can set the target capacity to
	// 0.
	//
	ModifyFleet(arg1 context.Context, arg2 *ec2.ModifyFleetInput, arg3 ...func(*ec2.Options)) (*ec2.ModifyFleetOutput, error)
	// Modifies the specified attribute of the specified Amazon FPGA Image (AFI).
	//
	ModifyFpgaImageAttribute(arg1 context.Context, arg2 *ec2.ModifyFpgaImageAttributeInput, arg3 ...func(*ec2.Options)) (*ec2.ModifyFpgaImageAttributeOutput, error)
	// Modify the auto-placement setting of a Dedicated Host. When auto-placement is
	// enabled, any instances that you launch with a tenancy of host but without a
	// specific host ID are placed onto any available Dedicated Host in your account
	// that has auto-placement enabled. When auto-placement is disabled, you need to
	// provide a host ID to have the instance launch onto a specific host. If no host
	// ID is provided, the instance is launched onto a suitable host with
	// auto-placement enabled. You can also use this API action to modify a Dedicated
	// Host to support either multiple instance types in an instance family, or to
	// support a specific instance type only.
	//
	ModifyHosts(arg1 context.Context, arg2 *ec2.ModifyHostsInput, arg3 ...func(*ec2.Options)) (*ec2.ModifyHostsOutput, error)
	// Modifies the ID format for the specified resource on a per-Region basis. You can
	// specify that resources should receive longer IDs (17-character IDs) when they
	// are created. This request can only be used to modify longer ID settings for
	// resource types that are within the opt-in period. Resources currently in their
	// opt-in period include: bundle | conversion-task | customer-gateway |
	// dhcp-options | elastic-ip-allocation | elastic-ip-association | export-task |
	// flow-log | image | import-task | internet-gateway | network-acl |
	// network-acl-association | network-interface | network-interface-attachment |
	// prefix-list | route-table | route-table-association | security-group | subnet |
	// subnet-cidr-block-association | vpc | vpc-cidr-block-association | vpc-endpoint
	// | vpc-peering-connection | vpn-connection | vpn-gateway. This setting applies to
	// the IAM user who makes the request; it does not apply to the entire Amazon Web
	// Services account. By default, an IAM user defaults to the same settings as the
	// root user. If you're using this action as the root user, then these settings
	// apply to the entire account, unless an IAM user explicitly overrides these
	// settings for themselves. For more information, see Resource IDs
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/resource-ids.html) in the
	// Amazon Elastic Compute Cloud User Guide. Resources created with longer IDs are
	// visible to all IAM roles and users, regardless of these settings and provided
	// that they have permission to use the relevant Describe command for the resource
	// type.
	//
	ModifyIdFormat(arg1 context.Context, arg2 *ec2.ModifyIdFormatInput, arg3 ...func(*ec2.Options)) (*ec2.ModifyIdFormatOutput, error)
	// Modifies the ID format of a resource for a specified IAM user, IAM role, or the
	// root user for an account; or all IAM users, IAM roles, and the root user for an
	// account. You can specify that resources should receive longer IDs (17-character
	// IDs) when they are created. This request can only be used to modify longer ID
	// settings for resource types that are within the opt-in period. Resources
	// currently in their opt-in period include: bundle | conversion-task |
	// customer-gateway | dhcp-options | elastic-ip-allocation | elastic-ip-association
	// | export-task | flow-log | image | import-task | internet-gateway | network-acl
	// | network-acl-association | network-interface | network-interface-attachment |
	// prefix-list | route-table | route-table-association | security-group | subnet |
	// subnet-cidr-block-association | vpc | vpc-cidr-block-association | vpc-endpoint
	// | vpc-peering-connection | vpn-connection | vpn-gateway. For more information,
	// see Resource IDs
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/resource-ids.html) in the
	// Amazon Elastic Compute Cloud User Guide. This setting applies to the principal
	// specified in the request; it does not apply to the principal that makes the
	// request. Resources created with longer IDs are visible to all IAM roles and
	// users, regardless of these settings and provided that they have permission to
	// use the relevant Describe command for the resource type.
	//
	ModifyIdentityIdFormat(arg1 context.Context, arg2 *ec2.ModifyIdentityIdFormatInput, arg3 ...func(*ec2.Options)) (*ec2.ModifyIdentityIdFormatOutput, error)
	// Modifies the specified attribute of the specified AMI. You can specify only one
	// attribute at a time. You can use the Attribute parameter to specify the
	// attribute or one of the following parameters: Description or LaunchPermission.
	// Images with an Amazon Web Services Marketplace product code cannot be made
	// public. To enable the SriovNetSupport enhanced networking attribute of an image,
	// enable SriovNetSupport on an instance and create an AMI from the instance.
	//
	ModifyImageAttribute(arg1 context.Context, arg2 *ec2.ModifyImageAttributeInput, arg3 ...func(*ec2.Options)) (*ec2.ModifyImageAttributeOutput, error)
	// Modifies the specified attribute of the specified instance. You can specify only
	// one attribute at a time. Note: Using this action to change the security groups
	// associated with an elastic network interface (ENI) attached to an instance in a
	// VPC can result in an error if the instance has more than one ENI. To change the
	// security groups associated with an ENI attached to an instance that has multiple
	// ENIs, we recommend that you use the ModifyNetworkInterfaceAttribute action. To
	// modify some attributes, the instance must be stopped. For more information, see
	// Modify a stopped instance
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/Using_ChangingAttributesWhileInstanceStopped.html)
	// in the Amazon EC2 User Guide.
	//
	ModifyInstanceAttribute(arg1 context.Context, arg2 *ec2.ModifyInstanceAttributeInput, arg3 ...func(*ec2.Options)) (*ec2.ModifyInstanceAttributeOutput, error)
	// Modifies the Capacity Reservation settings for a stopped instance. Use this
	// action to configure an instance to target a specific Capacity Reservation, run
	// in any open Capacity Reservation with matching attributes, or run On-Demand
	// Instance capacity.
	//
	ModifyInstanceCapacityReservationAttributes(arg1 context.Context, arg2 *ec2.ModifyInstanceCapacityReservationAttributesInput, arg3 ...func(*ec2.Options)) (*ec2.ModifyInstanceCapacityReservationAttributesOutput, error)
	// Modifies the credit option for CPU usage on a running or stopped burstable
	// performance instance. The credit options are standard and unlimited. For more
	// information, see Burstable performance instances
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/burstable-performance-instances.html)
	// in the Amazon EC2 User Guide.
	//
	ModifyInstanceCreditSpecification(arg1 context.Context, arg2 *ec2.ModifyInstanceCreditSpecificationInput, arg3 ...func(*ec2.Options)) (*ec2.ModifyInstanceCreditSpecificationOutput, error)
	// Modifies the start time for a scheduled Amazon EC2 instance event.
	//
	ModifyInstanceEventStartTime(arg1 context.Context, arg2 *ec2.ModifyInstanceEventStartTimeInput, arg3 ...func(*ec2.Options)) (*ec2.ModifyInstanceEventStartTimeOutput, error)
	// Modifies the specified event window. You can define either a set of time ranges
	// or a cron expression when modifying the event window, but not both. To modify
	// the targets associated with the event window, use the
	// AssociateInstanceEventWindow and DisassociateInstanceEventWindow API. If Amazon
	// Web Services has already scheduled an event, modifying an event window won't
	// change the time of the scheduled event. For more information, see Define event
	// windows for scheduled events
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/event-windows.html) in the
	// Amazon EC2 User Guide.
	//
	ModifyInstanceEventWindow(arg1 context.Context, arg2 *ec2.ModifyInstanceEventWindowInput, arg3 ...func(*ec2.Options)) (*ec2.ModifyInstanceEventWindowOutput, error)
	// Modifies the recovery behavior of your instance to disable simplified automatic
	// recovery or set the recovery behavior to default. The default configuration will
	// not enable simplified automatic recovery for an unsupported instance type. For
	// more information, see Simplified automatic recovery
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-instance-recover.html#instance-configuration-recovery).
	//
	ModifyInstanceMaintenanceOptions(arg1 context.Context, arg2 *ec2.ModifyInstanceMaintenanceOptionsInput, arg3 ...func(*ec2.Options)) (*ec2.ModifyInstanceMaintenanceOptionsOutput, error)
	// Modify the instance metadata parameters on a running or stopped instance. When
	// you modify the parameters on a stopped instance, they are applied when the
	// instance is started. When you modify the parameters on a running instance, the
	// API responds with a state of “pending”. After the parameter modifications are
	// successfully applied to the instance, the state of the modifications changes
	// from “pending” to “applied” in subsequent describe-instances API calls. For more
	// information, see Instance metadata and user data
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-instance-metadata.html)
	// in the Amazon EC2 User Guide.
	//
	ModifyInstanceMetadataOptions(arg1 context.Context, arg2 *ec2.ModifyInstanceMetadataOptionsInput, arg3 ...func(*ec2.Options)) (*ec2.ModifyInstanceMetadataOptionsOutput, error)
	// Modifies the placement attributes for a specified instance. You can do the
	// following:
	//
	// * Modify the affinity between an instance and a Dedicated Host
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/dedicated-hosts-overview.html).
	// When affinity is set to host and the instance is not associated with a specific
	// Dedicated Host, the next time the instance is launched, it is automatically
	// associated with the host on which it lands. If the instance is restarted or
	// rebooted, this relationship persists.
	//
	// * Change the Dedicated Host with which an
	// instance is associated.
	//
	// * Change the instance tenancy of an instance.
	//
	// * Move
	// an instance to or from a placement group
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/placement-groups.html).
	//
	// At
	// least one attribute for affinity, host ID, tenancy, or placement group name must
	// be specified in the request. Affinity and tenancy can be modified in the same
	// request. To modify the host ID, tenancy, placement group, or partition for an
	// instance, the instance must be in the stopped state.
	//
	ModifyInstancePlacement(arg1 context.Context, arg2 *ec2.ModifyInstancePlacementInput, arg3 ...func(*ec2.Options)) (*ec2.ModifyInstancePlacementOutput, error)
	// Modify the configurations of an IPAM.
	//
	ModifyIpam(arg1 context.Context, arg2 *ec2.ModifyIpamInput, arg3 ...func(*ec2.Options)) (*ec2.ModifyIpamOutput, error)
	// Modify the configurations of an IPAM pool. For more information, see Modify a
	// pool (https://docs.aws.amazon.com/vpc/latest/ipam/mod-pool-ipam.html) in the
	// Amazon VPC IPAM User Guide.
	//
	ModifyIpamPool(arg1 context.Context, arg2 *ec2.ModifyIpamPoolInput, arg3 ...func(*ec2.Options)) (*ec2.ModifyIpamPoolOutput, error)
	// Modify a resource CIDR. You can use this action to transfer resource CIDRs
	// between scopes and ignore resource CIDRs that you do not want to manage. If set
	// to false, the resource will not be tracked for overlap, it cannot be
	// auto-imported into a pool, and it will be removed from any pool it has an
	// allocation in. For more information, see Move resource CIDRs between scopes
	// (https://docs.aws.amazon.com/vpc/latest/ipam/move-resource-ipam.html) and Change
	// the monitoring state of resource CIDRs
	// (https://docs.aws.amazon.com/vpc/latest/ipam/change-monitoring-state-ipam.html)
	// in the Amazon VPC IPAM User Guide.
	//
	ModifyIpamResourceCidr(arg1 context.Context, arg2 *ec2.ModifyIpamResourceCidrInput, arg3 ...func(*ec2.Options)) (*ec2.ModifyIpamResourceCidrOutput, error)
	// Modify an IPAM scope.
	//
	ModifyIpamScope(arg1 context.Context, arg2 *ec2.ModifyIpamScopeInput, arg3 ...func(*ec2.Options)) (*ec2.ModifyIpamScopeOutput, error)
	// Modifies a launch template. You can specify which version of the launch template
	// to set as the default version. When launching an instance, the default version
	// applies when a launch template version is not specified.
	//
	ModifyLaunchTemplate(arg1 context.Context, arg2 *ec2.ModifyLaunchTemplateInput, arg3 ...func(*ec2.Options)) (*ec2.ModifyLaunchTemplateOutput, error)
	// Modifies the specified managed prefix list. Adding or removing entries in a
	// prefix list creates a new version of the prefix list. Changing the name of the
	// prefix list does not affect the version. If you specify a current version number
	// that does not match the true current version number, the request fails.
	//
	ModifyManagedPrefixList(arg1 context.Context, arg2 *ec2.ModifyManagedPrefixListInput, arg3 ...func(*ec2.Options)) (*ec2.ModifyManagedPrefixListOutput, error)
	// Modifies the specified network interface attribute. You can specify only one
	// attribute at a time. You can use this action to attach and detach security
	// groups from an existing EC2 instance.
	//
	ModifyNetworkInterfaceAttribute(arg1 context.Context, arg2 *ec2.ModifyNetworkInterfaceAttributeInput, arg3 ...func(*ec2.Options)) (*ec2.ModifyNetworkInterfaceAttributeOutput, error)
	// Modifies the options for instance hostnames for the specified instance.
	//
	ModifyPrivateDnsNameOptions(arg1 context.Context, arg2 *ec2.ModifyPrivateDnsNameOptionsInput, arg3 ...func(*ec2.Options)) (*ec2.ModifyPrivateDnsNameOptionsOutput, error)
	// Modifies the Availability Zone, instance count, instance type, or network
	// platform (EC2-Classic or EC2-VPC) of your Reserved Instances. The Reserved
	// Instances to be modified must be identical, except for Availability Zone,
	// network platform, and instance type. For more information, see Modifying
	// Reserved Instances
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ri-modifying.html) in the
	// Amazon EC2 User Guide.
	//
	ModifyReservedInstances(arg1 context.Context, arg2 *ec2.ModifyReservedInstancesInput, arg3 ...func(*ec2.Options)) (*ec2.ModifyReservedInstancesOutput, error)
	// Modifies the rules of a security group.
	//
	ModifySecurityGroupRules(arg1 context.Context, arg2 *ec2.ModifySecurityGroupRulesInput, arg3 ...func(*ec2.Options)) (*ec2.ModifySecurityGroupRulesOutput, error)
	// Adds or removes permission settings for the specified snapshot. You may add or
	// remove specified Amazon Web Services account IDs from a snapshot's list of
	// create volume permissions, but you cannot do both in a single operation. If you
	// need to both add and remove account IDs for a snapshot, you must use multiple
	// operations. You can make up to 500 modifications to a snapshot in a single
	// operation. Encrypted snapshots and snapshots with Amazon Web Services
	// Marketplace product codes cannot be made public. Snapshots encrypted with your
	// default KMS key cannot be shared with other accounts. For more information about
	// modifying snapshot permissions, see Share a snapshot
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ebs-modifying-snapshot-permissions.html)
	// in the Amazon Elastic Compute Cloud User Guide.
	//
	ModifySnapshotAttribute(arg1 context.Context, arg2 *ec2.ModifySnapshotAttributeInput, arg3 ...func(*ec2.Options)) (*ec2.ModifySnapshotAttributeOutput, error)
	// Archives an Amazon EBS snapshot. When you archive a snapshot, it is converted to
	// a full snapshot that includes all of the blocks of data that were written to the
	// volume at the time the snapshot was created, and moved from the standard tier to
	// the archive tier. For more information, see Archive Amazon EBS snapshots
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/snapshot-archive.html) in
	// the Amazon Elastic Compute Cloud User Guide.
	//
	ModifySnapshotTier(arg1 context.Context, arg2 *ec2.ModifySnapshotTierInput, arg3 ...func(*ec2.Options)) (*ec2.ModifySnapshotTierOutput, error)
	// Modifies the specified Spot Fleet request. You can only modify a Spot Fleet
	// request of type maintain. While the Spot Fleet request is being modified, it is
	// in the modifying state. To scale up your Spot Fleet, increase its target
	// capacity. The Spot Fleet launches the additional Spot Instances according to the
	// allocation strategy for the Spot Fleet request. If the allocation strategy is
	// lowestPrice, the Spot Fleet launches instances using the Spot Instance pool with
	// the lowest price. If the allocation strategy is diversified, the Spot Fleet
	// distributes the instances across the Spot Instance pools. If the allocation
	// strategy is capacityOptimized, Spot Fleet launches instances from Spot Instance
	// pools with optimal capacity for the number of instances that are launching. To
	// scale down your Spot Fleet, decrease its target capacity. First, the Spot Fleet
	// cancels any open requests that exceed the new target capacity. You can request
	// that the Spot Fleet terminate Spot Instances until the size of the fleet no
	// longer exceeds the new target capacity. If the allocation strategy is
	// lowestPrice, the Spot Fleet terminates the instances with the highest price per
	// unit. If the allocation strategy is capacityOptimized, the Spot Fleet terminates
	// the instances in the Spot Instance pools that have the least available Spot
	// Instance capacity. If the allocation strategy is diversified, the Spot Fleet
	// terminates instances across the Spot Instance pools. Alternatively, you can
	// request that the Spot Fleet keep the fleet at its current size, but not replace
	// any Spot Instances that are interrupted or that you terminate manually. If you
	// are finished with your Spot Fleet for now, but will use it again later, you can
	// set the target capacity to 0.
	//
	ModifySpotFleetRequest(arg1 context.Context, arg2 *ec2.ModifySpotFleetRequestInput, arg3 ...func(*ec2.Options)) (*ec2.ModifySpotFleetRequestOutput, error)
	// Modifies a subnet attribute. You can only modify one attribute at a time. Use
	// this action to modify subnets on Amazon Web Services Outposts.
	//
	// * To modify a
	// subnet on an Outpost rack, set both MapCustomerOwnedIpOnLaunch and
	// CustomerOwnedIpv4Pool. These two parameters act as a single attribute.
	//
	// * To
	// modify a subnet on an Outpost server, set either EnableLniAtDeviceIndex or
	// DisableLniAtDeviceIndex.
	//
	// For more information about Amazon Web Services
	// Outposts, see the following:
	//
	// * Outpost servers
	// (https://docs.aws.amazon.com/outposts/latest/userguide/how-servers-work.html)
	//
	// *
	// Outpost racks
	// (https://docs.aws.amazon.com/outposts/latest/userguide/how-racks-work.html)
	//
	ModifySubnetAttribute(arg1 context.Context, arg2 *ec2.ModifySubnetAttributeInput, arg3 ...func(*ec2.Options)) (*ec2.ModifySubnetAttributeOutput, error)
	// Allows or restricts mirroring network services. By default, Amazon DNS network
	// services are not eligible for Traffic Mirror. Use AddNetworkServices to add
	// network services to a Traffic Mirror filter. When a network service is added to
	// the Traffic Mirror filter, all traffic related to that network service will be
	// mirrored. When you no longer want to mirror network services, use
	// RemoveNetworkServices to remove the network services from the Traffic Mirror
	// filter.
	//
	ModifyTrafficMirrorFilterNetworkServices(arg1 context.Context, arg2 *ec2.ModifyTrafficMirrorFilterNetworkServicesInput, arg3 ...func(*ec2.Options)) (*ec2.ModifyTrafficMirrorFilterNetworkServicesOutput, error)
	// Modifies the specified Traffic Mirror rule. DestinationCidrBlock and
	// SourceCidrBlock must both be an IPv4 range or an IPv6 range.
	//
	ModifyTrafficMirrorFilterRule(arg1 context.Context, arg2 *ec2.ModifyTrafficMirrorFilterRuleInput, arg3 ...func(*ec2.Options)) (*ec2.ModifyTrafficMirrorFilterRuleOutput, error)
	// Modifies a Traffic Mirror session.
	//
	ModifyTrafficMirrorSession(arg1 context.Context, arg2 *ec2.ModifyTrafficMirrorSessionInput, arg3 ...func(*ec2.Options)) (*ec2.ModifyTrafficMirrorSessionOutput, error)
	// Modifies the specified transit gateway. When you modify a transit gateway, the
	// modified options are applied to new transit gateway attachments only. Your
	// existing transit gateway attachments are not modified.
	//
	ModifyTransitGateway(arg1 context.Context, arg2 *ec2.ModifyTransitGatewayInput, arg3 ...func(*ec2.Options)) (*ec2.ModifyTransitGatewayOutput, error)
	// Modifies a reference (route) to a prefix list in a specified transit gateway
	// route table.
	//
	ModifyTransitGatewayPrefixListReference(arg1 context.Context, arg2 *ec2.ModifyTransitGatewayPrefixListReferenceInput, arg3 ...func(*ec2.Options)) (*ec2.ModifyTransitGatewayPrefixListReferenceOutput, error)
	// Modifies the specified VPC attachment.
	//
	ModifyTransitGatewayVpcAttachment(arg1 context.Context, arg2 *ec2.ModifyTransitGatewayVpcAttachmentInput, arg3 ...func(*ec2.Options)) (*ec2.ModifyTransitGatewayVpcAttachmentOutput, error)
	// You can modify several parameters of an existing EBS volume, including volume
	// size, volume type, and IOPS capacity. If your EBS volume is attached to a
	// current-generation EC2 instance type, you might be able to apply these changes
	// without stopping the instance or detaching the volume from it. For more
	// information about modifying EBS volumes, see Amazon EBS Elastic Volumes
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ebs-modify-volume.html)
	// (Linux instances) or Amazon EBS Elastic Volumes
	// (https://docs.aws.amazon.com/AWSEC2/latest/WindowsGuide/ebs-modify-volume.html)
	// (Windows instances). When you complete a resize operation on your volume, you
	// need to extend the volume's file-system size to take advantage of the new
	// storage capacity. For more information, see Extend a Linux file system
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ebs-expand-volume.html#recognize-expanded-volume-linux)
	// or Extend a Windows file system
	// (https://docs.aws.amazon.com/AWSEC2/latest/WindowsGuide/ebs-expand-volume.html#recognize-expanded-volume-windows).
	// You can use CloudWatch Events to check the status of a modification to an EBS
	// volume. For information about CloudWatch Events, see the Amazon CloudWatch
	// Events User Guide (https://docs.aws.amazon.com/AmazonCloudWatch/latest/events/).
	// You can also track the status of a modification using
	// DescribeVolumesModifications. For information about tracking status changes
	// using either method, see Monitor the progress of volume modifications
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/monitoring-volume-modifications.html).
	// With previous-generation instance types, resizing an EBS volume might require
	// detaching and reattaching the volume or stopping and restarting the instance.
	// After modifying a volume, you must wait at least six hours and ensure that the
	// volume is in the in-use or available state before you can modify the same
	// volume. This is sometimes referred to as a cooldown period.
	//
	ModifyVolume(arg1 context.Context, arg2 *ec2.ModifyVolumeInput, arg3 ...func(*ec2.Options)) (*ec2.ModifyVolumeOutput, error)
	// Modifies a volume attribute. By default, all I/O operations for the volume are
	// suspended when the data on the volume is determined to be potentially
	// inconsistent, to prevent undetectable, latent data corruption. The I/O access to
	// the volume can be resumed by first enabling I/O access and then checking the
	// data consistency on your volume. You can change the default behavior to resume
	// I/O operations. We recommend that you change this only for boot volumes or for
	// volumes that are stateless or disposable.
	//
	ModifyVolumeAttribute(arg1 context.Context, arg2 *ec2.ModifyVolumeAttributeInput, arg3 ...func(*ec2.Options)) (*ec2.ModifyVolumeAttributeOutput, error)
	// Modifies the specified attribute of the specified VPC.
	//
	ModifyVpcAttribute(arg1 context.Context, arg2 *ec2.ModifyVpcAttributeInput, arg3 ...func(*ec2.Options)) (*ec2.ModifyVpcAttributeOutput, error)
	// Modifies attributes of a specified VPC endpoint. The attributes that you can
	// modify depend on the type of VPC endpoint (interface, gateway, or Gateway Load
	// Balancer). For more information, see the Amazon Web Services PrivateLink Guide
	// (https://docs.aws.amazon.com/vpc/latest/privatelink/).
	//
	ModifyVpcEndpoint(arg1 context.Context, arg2 *ec2.ModifyVpcEndpointInput, arg3 ...func(*ec2.Options)) (*ec2.ModifyVpcEndpointOutput, error)
	// Modifies a connection notification for VPC endpoint or VPC endpoint service. You
	// can change the SNS topic for the notification, or the events for which to be
	// notified.
	//
	ModifyVpcEndpointConnectionNotification(arg1 context.Context, arg2 *ec2.ModifyVpcEndpointConnectionNotificationInput, arg3 ...func(*ec2.Options)) (*ec2.ModifyVpcEndpointConnectionNotificationOutput, error)
	// Modifies the attributes of your VPC endpoint service configuration. You can
	// change the Network Load Balancers or Gateway Load Balancers for your service,
	// and you can specify whether acceptance is required for requests to connect to
	// your endpoint service through an interface VPC endpoint. If you set or modify
	// the private DNS name, you must prove that you own the private DNS domain name.
	//
	ModifyVpcEndpointServiceConfiguration(arg1 context.Context, arg2 *ec2.ModifyVpcEndpointServiceConfigurationInput, arg3 ...func(*ec2.Options)) (*ec2.ModifyVpcEndpointServiceConfigurationOutput, error)
	// Modifies the payer responsibility for your VPC endpoint service.
	//
	ModifyVpcEndpointServicePayerResponsibility(arg1 context.Context, arg2 *ec2.ModifyVpcEndpointServicePayerResponsibilityInput, arg3 ...func(*ec2.Options)) (*ec2.ModifyVpcEndpointServicePayerResponsibilityOutput, error)
	// Modifies the permissions for your VPC endpoint service. You can add or remove
	// permissions for service consumers (IAM users, IAM roles, and Amazon Web Services
	// accounts) to connect to your endpoint service. If you grant permissions to all
	// principals, the service is public. Any users who know the name of a public
	// service can send a request to attach an endpoint. If the service does not
	// require manual approval, attachments are automatically approved.
	//
	ModifyVpcEndpointServicePermissions(arg1 context.Context, arg2 *ec2.ModifyVpcEndpointServicePermissionsInput, arg3 ...func(*ec2.Options)) (*ec2.ModifyVpcEndpointServicePermissionsOutput, error)
	// Modifies the VPC peering connection options on one side of a VPC peering
	// connection. You can do the following:
	//
	// * Enable/disable communication over the
	// peering connection between an EC2-Classic instance that's linked to your VPC
	// (using ClassicLink) and instances in the peer VPC.
	//
	// * Enable/disable
	// communication over the peering connection between instances in your VPC and an
	// EC2-Classic instance that's linked to the peer VPC.
	//
	// * Enable/disable the
	// ability to resolve public DNS hostnames to private IP addresses when queried
	// from instances in the peer VPC.
	//
	// If the peered VPCs are in the same Amazon Web
	// Services account, you can enable DNS resolution for queries from the local VPC.
	// This ensures that queries from the local VPC resolve to private IP addresses in
	// the peer VPC. This option is not available if the peered VPCs are in different
	// different Amazon Web Services accounts or different Regions. For peered VPCs in
	// different Amazon Web Services accounts, each Amazon Web Services account owner
	// must initiate a separate request to modify the peering connection options. For
	// inter-region peering connections, you must use the Region for the requester VPC
	// to modify the requester VPC peering options and the Region for the accepter VPC
	// to modify the accepter VPC peering options. To verify which VPCs are the
	// accepter and the requester for a VPC peering connection, use the
	// DescribeVpcPeeringConnections command.
	//
	ModifyVpcPeeringConnectionOptions(arg1 context.Context, arg2 *ec2.ModifyVpcPeeringConnectionOptionsInput, arg3 ...func(*ec2.Options)) (*ec2.ModifyVpcPeeringConnectionOptionsOutput, error)
	// Modifies the instance tenancy attribute of the specified VPC. You can change the
	// instance tenancy attribute of a VPC to default only. You cannot change the
	// instance tenancy attribute to dedicated. After you modify the tenancy of the
	// VPC, any new instances that you launch into the VPC have a tenancy of default,
	// unless you specify otherwise during launch. The tenancy of any existing
	// instances in the VPC is not affected. For more information, see Dedicated
	// Instances
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/dedicated-instance.html) in
	// the Amazon Elastic Compute Cloud User Guide.
	//
	ModifyVpcTenancy(arg1 context.Context, arg2 *ec2.ModifyVpcTenancyInput, arg3 ...func(*ec2.Options)) (*ec2.ModifyVpcTenancyOutput, error)
	// Modifies the customer gateway or the target gateway of an Amazon Web Services
	// Site-to-Site VPN connection. To modify the target gateway, the following
	// migration options are available:
	//
	// * An existing virtual private gateway to a new
	// virtual private gateway
	//
	// * An existing virtual private gateway to a transit
	// gateway
	//
	// * An existing transit gateway to a new transit gateway
	//
	// * An existing
	// transit gateway to a virtual private gateway
	//
	// Before you perform the migration
	// to the new gateway, you must configure the new gateway. Use CreateVpnGateway to
	// create a virtual private gateway, or CreateTransitGateway to create a transit
	// gateway. This step is required when you migrate from a virtual private gateway
	// with static routes to a transit gateway. You must delete the static routes
	// before you migrate to the new gateway. Keep a copy of the static route before
	// you delete it. You will need to add back these routes to the transit gateway
	// after the VPN connection migration is complete. After you migrate to the new
	// gateway, you might need to modify your VPC route table. Use CreateRoute and
	// DeleteRoute to make the changes described in Update VPC route tables
	// (https://docs.aws.amazon.com/vpn/latest/s2svpn/modify-vpn-target.html#step-update-routing)
	// in the Amazon Web Services Site-to-Site VPN User Guide. When the new gateway is
	// a transit gateway, modify the transit gateway route table to allow traffic
	// between the VPC and the Amazon Web Services Site-to-Site VPN connection. Use
	// CreateTransitGatewayRoute to add the routes. If you deleted VPN static routes,
	// you must add the static routes to the transit gateway route table. After you
	// perform this operation, the VPN endpoint's IP addresses on the Amazon Web
	// Services side and the tunnel options remain intact. Your Amazon Web Services
	// Site-to-Site VPN connection will be temporarily unavailable for a brief period
	// while we provision the new endpoints.
	//
	ModifyVpnConnection(arg1 context.Context, arg2 *ec2.ModifyVpnConnectionInput, arg3 ...func(*ec2.Options)) (*ec2.ModifyVpnConnectionOutput, error)
	// Modifies the connection options for your Site-to-Site VPN connection. When you
	// modify the VPN connection options, the VPN endpoint IP addresses on the Amazon
	// Web Services side do not change, and the tunnel options do not change. Your VPN
	// connection will be temporarily unavailable for a brief period while the VPN
	// connection is updated.
	//
	ModifyVpnConnectionOptions(arg1 context.Context, arg2 *ec2.ModifyVpnConnectionOptionsInput, arg3 ...func(*ec2.Options)) (*ec2.ModifyVpnConnectionOptionsOutput, error)
	// Modifies the VPN tunnel endpoint certificate.
	//
	ModifyVpnTunnelCertificate(arg1 context.Context, arg2 *ec2.ModifyVpnTunnelCertificateInput, arg3 ...func(*ec2.Options)) (*ec2.ModifyVpnTunnelCertificateOutput, error)
	// Modifies the options for a VPN tunnel in an Amazon Web Services Site-to-Site VPN
	// connection. You can modify multiple options for a tunnel in a single request,
	// but you can only modify one tunnel at a time. For more information, see
	// Site-to-Site VPN tunnel options for your Site-to-Site VPN connection
	// (https://docs.aws.amazon.com/vpn/latest/s2svpn/VPNTunnels.html) in the Amazon
	// Web Services Site-to-Site VPN User Guide.
	//
	ModifyVpnTunnelOptions(arg1 context.Context, arg2 *ec2.ModifyVpnTunnelOptionsInput, arg3 ...func(*ec2.Options)) (*ec2.ModifyVpnTunnelOptionsOutput, error)
	// Enables detailed monitoring for a running instance. Otherwise, basic monitoring
	// is enabled. For more information, see Monitor your instances using CloudWatch
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/using-cloudwatch.html) in
	// the Amazon EC2 User Guide. To disable detailed monitoring, see
	// UnmonitorInstances
	// (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_UnmonitorInstances.html).
	//
	MonitorInstances(arg1 context.Context, arg2 *ec2.MonitorInstancesInput, arg3 ...func(*ec2.Options)) (*ec2.MonitorInstancesOutput, error)
	// Moves an Elastic IP address from the EC2-Classic platform to the EC2-VPC
	// platform. The Elastic IP address must be allocated to your account for more than
	// 24 hours, and it must not be associated with an instance. After the Elastic IP
	// address is moved, it is no longer available for use in the EC2-Classic platform,
	// unless you move it back using the RestoreAddressToClassic request. You cannot
	// move an Elastic IP address that was originally allocated for use in the EC2-VPC
	// platform to the EC2-Classic platform.
	//
	MoveAddressToVpc(arg1 context.Context, arg2 *ec2.MoveAddressToVpcInput, arg3 ...func(*ec2.Options)) (*ec2.MoveAddressToVpcOutput, error)
	// Move an BYOIP IPv4 CIDR to IPAM from a public IPv4 pool. If you already have an
	// IPv4 BYOIP CIDR with Amazon Web Services, you can move the CIDR to IPAM from a
	// public IPv4 pool. You cannot move an IPv6 CIDR to IPAM. If you are bringing a
	// new IP address to Amazon Web Services for the first time, complete the steps in
	// Tutorial: BYOIP address CIDRs to IPAM
	// (https://docs.aws.amazon.com/vpc/latest/ipam/tutorials-byoip-ipam.html).
	//
	MoveByoipCidrToIpam(arg1 context.Context, arg2 *ec2.MoveByoipCidrToIpamInput, arg3 ...func(*ec2.Options)) (*ec2.MoveByoipCidrToIpamOutput, error)
	// Provisions an IPv4 or IPv6 address range for use with your Amazon Web Services
	// resources through bring your own IP addresses (BYOIP) and creates a
	// corresponding address pool. After the address range is provisioned, it is ready
	// to be advertised using AdvertiseByoipCidr. Amazon Web Services verifies that you
	// own the address range and are authorized to advertise it. You must ensure that
	// the address range is registered to you and that you created an RPKI ROA to
	// authorize Amazon ASNs 16509 and 14618 to advertise the address range. For more
	// information, see Bring your own IP addresses (BYOIP)
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-byoip.html) in the
	// Amazon Elastic Compute Cloud User Guide. Provisioning an address range is an
	// asynchronous operation, so the call returns immediately, but the address range
	// is not ready to use until its status changes from pending-provision to
	// provisioned. To monitor the status of an address range, use DescribeByoipCidrs.
	// To allocate an Elastic IP address from your IPv4 address pool, use
	// AllocateAddress with either the specific address from the address pool or the ID
	// of the address pool.
	//
	ProvisionByoipCidr(arg1 context.Context, arg2 *ec2.ProvisionByoipCidrInput, arg3 ...func(*ec2.Options)) (*ec2.ProvisionByoipCidrOutput, error)
	// Provision a CIDR to an IPAM pool. You can use this action to provision new CIDRs
	// to a top-level pool or to transfer a CIDR from a top-level pool to a pool within
	// it. For more information, see Provision CIDRs to pools
	// (https://docs.aws.amazon.com/vpc/latest/ipam/prov-cidr-ipam.html) in the Amazon
	// VPC IPAM User Guide.
	//
	ProvisionIpamPoolCidr(arg1 context.Context, arg2 *ec2.ProvisionIpamPoolCidrInput, arg3 ...func(*ec2.Options)) (*ec2.ProvisionIpamPoolCidrOutput, error)
	// Provision a CIDR to a public IPv4 pool. For more information about IPAM, see
	// What is IPAM? (https://docs.aws.amazon.com/vpc/latest/ipam/what-is-it-ipam.html)
	// in the Amazon VPC IPAM User Guide.
	//
	ProvisionPublicIpv4PoolCidr(arg1 context.Context, arg2 *ec2.ProvisionPublicIpv4PoolCidrInput, arg3 ...func(*ec2.Options)) (*ec2.ProvisionPublicIpv4PoolCidrOutput, error)
	// Purchase a reservation with configurations that match those of your Dedicated
	// Host. You must have active Dedicated Hosts in your account before you purchase a
	// reservation. This action results in the specified reservation being purchased
	// and charged to your account.
	//
	PurchaseHostReservation(arg1 context.Context, arg2 *ec2.PurchaseHostReservationInput, arg3 ...func(*ec2.Options)) (*ec2.PurchaseHostReservationOutput, error)
	// Purchases a Reserved Instance for use with your account. With Reserved
	// Instances, you pay a lower hourly rate compared to On-Demand instance pricing.
	// Use DescribeReservedInstancesOfferings to get a list of Reserved Instance
	// offerings that match your specifications. After you've purchased a Reserved
	// Instance, you can check for your new Reserved Instance with
	// DescribeReservedInstances. To queue a purchase for a future date and time,
	// specify a purchase time. If you do not specify a purchase time, the default is
	// the current time. For more information, see Reserved Instances
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/concepts-on-demand-reserved-instances.html)
	// and Reserved Instance Marketplace
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ri-market-general.html) in
	// the Amazon EC2 User Guide.
	//
	PurchaseReservedInstancesOffering(arg1 context.Context, arg2 *ec2.PurchaseReservedInstancesOfferingInput, arg3 ...func(*ec2.Options)) (*ec2.PurchaseReservedInstancesOfferingOutput, error)
	// Purchases the Scheduled Instances with the specified schedule. Scheduled
	// Instances enable you to purchase Amazon EC2 compute capacity by the hour for a
	// one-year term. Before you can purchase a Scheduled Instance, you must call
	// DescribeScheduledInstanceAvailability to check for available schedules and
	// obtain a purchase token. After you purchase a Scheduled Instance, you must call
	// RunScheduledInstances during each scheduled time period. After you purchase a
	// Scheduled Instance, you can't cancel, modify, or resell your purchase.
	//
	PurchaseScheduledInstances(arg1 context.Context, arg2 *ec2.PurchaseScheduledInstancesInput, arg3 ...func(*ec2.Options)) (*ec2.PurchaseScheduledInstancesOutput, error)
	// Requests a reboot of the specified instances. This operation is asynchronous; it
	// only queues a request to reboot the specified instances. The operation succeeds
	// if the instances are valid and belong to you. Requests to reboot terminated
	// instances are ignored. If an instance does not cleanly shut down within a few
	// minutes, Amazon EC2 performs a hard reboot. For more information about
	// troubleshooting, see Troubleshoot an unreachable instance
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instance-console.html) in
	// the Amazon EC2 User Guide.
	//
	RebootInstances(arg1 context.Context, arg2 *ec2.RebootInstancesInput, arg3 ...func(*ec2.Options)) (*ec2.RebootInstancesOutput, error)
	// Registers an AMI. When you're creating an AMI, this is the final step you must
	// complete before you can launch an instance from the AMI. For more information
	// about creating AMIs, see Creating your own AMIs
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/creating-an-ami.html) in
	// the Amazon Elastic Compute Cloud User Guide. For Amazon EBS-backed instances,
	// CreateImage creates and registers the AMI in a single request, so you don't have
	// to register the AMI yourself. We recommend that you always use CreateImage
	// unless you have a specific reason to use RegisterImage. If needed, you can
	// deregister an AMI at any time. Any modifications you make to an AMI backed by an
	// instance store volume invalidates its registration. If you make changes to an
	// image, deregister the previous image and register the new image. Register a
	// snapshot of a root device volume You can use RegisterImage to create an Amazon
	// EBS-backed Linux AMI from a snapshot of a root device volume. You specify the
	// snapshot using a block device mapping. You can't set the encryption state of the
	// volume using the block device mapping. If the snapshot is encrypted, or
	// encryption by default is enabled, the root volume of an instance launched from
	// the AMI is encrypted. For more information, see Create a Linux AMI from a
	// snapshot
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/creating-an-ami-ebs.html#creating-launching-ami-from-snapshot)
	// and Use encryption with Amazon EBS-backed AMIs
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/AMIEncryption.html) in the
	// Amazon Elastic Compute Cloud User Guide. Amazon Web Services Marketplace product
	// codes If any snapshots have Amazon Web Services Marketplace product codes, they
	// are copied to the new AMI. Windows and some Linux distributions, such as Red Hat
	// Enterprise Linux (RHEL) and SUSE Linux Enterprise Server (SLES), use the Amazon
	// EC2 billing product code associated with an AMI to verify the subscription
	// status for package updates. To create a new AMI for operating systems that
	// require a billing product code, instead of registering the AMI, do the following
	// to preserve the billing product code association:
	//
	// * Launch an instance from an
	// existing AMI with that billing product code.
	//
	// * Customize the instance.
	//
	// *
	// Create an AMI from the instance using CreateImage.
	//
	// If you purchase a Reserved
	// Instance to apply to an On-Demand Instance that was launched from an AMI with a
	// billing product code, make sure that the Reserved Instance has the matching
	// billing product code. If you purchase a Reserved Instance without the matching
	// billing product code, the Reserved Instance will not be applied to the On-Demand
	// Instance. For information about how to obtain the platform details and billing
	// information of an AMI, see Understanding AMI billing
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ami-billing-info.html) in
	// the Amazon Elastic Compute Cloud User Guide.
	//
	RegisterImage(arg1 context.Context, arg2 *ec2.RegisterImageInput, arg3 ...func(*ec2.Options)) (*ec2.RegisterImageOutput, error)
	// Registers a set of tag keys to include in scheduled event notifications for your
	// resources. To remove tags, use DeregisterInstanceEventNotificationAttributes
	// (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DeregisterInstanceEventNotificationAttributes.html).
	//
	RegisterInstanceEventNotificationAttributes(arg1 context.Context, arg2 *ec2.RegisterInstanceEventNotificationAttributesInput, arg3 ...func(*ec2.Options)) (*ec2.RegisterInstanceEventNotificationAttributesOutput, error)
	// Registers members (network interfaces) with the transit gateway multicast group.
	// A member is a network interface associated with a supported EC2 instance that
	// receives multicast traffic. For information about supported instances, see
	// Multicast Consideration
	// (https://docs.aws.amazon.com/vpc/latest/tgw/transit-gateway-limits.html#multicast-limits)
	// in Amazon VPC Transit Gateways. After you add the members, use
	// SearchTransitGatewayMulticastGroups
	// (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_SearchTransitGatewayMulticastGroups.html)
	// to verify that the members were added to the transit gateway multicast group.
	//
	RegisterTransitGatewayMulticastGroupMembers(arg1 context.Context, arg2 *ec2.RegisterTransitGatewayMulticastGroupMembersInput, arg3 ...func(*ec2.Options)) (*ec2.RegisterTransitGatewayMulticastGroupMembersOutput, error)
	// Registers sources (network interfaces) with the specified transit gateway
	// multicast group. A multicast source is a network interface attached to a
	// supported instance that sends multicast traffic. For information about supported
	// instances, see Multicast Considerations
	// (https://docs.aws.amazon.com/vpc/latest/tgw/transit-gateway-limits.html#multicast-limits)
	// in Amazon VPC Transit Gateways. After you add the source, use
	// SearchTransitGatewayMulticastGroups
	// (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_SearchTransitGatewayMulticastGroups.html)
	// to verify that the source was added to the multicast group.
	//
	RegisterTransitGatewayMulticastGroupSources(arg1 context.Context, arg2 *ec2.RegisterTransitGatewayMulticastGroupSourcesInput, arg3 ...func(*ec2.Options)) (*ec2.RegisterTransitGatewayMulticastGroupSourcesOutput, error)
	// Rejects a request to associate cross-account subnets with a transit gateway
	// multicast domain.
	//
	RejectTransitGatewayMulticastDomainAssociations(arg1 context.Context, arg2 *ec2.RejectTransitGatewayMulticastDomainAssociationsInput, arg3 ...func(*ec2.Options)) (*ec2.RejectTransitGatewayMulticastDomainAssociationsOutput, error)
	// Rejects a transit gateway peering attachment request.
	//
	RejectTransitGatewayPeeringAttachment(arg1 context.Context, arg2 *ec2.RejectTransitGatewayPeeringAttachmentInput, arg3 ...func(*ec2.Options)) (*ec2.RejectTransitGatewayPeeringAttachmentOutput, error)
	// Rejects a request to attach a VPC to a transit gateway. The VPC attachment must
	// be in the pendingAcceptance state. Use DescribeTransitGatewayVpcAttachments to
	// view your pending VPC attachment requests. Use AcceptTransitGatewayVpcAttachment
	// to accept a VPC attachment request.
	//
	RejectTransitGatewayVpcAttachment(arg1 context.Context, arg2 *ec2.RejectTransitGatewayVpcAttachmentInput, arg3 ...func(*ec2.Options)) (*ec2.RejectTransitGatewayVpcAttachmentOutput, error)
	// Rejects one or more VPC endpoint connection requests to your VPC endpoint
	// service.
	//
	RejectVpcEndpointConnections(arg1 context.Context, arg2 *ec2.RejectVpcEndpointConnectionsInput, arg3 ...func(*ec2.Options)) (*ec2.RejectVpcEndpointConnectionsOutput, error)
	// Rejects a VPC peering connection request. The VPC peering connection must be in
	// the pending-acceptance state. Use the DescribeVpcPeeringConnections request to
	// view your outstanding VPC peering connection requests. To delete an active VPC
	// peering connection, or to delete a VPC peering connection request that you
	// initiated, use DeleteVpcPeeringConnection.
	//
	RejectVpcPeeringConnection(arg1 context.Context, arg2 *ec2.RejectVpcPeeringConnectionInput, arg3 ...func(*ec2.Options)) (*ec2.RejectVpcPeeringConnectionOutput, error)
	// Releases the specified Elastic IP address. [EC2-Classic, default VPC] Releasing
	// an Elastic IP address automatically disassociates it from any instance that it's
	// associated with. To disassociate an Elastic IP address without releasing it, use
	// DisassociateAddress. [Nondefault VPC] You must use DisassociateAddress to
	// disassociate the Elastic IP address before you can release it. Otherwise, Amazon
	// EC2 returns an error (InvalidIPAddress.InUse). After releasing an Elastic IP
	// address, it is released to the IP address pool. Be sure to update your DNS
	// records and any servers or devices that communicate with the address. If you
	// attempt to release an Elastic IP address that you already released, you'll get
	// an AuthFailure error if the address is already allocated to another Amazon Web
	// Services account. [EC2-VPC] After you release an Elastic IP address for use in a
	// VPC, you might be able to recover it. For more information, see AllocateAddress.
	// For more information, see Elastic IP Addresses
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/elastic-ip-addresses-eip.html)
	// in the Amazon Elastic Compute Cloud User Guide.
	//
	ReleaseAddress(arg1 context.Context, arg2 *ec2.ReleaseAddressInput, arg3 ...func(*ec2.Options)) (*ec2.ReleaseAddressOutput, error)
	// When you no longer want to use an On-Demand Dedicated Host it can be released.
	// On-Demand billing is stopped and the host goes into released state. The host ID
	// of Dedicated Hosts that have been released can no longer be specified in another
	// request, for example, to modify the host. You must stop or terminate all
	// instances on a host before it can be released. When Dedicated Hosts are
	// released, it may take some time for them to stop counting toward your limit and
	// you may receive capacity errors when trying to allocate new Dedicated Hosts.
	// Wait a few minutes and then try again. Released hosts still appear in a
	// DescribeHosts response.
	//
	ReleaseHosts(arg1 context.Context, arg2 *ec2.ReleaseHostsInput, arg3 ...func(*ec2.Options)) (*ec2.ReleaseHostsOutput, error)
	// Release an allocation within an IPAM pool. You can only use this action to
	// release manual allocations. To remove an allocation for a resource without
	// deleting the resource, set its monitored state to false using
	// ModifyIpamResourceCidr
	// (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_ModifyIpamResourceCidr.html).
	// For more information, see Release an allocation
	// (https://docs.aws.amazon.com/vpc/latest/ipam/release-pool-alloc-ipam.html) in
	// the Amazon VPC IPAM User Guide.
	//
	ReleaseIpamPoolAllocation(arg1 context.Context, arg2 *ec2.ReleaseIpamPoolAllocationInput, arg3 ...func(*ec2.Options)) (*ec2.ReleaseIpamPoolAllocationOutput, error)
	// Replaces an IAM instance profile for the specified running instance. You can use
	// this action to change the IAM instance profile that's associated with an
	// instance without having to disassociate the existing IAM instance profile first.
	// Use DescribeIamInstanceProfileAssociations to get the association ID.
	//
	ReplaceIamInstanceProfileAssociation(arg1 context.Context, arg2 *ec2.ReplaceIamInstanceProfileAssociationInput, arg3 ...func(*ec2.Options)) (*ec2.ReplaceIamInstanceProfileAssociationOutput, error)
	// Changes which network ACL a subnet is associated with. By default when you
	// create a subnet, it's automatically associated with the default network ACL. For
	// more information, see Network ACLs
	// (https://docs.aws.amazon.com/vpc/latest/userguide/VPC_ACLs.html) in the Amazon
	// Virtual Private Cloud User Guide. This is an idempotent operation.
	//
	ReplaceNetworkAclAssociation(arg1 context.Context, arg2 *ec2.ReplaceNetworkAclAssociationInput, arg3 ...func(*ec2.Options)) (*ec2.ReplaceNetworkAclAssociationOutput, error)
	// Replaces an entry (rule) in a network ACL. For more information, see Network
	// ACLs (https://docs.aws.amazon.com/vpc/latest/userguide/VPC_ACLs.html) in the
	// Amazon Virtual Private Cloud User Guide.
	//
	ReplaceNetworkAclEntry(arg1 context.Context, arg2 *ec2.ReplaceNetworkAclEntryInput, arg3 ...func(*ec2.Options)) (*ec2.ReplaceNetworkAclEntryOutput, error)
	// Replaces an existing route within a route table in a VPC. You must specify
	// either a destination CIDR block or a prefix list ID. You must also specify
	// exactly one of the resources from the parameter list, or reset the local route
	// to its default target. For more information, see Route tables
	// (https://docs.aws.amazon.com/vpc/latest/userguide/VPC_Route_Tables.html) in the
	// Amazon Virtual Private Cloud User Guide.
	//
	ReplaceRoute(arg1 context.Context, arg2 *ec2.ReplaceRouteInput, arg3 ...func(*ec2.Options)) (*ec2.ReplaceRouteOutput, error)
	// Changes the route table associated with a given subnet, internet gateway, or
	// virtual private gateway in a VPC. After the operation completes, the subnet or
	// gateway uses the routes in the new route table. For more information about route
	// tables, see Route tables
	// (https://docs.aws.amazon.com/vpc/latest/userguide/VPC_Route_Tables.html) in the
	// Amazon Virtual Private Cloud User Guide. You can also use this operation to
	// change which table is the main route table in the VPC. Specify the main route
	// table's association ID and the route table ID of the new main route table.
	//
	ReplaceRouteTableAssociation(arg1 context.Context, arg2 *ec2.ReplaceRouteTableAssociationInput, arg3 ...func(*ec2.Options)) (*ec2.ReplaceRouteTableAssociationOutput, error)
	// Replaces the specified route in the specified transit gateway route table.
	//
	ReplaceTransitGatewayRoute(arg1 context.Context, arg2 *ec2.ReplaceTransitGatewayRouteInput, arg3 ...func(*ec2.Options)) (*ec2.ReplaceTransitGatewayRouteOutput, error)
	// Submits feedback about the status of an instance. The instance must be in the
	// running state. If your experience with the instance differs from the instance
	// status returned by DescribeInstanceStatus, use ReportInstanceStatus to report
	// your experience with the instance. Amazon EC2 collects this information to
	// improve the accuracy of status checks. Use of this action does not change the
	// value returned by DescribeInstanceStatus.
	//
	ReportInstanceStatus(arg1 context.Context, arg2 *ec2.ReportInstanceStatusInput, arg3 ...func(*ec2.Options)) (*ec2.ReportInstanceStatusOutput, error)
	// Creates a Spot Fleet request. The Spot Fleet request specifies the total target
	// capacity and the On-Demand target capacity. Amazon EC2 calculates the difference
	// between the total capacity and On-Demand capacity, and launches the difference
	// as Spot capacity. You can submit a single request that includes multiple launch
	// specifications that vary by instance type, AMI, Availability Zone, or subnet. By
	// default, the Spot Fleet requests Spot Instances in the Spot Instance pool where
	// the price per unit is the lowest. Each launch specification can include its own
	// instance weighting that reflects the value of the instance type to your
	// application workload. Alternatively, you can specify that the Spot Fleet
	// distribute the target capacity across the Spot pools included in its launch
	// specifications. By ensuring that the Spot Instances in your Spot Fleet are in
	// different Spot pools, you can improve the availability of your fleet. You can
	// specify tags for the Spot Fleet request and instances launched by the fleet. You
	// cannot tag other resource types in a Spot Fleet request because only the
	// spot-fleet-request and instance resource types are supported. For more
	// information, see Spot Fleet requests
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/spot-fleet-requests.html)
	// in the Amazon EC2 User Guide for Linux Instances. We strongly discourage using
	// the RequestSpotFleet API because it is a legacy API with no planned investment.
	// For options for requesting Spot Instances, see Which is the best Spot request
	// method to use?
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/spot-best-practices.html#which-spot-request-method-to-use)
	// in the Amazon EC2 User Guide for Linux Instances.
	//
	RequestSpotFleet(arg1 context.Context, arg2 *ec2.RequestSpotFleetInput, arg3 ...func(*ec2.Options)) (*ec2.RequestSpotFleetOutput, error)
	// Creates a Spot Instance request. For more information, see Spot Instance
	// requests
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/spot-requests.html) in the
	// Amazon EC2 User Guide for Linux Instances. We strongly discourage using the
	// RequestSpotInstances API because it is a legacy API with no planned investment.
	// For options for requesting Spot Instances, see Which is the best Spot request
	// method to use?
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/spot-best-practices.html#which-spot-request-method-to-use)
	// in the Amazon EC2 User Guide for Linux Instances.
	//
	RequestSpotInstances(arg1 context.Context, arg2 *ec2.RequestSpotInstancesInput, arg3 ...func(*ec2.Options)) (*ec2.RequestSpotInstancesOutput, error)
	// Resets the attribute of the specified IP address. For requirements, see Using
	// reverse DNS for email applications
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/elastic-ip-addresses-eip.html#Using_Elastic_Addressing_Reverse_DNS).
	//
	ResetAddressAttribute(arg1 context.Context, arg2 *ec2.ResetAddressAttributeInput, arg3 ...func(*ec2.Options)) (*ec2.ResetAddressAttributeOutput, error)
	// Resets the default KMS key for EBS encryption for your account in this Region to
	// the Amazon Web Services managed KMS key for EBS. After resetting the default KMS
	// key to the Amazon Web Services managed KMS key, you can continue to encrypt by a
	// customer managed KMS key by specifying it when you create the volume. For more
	// information, see Amazon EBS encryption
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSEncryption.html) in the
	// Amazon Elastic Compute Cloud User Guide.
	//
	ResetEbsDefaultKmsKeyId(arg1 context.Context, arg2 *ec2.ResetEbsDefaultKmsKeyIdInput, arg3 ...func(*ec2.Options)) (*ec2.ResetEbsDefaultKmsKeyIdOutput, error)
	// Resets the specified attribute of the specified Amazon FPGA Image (AFI) to its
	// default value. You can only reset the load permission attribute.
	//
	ResetFpgaImageAttribute(arg1 context.Context, arg2 *ec2.ResetFpgaImageAttributeInput, arg3 ...func(*ec2.Options)) (*ec2.ResetFpgaImageAttributeOutput, error)
	// Resets an attribute of an AMI to its default value.
	//
	ResetImageAttribute(arg1 context.Context, arg2 *ec2.ResetImageAttributeInput, arg3 ...func(*ec2.Options)) (*ec2.ResetImageAttributeOutput, error)
	// Resets an attribute of an instance to its default value. To reset the kernel or
	// ramdisk, the instance must be in a stopped state. To reset the sourceDestCheck,
	// the instance can be either running or stopped. The sourceDestCheck attribute
	// controls whether source/destination checking is enabled. The default value is
	// true, which means checking is enabled. This value must be false for a NAT
	// instance to perform NAT. For more information, see NAT Instances
	// (https://docs.aws.amazon.com/AmazonVPC/latest/UserGuide/VPC_NAT_Instance.html)
	// in the Amazon VPC User Guide.
	//
	ResetInstanceAttribute(arg1 context.Context, arg2 *ec2.ResetInstanceAttributeInput, arg3 ...func(*ec2.Options)) (*ec2.ResetInstanceAttributeOutput, error)
	// Resets a network interface attribute. You can specify only one attribute at a
	// time.
	//
	ResetNetworkInterfaceAttribute(arg1 context.Context, arg2 *ec2.ResetNetworkInterfaceAttributeInput, arg3 ...func(*ec2.Options)) (*ec2.ResetNetworkInterfaceAttributeOutput, error)
	// Resets permission settings for the specified snapshot. For more information
	// about modifying snapshot permissions, see Share a snapshot
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ebs-modifying-snapshot-permissions.html)
	// in the Amazon Elastic Compute Cloud User Guide.
	//
	ResetSnapshotAttribute(arg1 context.Context, arg2 *ec2.ResetSnapshotAttributeInput, arg3 ...func(*ec2.Options)) (*ec2.ResetSnapshotAttributeOutput, error)
	// Restores an Elastic IP address that was previously moved to the EC2-VPC platform
	// back to the EC2-Classic platform. You cannot move an Elastic IP address that was
	// originally allocated for use in EC2-VPC. The Elastic IP address must not be
	// associated with an instance or network interface.
	//
	RestoreAddressToClassic(arg1 context.Context, arg2 *ec2.RestoreAddressToClassicInput, arg3 ...func(*ec2.Options)) (*ec2.RestoreAddressToClassicOutput, error)
	// Restores an AMI from the Recycle Bin. For more information, see Recycle Bin
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/recycle-bin.html) in the
	// Amazon Elastic Compute Cloud User Guide.
	//
	RestoreImageFromRecycleBin(arg1 context.Context, arg2 *ec2.RestoreImageFromRecycleBinInput, arg3 ...func(*ec2.Options)) (*ec2.RestoreImageFromRecycleBinOutput, error)
	// Restores the entries from a previous version of a managed prefix list to a new
	// version of the prefix list.
	//
	RestoreManagedPrefixListVersion(arg1 context.Context, arg2 *ec2.RestoreManagedPrefixListVersionInput, arg3 ...func(*ec2.Options)) (*ec2.RestoreManagedPrefixListVersionOutput, error)
	// Restores a snapshot from the Recycle Bin. For more information, see Restore
	// snapshots from the Recycle Bin
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/recycle-bin-working-with-snaps.html#recycle-bin-restore-snaps)
	// in the Amazon Elastic Compute Cloud User Guide.
	//
	RestoreSnapshotFromRecycleBin(arg1 context.Context, arg2 *ec2.RestoreSnapshotFromRecycleBinInput, arg3 ...func(*ec2.Options)) (*ec2.RestoreSnapshotFromRecycleBinOutput, error)
	// Restores an archived Amazon EBS snapshot for use temporarily or permanently, or
	// modifies the restore period or restore type for a snapshot that was previously
	// temporarily restored. For more information see  Restore an archived snapshot
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/working-with-snapshot-archiving.html#restore-archived-snapshot)
	// and  modify the restore period or restore type for a temporarily restored
	// snapshot
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/working-with-snapshot-archiving.html#modify-temp-restore-period)
	// in the Amazon Elastic Compute Cloud User Guide.
	//
	RestoreSnapshotTier(arg1 context.Context, arg2 *ec2.RestoreSnapshotTierInput, arg3 ...func(*ec2.Options)) (*ec2.RestoreSnapshotTierOutput, error)
	// Removes an ingress authorization rule from a Client VPN endpoint.
	//
	RevokeClientVpnIngress(arg1 context.Context, arg2 *ec2.RevokeClientVpnIngressInput, arg3 ...func(*ec2.Options)) (*ec2.RevokeClientVpnIngressOutput, error)
	// [VPC only] Removes the specified outbound (egress) rules from a security group
	// for EC2-VPC. This action does not apply to security groups for use in
	// EC2-Classic. You can specify rules using either rule IDs or security group rule
	// properties. If you use rule properties, the values that you specify (for
	// example, ports) must match the existing rule's values exactly. Each rule has a
	// protocol, from and to ports, and destination (CIDR range, security group, or
	// prefix list). For the TCP and UDP protocols, you must also specify the
	// destination port or range of ports. For the ICMP protocol, you must also specify
	// the ICMP type and code. If the security group rule has a description, you do not
	// need to specify the description to revoke the rule. [Default VPC] If the values
	// you specify do not match the existing rule's values, no error is returned, and
	// the output describes the security group rules that were not revoked. Amazon Web
	// Services recommends that you describe the security group to verify that the
	// rules were removed. Rule changes are propagated to instances within the security
	// group as quickly as possible. However, a small delay might occur.
	//
	RevokeSecurityGroupEgress(arg1 context.Context, arg2 *ec2.RevokeSecurityGroupEgressInput, arg3 ...func(*ec2.Options)) (*ec2.RevokeSecurityGroupEgressOutput, error)
	// Removes the specified inbound (ingress) rules from a security group. You can
	// specify rules using either rule IDs or security group rule properties. If you
	// use rule properties, the values that you specify (for example, ports) must match
	// the existing rule's values exactly. Each rule has a protocol, from and to ports,
	// and source (CIDR range, security group, or prefix list). For the TCP and UDP
	// protocols, you must also specify the destination port or range of ports. For the
	// ICMP protocol, you must also specify the ICMP type and code. If the security
	// group rule has a description, you do not need to specify the description to
	// revoke the rule. [EC2-Classic, default VPC] If the values you specify do not
	// match the existing rule's values, no error is returned, and the output describes
	// the security group rules that were not revoked. Amazon Web Services recommends
	// that you describe the security group to verify that the rules were removed. Rule
	// changes are propagated to instances within the security group as quickly as
	// possible. However, a small delay might occur.
	//
	RevokeSecurityGroupIngress(arg1 context.Context, arg2 *ec2.RevokeSecurityGroupIngressInput, arg3 ...func(*ec2.Options)) (*ec2.RevokeSecurityGroupIngressOutput, error)
	// Launches the specified number of instances using an AMI for which you have
	// permissions. You can specify a number of options, or leave the default options.
	// The following rules apply:
	//
	// * [EC2-VPC] If you don't specify a subnet ID, we
	// choose a default subnet from your default VPC for you. If you don't have a
	// default VPC, you must specify a subnet ID in the request.
	//
	// * [EC2-Classic] If
	// don't specify an Availability Zone, we choose one for you.
	//
	// * Some instance
	// types must be launched into a VPC. If you do not have a default VPC, or if you
	// do not specify a subnet ID, the request fails. For more information, see
	// Instance types available only in a VPC
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/using-vpc.html#vpc-only-instance-types).
	//
	// *
	// [EC2-VPC] All instances have a network interface with a primary private IPv4
	// address. If you don't specify this address, we choose one from the IPv4 range of
	// your subnet.
	//
	// * Not all instance types support IPv6 addresses. For more
	// information, see Instance types
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instance-types.html).
	//
	// * If
	// you don't specify a security group ID, we use the default security group. For
	// more information, see Security groups
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/using-network-security.html).
	//
	// *
	// If any of the AMIs have a product code attached for which the user has not
	// subscribed, the request fails.
	//
	// You can create a launch template
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-launch-templates.html),
	// which is a resource that contains the parameters to launch an instance. When you
	// launch an instance using RunInstances, you can specify the launch template
	// instead of specifying the launch parameters. To ensure faster instance launches,
	// break up large requests into smaller batches. For example, create five separate
	// launch requests for 100 instances each instead of one launch request for 500
	// instances. An instance is ready for you to use when it's in the running state.
	// You can check the state of your instance using DescribeInstances. You can tag
	// instances and EBS volumes during launch, after launch, or both. For more
	// information, see CreateTags and Tagging your Amazon EC2 resources
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/Using_Tags.html). Linux
	// instances have access to the public key of the key pair at boot. You can use
	// this key to provide secure access to the instance. Amazon EC2 public images use
	// this feature to provide secure access without passwords. For more information,
	// see Key pairs
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-key-pairs.html). For
	// troubleshooting, see What to do if an instance immediately terminates
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/Using_InstanceStraightToTerminated.html),
	// and Troubleshooting connecting to your instance
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/TroubleshootingInstancesConnecting.html).
	//
	RunInstances(arg1 context.Context, arg2 *ec2.RunInstancesInput, arg3 ...func(*ec2.Options)) (*ec2.RunInstancesOutput, error)
	// Launches the specified Scheduled Instances. Before you can launch a Scheduled
	// Instance, you must purchase it and obtain an identifier using
	// PurchaseScheduledInstances. You must launch a Scheduled Instance during its
	// scheduled time period. You can't stop or reboot a Scheduled Instance, but you
	// can terminate it as needed. If you terminate a Scheduled Instance before the
	// current scheduled time period ends, you can launch it again after a few minutes.
	// For more information, see Scheduled Instances
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-scheduled-instances.html)
	// in the Amazon EC2 User Guide.
	//
	RunScheduledInstances(arg1 context.Context, arg2 *ec2.RunScheduledInstancesInput, arg3 ...func(*ec2.Options)) (*ec2.RunScheduledInstancesOutput, error)
	// Searches for routes in the specified local gateway route table.
	//
	SearchLocalGatewayRoutes(arg1 context.Context, arg2 *ec2.SearchLocalGatewayRoutesInput, arg3 ...func(*ec2.Options)) (*ec2.SearchLocalGatewayRoutesOutput, error)
	// Searches one or more transit gateway multicast groups and returns the group
	// membership information.
	//
	SearchTransitGatewayMulticastGroups(arg1 context.Context, arg2 *ec2.SearchTransitGatewayMulticastGroupsInput, arg3 ...func(*ec2.Options)) (*ec2.SearchTransitGatewayMulticastGroupsOutput, error)
	// Searches for routes in the specified transit gateway route table.
	//
	SearchTransitGatewayRoutes(arg1 context.Context, arg2 *ec2.SearchTransitGatewayRoutesInput, arg3 ...func(*ec2.Options)) (*ec2.SearchTransitGatewayRoutesOutput, error)
	// Sends a diagnostic interrupt to the specified Amazon EC2 instance to trigger a
	// kernel panic (on Linux instances), or a blue screen/stop error (on Windows
	// instances). For instances based on Intel and AMD processors, the interrupt is
	// received as a non-maskable interrupt (NMI). In general, the operating system
	// crashes and reboots when a kernel panic or stop error is triggered. The
	// operating system can also be configured to perform diagnostic tasks, such as
	// generating a memory dump file, loading a secondary kernel, or obtaining a call
	// trace. Before sending a diagnostic interrupt to your instance, ensure that its
	// operating system is configured to perform the required diagnostic tasks. For
	// more information about configuring your operating system to generate a crash
	// dump when a kernel panic or stop error occurs, see Send a diagnostic interrupt
	// (for advanced users)
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/diagnostic-interrupt.html)
	// (Linux instances) or Send a diagnostic interrupt (for advanced users)
	// (https://docs.aws.amazon.com/AWSEC2/latest/WindowsGuide/diagnostic-interrupt.html)
	// (Windows instances).
	//
	SendDiagnosticInterrupt(arg1 context.Context, arg2 *ec2.SendDiagnosticInterruptInput, arg3 ...func(*ec2.Options)) (*ec2.SendDiagnosticInterruptOutput, error)
	// Starts an Amazon EBS-backed instance that you've previously stopped. Instances
	// that use Amazon EBS volumes as their root devices can be quickly stopped and
	// started. When an instance is stopped, the compute resources are released and you
	// are not billed for instance usage. However, your root partition Amazon EBS
	// volume remains and continues to persist your data, and you are charged for
	// Amazon EBS volume usage. You can restart your instance at any time. Every time
	// you start your instance, Amazon EC2 charges a one-minute minimum for instance
	// usage, and thereafter charges per second for instance usage. Before stopping an
	// instance, make sure it is in a state from which it can be restarted. Stopping an
	// instance does not preserve data stored in RAM. Performing this operation on an
	// instance that uses an instance store as its root device returns an error. If you
	// attempt to start a T3 instance with host tenancy and the unlimted CPU credit
	// option, the request fails. The unlimited CPU credit option is not supported on
	// Dedicated Hosts. Before you start the instance, either change its CPU credit
	// option to standard, or change its tenancy to default or dedicated. For more
	// information, see Stop and start your instance
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/Stop_Start.html) in the
	// Amazon EC2 User Guide.
	//
	StartInstances(arg1 context.Context, arg2 *ec2.StartInstancesInput, arg3 ...func(*ec2.Options)) (*ec2.StartInstancesOutput, error)
	// Starts analyzing the specified Network Access Scope.
	//
	StartNetworkInsightsAccessScopeAnalysis(arg1 context.Context, arg2 *ec2.StartNetworkInsightsAccessScopeAnalysisInput, arg3 ...func(*ec2.Options)) (*ec2.StartNetworkInsightsAccessScopeAnalysisOutput, error)
	// Starts analyzing the specified path. If the path is reachable, the operation
	// returns the shortest feasible path.
	//
	StartNetworkInsightsAnalysis(arg1 context.Context, arg2 *ec2.StartNetworkInsightsAnalysisInput, arg3 ...func(*ec2.Options)) (*ec2.StartNetworkInsightsAnalysisOutput, error)
	// Initiates the verification process to prove that the service provider owns the
	// private DNS name domain for the endpoint service. The service provider must
	// successfully perform the verification before the consumer can use the name to
	// access the service. Before the service provider runs this command, they must add
	// a record to the DNS server.
	//
	StartVpcEndpointServicePrivateDnsVerification(arg1 context.Context, arg2 *ec2.StartVpcEndpointServicePrivateDnsVerificationInput, arg3 ...func(*ec2.Options)) (*ec2.StartVpcEndpointServicePrivateDnsVerificationOutput, error)
	// Stops an Amazon EBS-backed instance. For more information, see Stop and start
	// your instance
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/Stop_Start.html) in the
	// Amazon EC2 User Guide. You can use the Stop action to hibernate an instance if
	// the instance is enabled for hibernation
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/Hibernate.html#enabling-hibernation)
	// and it meets the hibernation prerequisites
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/Hibernate.html#hibernating-prerequisites).
	// For more information, see Hibernate your instance
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/Hibernate.html) in the
	// Amazon EC2 User Guide. We don't charge usage for a stopped instance, or data
	// transfer fees; however, your root partition Amazon EBS volume remains and
	// continues to persist your data, and you are charged for Amazon EBS volume usage.
	// Every time you start your instance, Amazon EC2 charges a one-minute minimum for
	// instance usage, and thereafter charges per second for instance usage. You can't
	// stop or hibernate instance store-backed instances. You can't use the Stop action
	// to hibernate Spot Instances, but you can specify that Amazon EC2 should
	// hibernate Spot Instances when they are interrupted. For more information, see
	// Hibernating interrupted Spot Instances
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/spot-interruptions.html#hibernate-spot-instances)
	// in the Amazon EC2 User Guide. When you stop or hibernate an instance, we shut it
	// down. You can restart your instance at any time. Before stopping or hibernating
	// an instance, make sure it is in a state from which it can be restarted. Stopping
	// an instance does not preserve data stored in RAM, but hibernating an instance
	// does preserve data stored in RAM. If an instance cannot hibernate successfully,
	// a normal shutdown occurs. Stopping and hibernating an instance is different to
	// rebooting or terminating it. For example, when you stop or hibernate an
	// instance, the root device and any other devices attached to the instance
	// persist. When you terminate an instance, the root device and any other devices
	// attached during the instance launch are automatically deleted. For more
	// information about the differences between rebooting, stopping, hibernating, and
	// terminating instances, see Instance lifecycle
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-instance-lifecycle.html)
	// in the Amazon EC2 User Guide. When you stop an instance, we attempt to shut it
	// down forcibly after a short while. If your instance appears stuck in the
	// stopping state after a period of time, there may be an issue with the underlying
	// host computer. For more information, see Troubleshoot stopping your instance
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/TroubleshootingInstancesStopping.html)
	// in the Amazon EC2 User Guide.
	//
	StopInstances(arg1 context.Context, arg2 *ec2.StopInstancesInput, arg3 ...func(*ec2.Options)) (*ec2.StopInstancesOutput, error)
	// Terminates active Client VPN endpoint connections. This action can be used to
	// terminate a specific client connection, or up to five connections established by
	// a specific user.
	//
	TerminateClientVpnConnections(arg1 context.Context, arg2 *ec2.TerminateClientVpnConnectionsInput, arg3 ...func(*ec2.Options)) (*ec2.TerminateClientVpnConnectionsOutput, error)
	// Shuts down the specified instances. This operation is idempotent; if you
	// terminate an instance more than once, each call succeeds. If you specify
	// multiple instances and the request fails (for example, because of a single
	// incorrect instance ID), none of the instances are terminated. If you terminate
	// multiple instances across multiple Availability Zones, and one or more of the
	// specified instances are enabled for termination protection, the request fails
	// with the following results:
	//
	// * The specified instances that are in the same
	// Availability Zone as the protected instance are not terminated.
	//
	// * The specified
	// instances that are in different Availability Zones, where no other specified
	// instances are protected, are successfully terminated.
	//
	// For example, say you have
	// the following instances:
	//
	// * Instance A: us-east-1a; Not protected
	//
	// * Instance B:
	// us-east-1a; Not protected
	//
	// * Instance C: us-east-1b; Protected
	//
	// * Instance D:
	// us-east-1b; not protected
	//
	// If you attempt to terminate all of these instances in
	// the same request, the request reports failure with the following results:
	//
	// *
	// Instance A and Instance B are successfully terminated because none of the
	// specified instances in us-east-1a are enabled for termination protection.
	//
	// *
	// Instance C and Instance D fail to terminate because at least one of the
	// specified instances in us-east-1b (Instance C) is enabled for termination
	// protection.
	//
	// Terminated instances remain visible after termination (for
	// approximately one hour). By default, Amazon EC2 deletes all EBS volumes that
	// were attached when the instance launched. Volumes attached after instance launch
	// continue running. You can stop, start, and terminate EBS-backed instances. You
	// can only terminate instance store-backed instances. What happens to an instance
	// differs if you stop it or terminate it. For example, when you stop an instance,
	// the root device and any other devices attached to the instance persist. When you
	// terminate an instance, any attached EBS volumes with the DeleteOnTermination
	// block device mapping parameter set to true are automatically deleted. For more
	// information about the differences between stopping and terminating instances,
	// see Instance lifecycle
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-instance-lifecycle.html)
	// in the Amazon EC2 User Guide. For more information about troubleshooting, see
	// Troubleshooting terminating your instance
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/TroubleshootingInstancesShuttingDown.html)
	// in the Amazon EC2 User Guide.
	//
	TerminateInstances(arg1 context.Context, arg2 *ec2.TerminateInstancesInput, arg3 ...func(*ec2.Options)) (*ec2.TerminateInstancesOutput, error)
	// Unassigns one or more IPv6 addresses IPv4 Prefix Delegation prefixes from a
	// network interface.
	//
	UnassignIpv6Addresses(arg1 context.Context, arg2 *ec2.UnassignIpv6AddressesInput, arg3 ...func(*ec2.Options)) (*ec2.UnassignIpv6AddressesOutput, error)
	// Unassigns one or more secondary private IP addresses, or IPv4 Prefix Delegation
	// prefixes from a network interface.
	//
	UnassignPrivateIpAddresses(arg1 context.Context, arg2 *ec2.UnassignPrivateIpAddressesInput, arg3 ...func(*ec2.Options)) (*ec2.UnassignPrivateIpAddressesOutput, error)
	// Disables detailed monitoring for a running instance. For more information, see
	// Monitoring your instances and volumes
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/using-cloudwatch.html) in
	// the Amazon EC2 User Guide.
	//
	UnmonitorInstances(arg1 context.Context, arg2 *ec2.UnmonitorInstancesInput, arg3 ...func(*ec2.Options)) (*ec2.UnmonitorInstancesOutput, error)
	// [VPC only] Updates the description of an egress (outbound) security group rule.
	// You can replace an existing description, or add a description to a rule that did
	// not have one previously. You can remove a description for a security group rule
	// by omitting the description parameter in the request.
	//
	UpdateSecurityGroupRuleDescriptionsEgress(arg1 context.Context, arg2 *ec2.UpdateSecurityGroupRuleDescriptionsEgressInput, arg3 ...func(*ec2.Options)) (*ec2.UpdateSecurityGroupRuleDescriptionsEgressOutput, error)
	// Updates the description of an ingress (inbound) security group rule. You can
	// replace an existing description, or add a description to a rule that did not
	// have one previously. You can remove a description for a security group rule by
	// omitting the description parameter in the request.
	//
	UpdateSecurityGroupRuleDescriptionsIngress(arg1 context.Context, arg2 *ec2.UpdateSecurityGroupRuleDescriptionsIngressInput, arg3 ...func(*ec2.Options)) (*ec2.UpdateSecurityGroupRuleDescriptionsIngressOutput, error)
	// Stops advertising an address range that is provisioned as an address pool. You
	// can perform this operation at most once every 10 seconds, even if you specify
	// different address ranges each time. It can take a few minutes before traffic to
	// the specified addresses stops routing to Amazon Web Services because of BGP
	// propagation delays.
	//
	WithdrawByoipCidr(arg1 context.Context, arg2 *ec2.WithdrawByoipCidrInput, arg3 ...func(*ec2.Options)) (*ec2.WithdrawByoipCidrOutput, error)
}

// EC2ClientMock generic client mock
type EC2ClientMock struct {
	AcceptReservedInstancesExchangeQuoteMock                            func(arg1 context.Context, arg2 *ec2.AcceptReservedInstancesExchangeQuoteInput, arg3 ...func(*ec2.Options)) (*ec2.AcceptReservedInstancesExchangeQuoteOutput, error)
	AcceptTransitGatewayMulticastDomainAssociationsMock                 func(arg1 context.Context, arg2 *ec2.AcceptTransitGatewayMulticastDomainAssociationsInput, arg3 ...func(*ec2.Options)) (*ec2.AcceptTransitGatewayMulticastDomainAssociationsOutput, error)
	AcceptTransitGatewayPeeringAttachmentMock                           func(arg1 context.Context, arg2 *ec2.AcceptTransitGatewayPeeringAttachmentInput, arg3 ...func(*ec2.Options)) (*ec2.AcceptTransitGatewayPeeringAttachmentOutput, error)
	AcceptTransitGatewayVpcAttachmentMock                               func(arg1 context.Context, arg2 *ec2.AcceptTransitGatewayVpcAttachmentInput, arg3 ...func(*ec2.Options)) (*ec2.AcceptTransitGatewayVpcAttachmentOutput, error)
	AcceptVpcEndpointConnectionsMock                                    func(arg1 context.Context, arg2 *ec2.AcceptVpcEndpointConnectionsInput, arg3 ...func(*ec2.Options)) (*ec2.AcceptVpcEndpointConnectionsOutput, error)
	AcceptVpcPeeringConnectionMock                                      func(arg1 context.Context, arg2 *ec2.AcceptVpcPeeringConnectionInput, arg3 ...func(*ec2.Options)) (*ec2.AcceptVpcPeeringConnectionOutput, error)
	AdvertiseByoipCidrMock                                              func(arg1 context.Context, arg2 *ec2.AdvertiseByoipCidrInput, arg3 ...func(*ec2.Options)) (*ec2.AdvertiseByoipCidrOutput, error)
	AllocateAddressMock                                                 func(arg1 context.Context, arg2 *ec2.AllocateAddressInput, arg3 ...func(*ec2.Options)) (*ec2.AllocateAddressOutput, error)
	AllocateHostsMock                                                   func(arg1 context.Context, arg2 *ec2.AllocateHostsInput, arg3 ...func(*ec2.Options)) (*ec2.AllocateHostsOutput, error)
	AllocateIpamPoolCidrMock                                            func(arg1 context.Context, arg2 *ec2.AllocateIpamPoolCidrInput, arg3 ...func(*ec2.Options)) (*ec2.AllocateIpamPoolCidrOutput, error)
	ApplySecurityGroupsToClientVpnTargetNetworkMock                     func(arg1 context.Context, arg2 *ec2.ApplySecurityGroupsToClientVpnTargetNetworkInput, arg3 ...func(*ec2.Options)) (*ec2.ApplySecurityGroupsToClientVpnTargetNetworkOutput, error)
	AssignIpv6AddressesMock                                             func(arg1 context.Context, arg2 *ec2.AssignIpv6AddressesInput, arg3 ...func(*ec2.Options)) (*ec2.AssignIpv6AddressesOutput, error)
	AssignPrivateIpAddressesMock                                        func(arg1 context.Context, arg2 *ec2.AssignPrivateIpAddressesInput, arg3 ...func(*ec2.Options)) (*ec2.AssignPrivateIpAddressesOutput, error)
	AssociateAddressMock                                                func(arg1 context.Context, arg2 *ec2.AssociateAddressInput, arg3 ...func(*ec2.Options)) (*ec2.AssociateAddressOutput, error)
	AssociateClientVpnTargetNetworkMock                                 func(arg1 context.Context, arg2 *ec2.AssociateClientVpnTargetNetworkInput, arg3 ...func(*ec2.Options)) (*ec2.AssociateClientVpnTargetNetworkOutput, error)
	AssociateDhcpOptionsMock                                            func(arg1 context.Context, arg2 *ec2.AssociateDhcpOptionsInput, arg3 ...func(*ec2.Options)) (*ec2.AssociateDhcpOptionsOutput, error)
	AssociateEnclaveCertificateIamRoleMock                              func(arg1 context.Context, arg2 *ec2.AssociateEnclaveCertificateIamRoleInput, arg3 ...func(*ec2.Options)) (*ec2.AssociateEnclaveCertificateIamRoleOutput, error)
	AssociateIamInstanceProfileMock                                     func(arg1 context.Context, arg2 *ec2.AssociateIamInstanceProfileInput, arg3 ...func(*ec2.Options)) (*ec2.AssociateIamInstanceProfileOutput, error)
	AssociateInstanceEventWindowMock                                    func(arg1 context.Context, arg2 *ec2.AssociateInstanceEventWindowInput, arg3 ...func(*ec2.Options)) (*ec2.AssociateInstanceEventWindowOutput, error)
	AssociateRouteTableMock                                             func(arg1 context.Context, arg2 *ec2.AssociateRouteTableInput, arg3 ...func(*ec2.Options)) (*ec2.AssociateRouteTableOutput, error)
	AssociateSubnetCidrBlockMock                                        func(arg1 context.Context, arg2 *ec2.AssociateSubnetCidrBlockInput, arg3 ...func(*ec2.Options)) (*ec2.AssociateSubnetCidrBlockOutput, error)
	AssociateTransitGatewayMulticastDomainMock                          func(arg1 context.Context, arg2 *ec2.AssociateTransitGatewayMulticastDomainInput, arg3 ...func(*ec2.Options)) (*ec2.AssociateTransitGatewayMulticastDomainOutput, error)
	AssociateTransitGatewayPolicyTableMock                              func(arg1 context.Context, arg2 *ec2.AssociateTransitGatewayPolicyTableInput, arg3 ...func(*ec2.Options)) (*ec2.AssociateTransitGatewayPolicyTableOutput, error)
	AssociateTransitGatewayRouteTableMock                               func(arg1 context.Context, arg2 *ec2.AssociateTransitGatewayRouteTableInput, arg3 ...func(*ec2.Options)) (*ec2.AssociateTransitGatewayRouteTableOutput, error)
	AssociateTrunkInterfaceMock                                         func(arg1 context.Context, arg2 *ec2.AssociateTrunkInterfaceInput, arg3 ...func(*ec2.Options)) (*ec2.AssociateTrunkInterfaceOutput, error)
	AssociateVpcCidrBlockMock                                           func(arg1 context.Context, arg2 *ec2.AssociateVpcCidrBlockInput, arg3 ...func(*ec2.Options)) (*ec2.AssociateVpcCidrBlockOutput, error)
	AttachClassicLinkVpcMock                                            func(arg1 context.Context, arg2 *ec2.AttachClassicLinkVpcInput, arg3 ...func(*ec2.Options)) (*ec2.AttachClassicLinkVpcOutput, error)
	AttachInternetGatewayMock                                           func(arg1 context.Context, arg2 *ec2.AttachInternetGatewayInput, arg3 ...func(*ec2.Options)) (*ec2.AttachInternetGatewayOutput, error)
	AttachNetworkInterfaceMock                                          func(arg1 context.Context, arg2 *ec2.AttachNetworkInterfaceInput, arg3 ...func(*ec2.Options)) (*ec2.AttachNetworkInterfaceOutput, error)
	AttachVolumeMock                                                    func(arg1 context.Context, arg2 *ec2.AttachVolumeInput, arg3 ...func(*ec2.Options)) (*ec2.AttachVolumeOutput, error)
	AttachVpnGatewayMock                                                func(arg1 context.Context, arg2 *ec2.AttachVpnGatewayInput, arg3 ...func(*ec2.Options)) (*ec2.AttachVpnGatewayOutput, error)
	AuthorizeClientVpnIngressMock                                       func(arg1 context.Context, arg2 *ec2.AuthorizeClientVpnIngressInput, arg3 ...func(*ec2.Options)) (*ec2.AuthorizeClientVpnIngressOutput, error)
	AuthorizeSecurityGroupEgressMock                                    func(arg1 context.Context, arg2 *ec2.AuthorizeSecurityGroupEgressInput, arg3 ...func(*ec2.Options)) (*ec2.AuthorizeSecurityGroupEgressOutput, error)
	AuthorizeSecurityGroupIngressMock                                   func(arg1 context.Context, arg2 *ec2.AuthorizeSecurityGroupIngressInput, arg3 ...func(*ec2.Options)) (*ec2.AuthorizeSecurityGroupIngressOutput, error)
	BundleInstanceMock                                                  func(arg1 context.Context, arg2 *ec2.BundleInstanceInput, arg3 ...func(*ec2.Options)) (*ec2.BundleInstanceOutput, error)
	CancelBundleTaskMock                                                func(arg1 context.Context, arg2 *ec2.CancelBundleTaskInput, arg3 ...func(*ec2.Options)) (*ec2.CancelBundleTaskOutput, error)
	CancelCapacityReservationMock                                       func(arg1 context.Context, arg2 *ec2.CancelCapacityReservationInput, arg3 ...func(*ec2.Options)) (*ec2.CancelCapacityReservationOutput, error)
	CancelCapacityReservationFleetsMock                                 func(arg1 context.Context, arg2 *ec2.CancelCapacityReservationFleetsInput, arg3 ...func(*ec2.Options)) (*ec2.CancelCapacityReservationFleetsOutput, error)
	CancelConversionTaskMock                                            func(arg1 context.Context, arg2 *ec2.CancelConversionTaskInput, arg3 ...func(*ec2.Options)) (*ec2.CancelConversionTaskOutput, error)
	CancelExportTaskMock                                                func(arg1 context.Context, arg2 *ec2.CancelExportTaskInput, arg3 ...func(*ec2.Options)) (*ec2.CancelExportTaskOutput, error)
	CancelImportTaskMock                                                func(arg1 context.Context, arg2 *ec2.CancelImportTaskInput, arg3 ...func(*ec2.Options)) (*ec2.CancelImportTaskOutput, error)
	CancelReservedInstancesListingMock                                  func(arg1 context.Context, arg2 *ec2.CancelReservedInstancesListingInput, arg3 ...func(*ec2.Options)) (*ec2.CancelReservedInstancesListingOutput, error)
	CancelSpotFleetRequestsMock                                         func(arg1 context.Context, arg2 *ec2.CancelSpotFleetRequestsInput, arg3 ...func(*ec2.Options)) (*ec2.CancelSpotFleetRequestsOutput, error)
	CancelSpotInstanceRequestsMock                                      func(arg1 context.Context, arg2 *ec2.CancelSpotInstanceRequestsInput, arg3 ...func(*ec2.Options)) (*ec2.CancelSpotInstanceRequestsOutput, error)
	ConfirmProductInstanceMock                                          func(arg1 context.Context, arg2 *ec2.ConfirmProductInstanceInput, arg3 ...func(*ec2.Options)) (*ec2.ConfirmProductInstanceOutput, error)
	CopyFpgaImageMock                                                   func(arg1 context.Context, arg2 *ec2.CopyFpgaImageInput, arg3 ...func(*ec2.Options)) (*ec2.CopyFpgaImageOutput, error)
	CopyImageMock                                                       func(arg1 context.Context, arg2 *ec2.CopyImageInput, arg3 ...func(*ec2.Options)) (*ec2.CopyImageOutput, error)
	CopySnapshotMock                                                    func(arg1 context.Context, arg2 *ec2.CopySnapshotInput, arg3 ...func(*ec2.Options)) (*ec2.CopySnapshotOutput, error)
	CreateCapacityReservationMock                                       func(arg1 context.Context, arg2 *ec2.CreateCapacityReservationInput, arg3 ...func(*ec2.Options)) (*ec2.CreateCapacityReservationOutput, error)
	CreateCapacityReservationFleetMock                                  func(arg1 context.Context, arg2 *ec2.CreateCapacityReservationFleetInput, arg3 ...func(*ec2.Options)) (*ec2.CreateCapacityReservationFleetOutput, error)
	CreateCarrierGatewayMock                                            func(arg1 context.Context, arg2 *ec2.CreateCarrierGatewayInput, arg3 ...func(*ec2.Options)) (*ec2.CreateCarrierGatewayOutput, error)
	CreateClientVpnEndpointMock                                         func(arg1 context.Context, arg2 *ec2.CreateClientVpnEndpointInput, arg3 ...func(*ec2.Options)) (*ec2.CreateClientVpnEndpointOutput, error)
	CreateClientVpnRouteMock                                            func(arg1 context.Context, arg2 *ec2.CreateClientVpnRouteInput, arg3 ...func(*ec2.Options)) (*ec2.CreateClientVpnRouteOutput, error)
	CreateCustomerGatewayMock                                           func(arg1 context.Context, arg2 *ec2.CreateCustomerGatewayInput, arg3 ...func(*ec2.Options)) (*ec2.CreateCustomerGatewayOutput, error)
	CreateDefaultSubnetMock                                             func(arg1 context.Context, arg2 *ec2.CreateDefaultSubnetInput, arg3 ...func(*ec2.Options)) (*ec2.CreateDefaultSubnetOutput, error)
	CreateDefaultVpcMock                                                func(arg1 context.Context, arg2 *ec2.CreateDefaultVpcInput, arg3 ...func(*ec2.Options)) (*ec2.CreateDefaultVpcOutput, error)
	CreateDhcpOptionsMock                                               func(arg1 context.Context, arg2 *ec2.CreateDhcpOptionsInput, arg3 ...func(*ec2.Options)) (*ec2.CreateDhcpOptionsOutput, error)
	CreateEgressOnlyInternetGatewayMock                                 func(arg1 context.Context, arg2 *ec2.CreateEgressOnlyInternetGatewayInput, arg3 ...func(*ec2.Options)) (*ec2.CreateEgressOnlyInternetGatewayOutput, error)
	CreateFleetMock                                                     func(arg1 context.Context, arg2 *ec2.CreateFleetInput, arg3 ...func(*ec2.Options)) (*ec2.CreateFleetOutput, error)
	CreateFlowLogsMock                                                  func(arg1 context.Context, arg2 *ec2.CreateFlowLogsInput, arg3 ...func(*ec2.Options)) (*ec2.CreateFlowLogsOutput, error)
	CreateFpgaImageMock                                                 func(arg1 context.Context, arg2 *ec2.CreateFpgaImageInput, arg3 ...func(*ec2.Options)) (*ec2.CreateFpgaImageOutput, error)
	CreateImageMock                                                     func(arg1 context.Context, arg2 *ec2.CreateImageInput, arg3 ...func(*ec2.Options)) (*ec2.CreateImageOutput, error)
	CreateInstanceEventWindowMock                                       func(arg1 context.Context, arg2 *ec2.CreateInstanceEventWindowInput, arg3 ...func(*ec2.Options)) (*ec2.CreateInstanceEventWindowOutput, error)
	CreateInstanceExportTaskMock                                        func(arg1 context.Context, arg2 *ec2.CreateInstanceExportTaskInput, arg3 ...func(*ec2.Options)) (*ec2.CreateInstanceExportTaskOutput, error)
	CreateInternetGatewayMock                                           func(arg1 context.Context, arg2 *ec2.CreateInternetGatewayInput, arg3 ...func(*ec2.Options)) (*ec2.CreateInternetGatewayOutput, error)
	CreateIpamMock                                                      func(arg1 context.Context, arg2 *ec2.CreateIpamInput, arg3 ...func(*ec2.Options)) (*ec2.CreateIpamOutput, error)
	CreateIpamPoolMock                                                  func(arg1 context.Context, arg2 *ec2.CreateIpamPoolInput, arg3 ...func(*ec2.Options)) (*ec2.CreateIpamPoolOutput, error)
	CreateIpamScopeMock                                                 func(arg1 context.Context, arg2 *ec2.CreateIpamScopeInput, arg3 ...func(*ec2.Options)) (*ec2.CreateIpamScopeOutput, error)
	CreateKeyPairMock                                                   func(arg1 context.Context, arg2 *ec2.CreateKeyPairInput, arg3 ...func(*ec2.Options)) (*ec2.CreateKeyPairOutput, error)
	CreateLaunchTemplateMock                                            func(arg1 context.Context, arg2 *ec2.CreateLaunchTemplateInput, arg3 ...func(*ec2.Options)) (*ec2.CreateLaunchTemplateOutput, error)
	CreateLaunchTemplateVersionMock                                     func(arg1 context.Context, arg2 *ec2.CreateLaunchTemplateVersionInput, arg3 ...func(*ec2.Options)) (*ec2.CreateLaunchTemplateVersionOutput, error)
	CreateLocalGatewayRouteMock                                         func(arg1 context.Context, arg2 *ec2.CreateLocalGatewayRouteInput, arg3 ...func(*ec2.Options)) (*ec2.CreateLocalGatewayRouteOutput, error)
	CreateLocalGatewayRouteTableVpcAssociationMock                      func(arg1 context.Context, arg2 *ec2.CreateLocalGatewayRouteTableVpcAssociationInput, arg3 ...func(*ec2.Options)) (*ec2.CreateLocalGatewayRouteTableVpcAssociationOutput, error)
	CreateManagedPrefixListMock                                         func(arg1 context.Context, arg2 *ec2.CreateManagedPrefixListInput, arg3 ...func(*ec2.Options)) (*ec2.CreateManagedPrefixListOutput, error)
	CreateNatGatewayMock                                                func(arg1 context.Context, arg2 *ec2.CreateNatGatewayInput, arg3 ...func(*ec2.Options)) (*ec2.CreateNatGatewayOutput, error)
	CreateNetworkAclMock                                                func(arg1 context.Context, arg2 *ec2.CreateNetworkAclInput, arg3 ...func(*ec2.Options)) (*ec2.CreateNetworkAclOutput, error)
	CreateNetworkAclEntryMock                                           func(arg1 context.Context, arg2 *ec2.CreateNetworkAclEntryInput, arg3 ...func(*ec2.Options)) (*ec2.CreateNetworkAclEntryOutput, error)
	CreateNetworkInsightsAccessScopeMock                                func(arg1 context.Context, arg2 *ec2.CreateNetworkInsightsAccessScopeInput, arg3 ...func(*ec2.Options)) (*ec2.CreateNetworkInsightsAccessScopeOutput, error)
	CreateNetworkInsightsPathMock                                       func(arg1 context.Context, arg2 *ec2.CreateNetworkInsightsPathInput, arg3 ...func(*ec2.Options)) (*ec2.CreateNetworkInsightsPathOutput, error)
	CreateNetworkInterfaceMock                                          func(arg1 context.Context, arg2 *ec2.CreateNetworkInterfaceInput, arg3 ...func(*ec2.Options)) (*ec2.CreateNetworkInterfaceOutput, error)
	CreateNetworkInterfacePermissionMock                                func(arg1 context.Context, arg2 *ec2.CreateNetworkInterfacePermissionInput, arg3 ...func(*ec2.Options)) (*ec2.CreateNetworkInterfacePermissionOutput, error)
	CreatePlacementGroupMock                                            func(arg1 context.Context, arg2 *ec2.CreatePlacementGroupInput, arg3 ...func(*ec2.Options)) (*ec2.CreatePlacementGroupOutput, error)
	CreatePublicIpv4PoolMock                                            func(arg1 context.Context, arg2 *ec2.CreatePublicIpv4PoolInput, arg3 ...func(*ec2.Options)) (*ec2.CreatePublicIpv4PoolOutput, error)
	CreateReplaceRootVolumeTaskMock                                     func(arg1 context.Context, arg2 *ec2.CreateReplaceRootVolumeTaskInput, arg3 ...func(*ec2.Options)) (*ec2.CreateReplaceRootVolumeTaskOutput, error)
	CreateReservedInstancesListingMock                                  func(arg1 context.Context, arg2 *ec2.CreateReservedInstancesListingInput, arg3 ...func(*ec2.Options)) (*ec2.CreateReservedInstancesListingOutput, error)
	CreateRestoreImageTaskMock                                          func(arg1 context.Context, arg2 *ec2.CreateRestoreImageTaskInput, arg3 ...func(*ec2.Options)) (*ec2.CreateRestoreImageTaskOutput, error)
	CreateRouteMock                                                     func(arg1 context.Context, arg2 *ec2.CreateRouteInput, arg3 ...func(*ec2.Options)) (*ec2.CreateRouteOutput, error)
	CreateRouteTableMock                                                func(arg1 context.Context, arg2 *ec2.CreateRouteTableInput, arg3 ...func(*ec2.Options)) (*ec2.CreateRouteTableOutput, error)
	CreateSecurityGroupMock                                             func(arg1 context.Context, arg2 *ec2.CreateSecurityGroupInput, arg3 ...func(*ec2.Options)) (*ec2.CreateSecurityGroupOutput, error)
	CreateSnapshotMock                                                  func(arg1 context.Context, arg2 *ec2.CreateSnapshotInput, arg3 ...func(*ec2.Options)) (*ec2.CreateSnapshotOutput, error)
	CreateSnapshotsMock                                                 func(arg1 context.Context, arg2 *ec2.CreateSnapshotsInput, arg3 ...func(*ec2.Options)) (*ec2.CreateSnapshotsOutput, error)
	CreateSpotDatafeedSubscriptionMock                                  func(arg1 context.Context, arg2 *ec2.CreateSpotDatafeedSubscriptionInput, arg3 ...func(*ec2.Options)) (*ec2.CreateSpotDatafeedSubscriptionOutput, error)
	CreateStoreImageTaskMock                                            func(arg1 context.Context, arg2 *ec2.CreateStoreImageTaskInput, arg3 ...func(*ec2.Options)) (*ec2.CreateStoreImageTaskOutput, error)
	CreateSubnetMock                                                    func(arg1 context.Context, arg2 *ec2.CreateSubnetInput, arg3 ...func(*ec2.Options)) (*ec2.CreateSubnetOutput, error)
	CreateSubnetCidrReservationMock                                     func(arg1 context.Context, arg2 *ec2.CreateSubnetCidrReservationInput, arg3 ...func(*ec2.Options)) (*ec2.CreateSubnetCidrReservationOutput, error)
	CreateTagsMock                                                      func(arg1 context.Context, arg2 *ec2.CreateTagsInput, arg3 ...func(*ec2.Options)) (*ec2.CreateTagsOutput, error)
	CreateTrafficMirrorFilterMock                                       func(arg1 context.Context, arg2 *ec2.CreateTrafficMirrorFilterInput, arg3 ...func(*ec2.Options)) (*ec2.CreateTrafficMirrorFilterOutput, error)
	CreateTrafficMirrorFilterRuleMock                                   func(arg1 context.Context, arg2 *ec2.CreateTrafficMirrorFilterRuleInput, arg3 ...func(*ec2.Options)) (*ec2.CreateTrafficMirrorFilterRuleOutput, error)
	CreateTrafficMirrorSessionMock                                      func(arg1 context.Context, arg2 *ec2.CreateTrafficMirrorSessionInput, arg3 ...func(*ec2.Options)) (*ec2.CreateTrafficMirrorSessionOutput, error)
	CreateTrafficMirrorTargetMock                                       func(arg1 context.Context, arg2 *ec2.CreateTrafficMirrorTargetInput, arg3 ...func(*ec2.Options)) (*ec2.CreateTrafficMirrorTargetOutput, error)
	CreateTransitGatewayMock                                            func(arg1 context.Context, arg2 *ec2.CreateTransitGatewayInput, arg3 ...func(*ec2.Options)) (*ec2.CreateTransitGatewayOutput, error)
	CreateTransitGatewayConnectMock                                     func(arg1 context.Context, arg2 *ec2.CreateTransitGatewayConnectInput, arg3 ...func(*ec2.Options)) (*ec2.CreateTransitGatewayConnectOutput, error)
	CreateTransitGatewayConnectPeerMock                                 func(arg1 context.Context, arg2 *ec2.CreateTransitGatewayConnectPeerInput, arg3 ...func(*ec2.Options)) (*ec2.CreateTransitGatewayConnectPeerOutput, error)
	CreateTransitGatewayMulticastDomainMock                             func(arg1 context.Context, arg2 *ec2.CreateTransitGatewayMulticastDomainInput, arg3 ...func(*ec2.Options)) (*ec2.CreateTransitGatewayMulticastDomainOutput, error)
	CreateTransitGatewayPeeringAttachmentMock                           func(arg1 context.Context, arg2 *ec2.CreateTransitGatewayPeeringAttachmentInput, arg3 ...func(*ec2.Options)) (*ec2.CreateTransitGatewayPeeringAttachmentOutput, error)
	CreateTransitGatewayPolicyTableMock                                 func(arg1 context.Context, arg2 *ec2.CreateTransitGatewayPolicyTableInput, arg3 ...func(*ec2.Options)) (*ec2.CreateTransitGatewayPolicyTableOutput, error)
	CreateTransitGatewayPrefixListReferenceMock                         func(arg1 context.Context, arg2 *ec2.CreateTransitGatewayPrefixListReferenceInput, arg3 ...func(*ec2.Options)) (*ec2.CreateTransitGatewayPrefixListReferenceOutput, error)
	CreateTransitGatewayRouteMock                                       func(arg1 context.Context, arg2 *ec2.CreateTransitGatewayRouteInput, arg3 ...func(*ec2.Options)) (*ec2.CreateTransitGatewayRouteOutput, error)
	CreateTransitGatewayRouteTableMock                                  func(arg1 context.Context, arg2 *ec2.CreateTransitGatewayRouteTableInput, arg3 ...func(*ec2.Options)) (*ec2.CreateTransitGatewayRouteTableOutput, error)
	CreateTransitGatewayRouteTableAnnouncementMock                      func(arg1 context.Context, arg2 *ec2.CreateTransitGatewayRouteTableAnnouncementInput, arg3 ...func(*ec2.Options)) (*ec2.CreateTransitGatewayRouteTableAnnouncementOutput, error)
	CreateTransitGatewayVpcAttachmentMock                               func(arg1 context.Context, arg2 *ec2.CreateTransitGatewayVpcAttachmentInput, arg3 ...func(*ec2.Options)) (*ec2.CreateTransitGatewayVpcAttachmentOutput, error)
	CreateVolumeMock                                                    func(arg1 context.Context, arg2 *ec2.CreateVolumeInput, arg3 ...func(*ec2.Options)) (*ec2.CreateVolumeOutput, error)
	CreateVpcMock                                                       func(arg1 context.Context, arg2 *ec2.CreateVpcInput, arg3 ...func(*ec2.Options)) (*ec2.CreateVpcOutput, error)
	CreateVpcEndpointMock                                               func(arg1 context.Context, arg2 *ec2.CreateVpcEndpointInput, arg3 ...func(*ec2.Options)) (*ec2.CreateVpcEndpointOutput, error)
	CreateVpcEndpointConnectionNotificationMock                         func(arg1 context.Context, arg2 *ec2.CreateVpcEndpointConnectionNotificationInput, arg3 ...func(*ec2.Options)) (*ec2.CreateVpcEndpointConnectionNotificationOutput, error)
	CreateVpcEndpointServiceConfigurationMock                           func(arg1 context.Context, arg2 *ec2.CreateVpcEndpointServiceConfigurationInput, arg3 ...func(*ec2.Options)) (*ec2.CreateVpcEndpointServiceConfigurationOutput, error)
	CreateVpcPeeringConnectionMock                                      func(arg1 context.Context, arg2 *ec2.CreateVpcPeeringConnectionInput, arg3 ...func(*ec2.Options)) (*ec2.CreateVpcPeeringConnectionOutput, error)
	CreateVpnConnectionMock                                             func(arg1 context.Context, arg2 *ec2.CreateVpnConnectionInput, arg3 ...func(*ec2.Options)) (*ec2.CreateVpnConnectionOutput, error)
	CreateVpnConnectionRouteMock                                        func(arg1 context.Context, arg2 *ec2.CreateVpnConnectionRouteInput, arg3 ...func(*ec2.Options)) (*ec2.CreateVpnConnectionRouteOutput, error)
	CreateVpnGatewayMock                                                func(arg1 context.Context, arg2 *ec2.CreateVpnGatewayInput, arg3 ...func(*ec2.Options)) (*ec2.CreateVpnGatewayOutput, error)
	DeleteCarrierGatewayMock                                            func(arg1 context.Context, arg2 *ec2.DeleteCarrierGatewayInput, arg3 ...func(*ec2.Options)) (*ec2.DeleteCarrierGatewayOutput, error)
	DeleteClientVpnEndpointMock                                         func(arg1 context.Context, arg2 *ec2.DeleteClientVpnEndpointInput, arg3 ...func(*ec2.Options)) (*ec2.DeleteClientVpnEndpointOutput, error)
	DeleteClientVpnRouteMock                                            func(arg1 context.Context, arg2 *ec2.DeleteClientVpnRouteInput, arg3 ...func(*ec2.Options)) (*ec2.DeleteClientVpnRouteOutput, error)
	DeleteCustomerGatewayMock                                           func(arg1 context.Context, arg2 *ec2.DeleteCustomerGatewayInput, arg3 ...func(*ec2.Options)) (*ec2.DeleteCustomerGatewayOutput, error)
	DeleteDhcpOptionsMock                                               func(arg1 context.Context, arg2 *ec2.DeleteDhcpOptionsInput, arg3 ...func(*ec2.Options)) (*ec2.DeleteDhcpOptionsOutput, error)
	DeleteEgressOnlyInternetGatewayMock                                 func(arg1 context.Context, arg2 *ec2.DeleteEgressOnlyInternetGatewayInput, arg3 ...func(*ec2.Options)) (*ec2.DeleteEgressOnlyInternetGatewayOutput, error)
	DeleteFleetsMock                                                    func(arg1 context.Context, arg2 *ec2.DeleteFleetsInput, arg3 ...func(*ec2.Options)) (*ec2.DeleteFleetsOutput, error)
	DeleteFlowLogsMock                                                  func(arg1 context.Context, arg2 *ec2.DeleteFlowLogsInput, arg3 ...func(*ec2.Options)) (*ec2.DeleteFlowLogsOutput, error)
	DeleteFpgaImageMock                                                 func(arg1 context.Context, arg2 *ec2.DeleteFpgaImageInput, arg3 ...func(*ec2.Options)) (*ec2.DeleteFpgaImageOutput, error)
	DeleteInstanceEventWindowMock                                       func(arg1 context.Context, arg2 *ec2.DeleteInstanceEventWindowInput, arg3 ...func(*ec2.Options)) (*ec2.DeleteInstanceEventWindowOutput, error)
	DeleteInternetGatewayMock                                           func(arg1 context.Context, arg2 *ec2.DeleteInternetGatewayInput, arg3 ...func(*ec2.Options)) (*ec2.DeleteInternetGatewayOutput, error)
	DeleteIpamMock                                                      func(arg1 context.Context, arg2 *ec2.DeleteIpamInput, arg3 ...func(*ec2.Options)) (*ec2.DeleteIpamOutput, error)
	DeleteIpamPoolMock                                                  func(arg1 context.Context, arg2 *ec2.DeleteIpamPoolInput, arg3 ...func(*ec2.Options)) (*ec2.DeleteIpamPoolOutput, error)
	DeleteIpamScopeMock                                                 func(arg1 context.Context, arg2 *ec2.DeleteIpamScopeInput, arg3 ...func(*ec2.Options)) (*ec2.DeleteIpamScopeOutput, error)
	DeleteKeyPairMock                                                   func(arg1 context.Context, arg2 *ec2.DeleteKeyPairInput, arg3 ...func(*ec2.Options)) (*ec2.DeleteKeyPairOutput, error)
	DeleteLaunchTemplateMock                                            func(arg1 context.Context, arg2 *ec2.DeleteLaunchTemplateInput, arg3 ...func(*ec2.Options)) (*ec2.DeleteLaunchTemplateOutput, error)
	DeleteLaunchTemplateVersionsMock                                    func(arg1 context.Context, arg2 *ec2.DeleteLaunchTemplateVersionsInput, arg3 ...func(*ec2.Options)) (*ec2.DeleteLaunchTemplateVersionsOutput, error)
	DeleteLocalGatewayRouteMock                                         func(arg1 context.Context, arg2 *ec2.DeleteLocalGatewayRouteInput, arg3 ...func(*ec2.Options)) (*ec2.DeleteLocalGatewayRouteOutput, error)
	DeleteLocalGatewayRouteTableVpcAssociationMock                      func(arg1 context.Context, arg2 *ec2.DeleteLocalGatewayRouteTableVpcAssociationInput, arg3 ...func(*ec2.Options)) (*ec2.DeleteLocalGatewayRouteTableVpcAssociationOutput, error)
	DeleteManagedPrefixListMock                                         func(arg1 context.Context, arg2 *ec2.DeleteManagedPrefixListInput, arg3 ...func(*ec2.Options)) (*ec2.DeleteManagedPrefixListOutput, error)
	DeleteNatGatewayMock                                                func(arg1 context.Context, arg2 *ec2.DeleteNatGatewayInput, arg3 ...func(*ec2.Options)) (*ec2.DeleteNatGatewayOutput, error)
	DeleteNetworkAclMock                                                func(arg1 context.Context, arg2 *ec2.DeleteNetworkAclInput, arg3 ...func(*ec2.Options)) (*ec2.DeleteNetworkAclOutput, error)
	DeleteNetworkAclEntryMock                                           func(arg1 context.Context, arg2 *ec2.DeleteNetworkAclEntryInput, arg3 ...func(*ec2.Options)) (*ec2.DeleteNetworkAclEntryOutput, error)
	DeleteNetworkInsightsAccessScopeMock                                func(arg1 context.Context, arg2 *ec2.DeleteNetworkInsightsAccessScopeInput, arg3 ...func(*ec2.Options)) (*ec2.DeleteNetworkInsightsAccessScopeOutput, error)
	DeleteNetworkInsightsAccessScopeAnalysisMock                        func(arg1 context.Context, arg2 *ec2.DeleteNetworkInsightsAccessScopeAnalysisInput, arg3 ...func(*ec2.Options)) (*ec2.DeleteNetworkInsightsAccessScopeAnalysisOutput, error)
	DeleteNetworkInsightsAnalysisMock                                   func(arg1 context.Context, arg2 *ec2.DeleteNetworkInsightsAnalysisInput, arg3 ...func(*ec2.Options)) (*ec2.DeleteNetworkInsightsAnalysisOutput, error)
	DeleteNetworkInsightsPathMock                                       func(arg1 context.Context, arg2 *ec2.DeleteNetworkInsightsPathInput, arg3 ...func(*ec2.Options)) (*ec2.DeleteNetworkInsightsPathOutput, error)
	DeleteNetworkInterfaceMock                                          func(arg1 context.Context, arg2 *ec2.DeleteNetworkInterfaceInput, arg3 ...func(*ec2.Options)) (*ec2.DeleteNetworkInterfaceOutput, error)
	DeleteNetworkInterfacePermissionMock                                func(arg1 context.Context, arg2 *ec2.DeleteNetworkInterfacePermissionInput, arg3 ...func(*ec2.Options)) (*ec2.DeleteNetworkInterfacePermissionOutput, error)
	DeletePlacementGroupMock                                            func(arg1 context.Context, arg2 *ec2.DeletePlacementGroupInput, arg3 ...func(*ec2.Options)) (*ec2.DeletePlacementGroupOutput, error)
	DeletePublicIpv4PoolMock                                            func(arg1 context.Context, arg2 *ec2.DeletePublicIpv4PoolInput, arg3 ...func(*ec2.Options)) (*ec2.DeletePublicIpv4PoolOutput, error)
	DeleteQueuedReservedInstancesMock                                   func(arg1 context.Context, arg2 *ec2.DeleteQueuedReservedInstancesInput, arg3 ...func(*ec2.Options)) (*ec2.DeleteQueuedReservedInstancesOutput, error)
	DeleteRouteMock                                                     func(arg1 context.Context, arg2 *ec2.DeleteRouteInput, arg3 ...func(*ec2.Options)) (*ec2.DeleteRouteOutput, error)
	DeleteRouteTableMock                                                func(arg1 context.Context, arg2 *ec2.DeleteRouteTableInput, arg3 ...func(*ec2.Options)) (*ec2.DeleteRouteTableOutput, error)
	DeleteSecurityGroupMock                                             func(arg1 context.Context, arg2 *ec2.DeleteSecurityGroupInput, arg3 ...func(*ec2.Options)) (*ec2.DeleteSecurityGroupOutput, error)
	DeleteSnapshotMock                                                  func(arg1 context.Context, arg2 *ec2.DeleteSnapshotInput, arg3 ...func(*ec2.Options)) (*ec2.DeleteSnapshotOutput, error)
	DeleteSpotDatafeedSubscriptionMock                                  func(arg1 context.Context, arg2 *ec2.DeleteSpotDatafeedSubscriptionInput, arg3 ...func(*ec2.Options)) (*ec2.DeleteSpotDatafeedSubscriptionOutput, error)
	DeleteSubnetMock                                                    func(arg1 context.Context, arg2 *ec2.DeleteSubnetInput, arg3 ...func(*ec2.Options)) (*ec2.DeleteSubnetOutput, error)
	DeleteSubnetCidrReservationMock                                     func(arg1 context.Context, arg2 *ec2.DeleteSubnetCidrReservationInput, arg3 ...func(*ec2.Options)) (*ec2.DeleteSubnetCidrReservationOutput, error)
	DeleteTagsMock                                                      func(arg1 context.Context, arg2 *ec2.DeleteTagsInput, arg3 ...func(*ec2.Options)) (*ec2.DeleteTagsOutput, error)
	DeleteTrafficMirrorFilterMock                                       func(arg1 context.Context, arg2 *ec2.DeleteTrafficMirrorFilterInput, arg3 ...func(*ec2.Options)) (*ec2.DeleteTrafficMirrorFilterOutput, error)
	DeleteTrafficMirrorFilterRuleMock                                   func(arg1 context.Context, arg2 *ec2.DeleteTrafficMirrorFilterRuleInput, arg3 ...func(*ec2.Options)) (*ec2.DeleteTrafficMirrorFilterRuleOutput, error)
	DeleteTrafficMirrorSessionMock                                      func(arg1 context.Context, arg2 *ec2.DeleteTrafficMirrorSessionInput, arg3 ...func(*ec2.Options)) (*ec2.DeleteTrafficMirrorSessionOutput, error)
	DeleteTrafficMirrorTargetMock                                       func(arg1 context.Context, arg2 *ec2.DeleteTrafficMirrorTargetInput, arg3 ...func(*ec2.Options)) (*ec2.DeleteTrafficMirrorTargetOutput, error)
	DeleteTransitGatewayMock                                            func(arg1 context.Context, arg2 *ec2.DeleteTransitGatewayInput, arg3 ...func(*ec2.Options)) (*ec2.DeleteTransitGatewayOutput, error)
	DeleteTransitGatewayConnectMock                                     func(arg1 context.Context, arg2 *ec2.DeleteTransitGatewayConnectInput, arg3 ...func(*ec2.Options)) (*ec2.DeleteTransitGatewayConnectOutput, error)
	DeleteTransitGatewayConnectPeerMock                                 func(arg1 context.Context, arg2 *ec2.DeleteTransitGatewayConnectPeerInput, arg3 ...func(*ec2.Options)) (*ec2.DeleteTransitGatewayConnectPeerOutput, error)
	DeleteTransitGatewayMulticastDomainMock                             func(arg1 context.Context, arg2 *ec2.DeleteTransitGatewayMulticastDomainInput, arg3 ...func(*ec2.Options)) (*ec2.DeleteTransitGatewayMulticastDomainOutput, error)
	DeleteTransitGatewayPeeringAttachmentMock                           func(arg1 context.Context, arg2 *ec2.DeleteTransitGatewayPeeringAttachmentInput, arg3 ...func(*ec2.Options)) (*ec2.DeleteTransitGatewayPeeringAttachmentOutput, error)
	DeleteTransitGatewayPolicyTableMock                                 func(arg1 context.Context, arg2 *ec2.DeleteTransitGatewayPolicyTableInput, arg3 ...func(*ec2.Options)) (*ec2.DeleteTransitGatewayPolicyTableOutput, error)
	DeleteTransitGatewayPrefixListReferenceMock                         func(arg1 context.Context, arg2 *ec2.DeleteTransitGatewayPrefixListReferenceInput, arg3 ...func(*ec2.Options)) (*ec2.DeleteTransitGatewayPrefixListReferenceOutput, error)
	DeleteTransitGatewayRouteMock                                       func(arg1 context.Context, arg2 *ec2.DeleteTransitGatewayRouteInput, arg3 ...func(*ec2.Options)) (*ec2.DeleteTransitGatewayRouteOutput, error)
	DeleteTransitGatewayRouteTableMock                                  func(arg1 context.Context, arg2 *ec2.DeleteTransitGatewayRouteTableInput, arg3 ...func(*ec2.Options)) (*ec2.DeleteTransitGatewayRouteTableOutput, error)
	DeleteTransitGatewayRouteTableAnnouncementMock                      func(arg1 context.Context, arg2 *ec2.DeleteTransitGatewayRouteTableAnnouncementInput, arg3 ...func(*ec2.Options)) (*ec2.DeleteTransitGatewayRouteTableAnnouncementOutput, error)
	DeleteTransitGatewayVpcAttachmentMock                               func(arg1 context.Context, arg2 *ec2.DeleteTransitGatewayVpcAttachmentInput, arg3 ...func(*ec2.Options)) (*ec2.DeleteTransitGatewayVpcAttachmentOutput, error)
	DeleteVolumeMock                                                    func(arg1 context.Context, arg2 *ec2.DeleteVolumeInput, arg3 ...func(*ec2.Options)) (*ec2.DeleteVolumeOutput, error)
	DeleteVpcMock                                                       func(arg1 context.Context, arg2 *ec2.DeleteVpcInput, arg3 ...func(*ec2.Options)) (*ec2.DeleteVpcOutput, error)
	DeleteVpcEndpointConnectionNotificationsMock                        func(arg1 context.Context, arg2 *ec2.DeleteVpcEndpointConnectionNotificationsInput, arg3 ...func(*ec2.Options)) (*ec2.DeleteVpcEndpointConnectionNotificationsOutput, error)
	DeleteVpcEndpointServiceConfigurationsMock                          func(arg1 context.Context, arg2 *ec2.DeleteVpcEndpointServiceConfigurationsInput, arg3 ...func(*ec2.Options)) (*ec2.DeleteVpcEndpointServiceConfigurationsOutput, error)
	DeleteVpcEndpointsMock                                              func(arg1 context.Context, arg2 *ec2.DeleteVpcEndpointsInput, arg3 ...func(*ec2.Options)) (*ec2.DeleteVpcEndpointsOutput, error)
	DeleteVpcPeeringConnectionMock                                      func(arg1 context.Context, arg2 *ec2.DeleteVpcPeeringConnectionInput, arg3 ...func(*ec2.Options)) (*ec2.DeleteVpcPeeringConnectionOutput, error)
	DeleteVpnConnectionMock                                             func(arg1 context.Context, arg2 *ec2.DeleteVpnConnectionInput, arg3 ...func(*ec2.Options)) (*ec2.DeleteVpnConnectionOutput, error)
	DeleteVpnConnectionRouteMock                                        func(arg1 context.Context, arg2 *ec2.DeleteVpnConnectionRouteInput, arg3 ...func(*ec2.Options)) (*ec2.DeleteVpnConnectionRouteOutput, error)
	DeleteVpnGatewayMock                                                func(arg1 context.Context, arg2 *ec2.DeleteVpnGatewayInput, arg3 ...func(*ec2.Options)) (*ec2.DeleteVpnGatewayOutput, error)
	DeprovisionByoipCidrMock                                            func(arg1 context.Context, arg2 *ec2.DeprovisionByoipCidrInput, arg3 ...func(*ec2.Options)) (*ec2.DeprovisionByoipCidrOutput, error)
	DeprovisionIpamPoolCidrMock                                         func(arg1 context.Context, arg2 *ec2.DeprovisionIpamPoolCidrInput, arg3 ...func(*ec2.Options)) (*ec2.DeprovisionIpamPoolCidrOutput, error)
	DeprovisionPublicIpv4PoolCidrMock                                   func(arg1 context.Context, arg2 *ec2.DeprovisionPublicIpv4PoolCidrInput, arg3 ...func(*ec2.Options)) (*ec2.DeprovisionPublicIpv4PoolCidrOutput, error)
	DeregisterImageMock                                                 func(arg1 context.Context, arg2 *ec2.DeregisterImageInput, arg3 ...func(*ec2.Options)) (*ec2.DeregisterImageOutput, error)
	DeregisterInstanceEventNotificationAttributesMock                   func(arg1 context.Context, arg2 *ec2.DeregisterInstanceEventNotificationAttributesInput, arg3 ...func(*ec2.Options)) (*ec2.DeregisterInstanceEventNotificationAttributesOutput, error)
	DeregisterTransitGatewayMulticastGroupMembersMock                   func(arg1 context.Context, arg2 *ec2.DeregisterTransitGatewayMulticastGroupMembersInput, arg3 ...func(*ec2.Options)) (*ec2.DeregisterTransitGatewayMulticastGroupMembersOutput, error)
	DeregisterTransitGatewayMulticastGroupSourcesMock                   func(arg1 context.Context, arg2 *ec2.DeregisterTransitGatewayMulticastGroupSourcesInput, arg3 ...func(*ec2.Options)) (*ec2.DeregisterTransitGatewayMulticastGroupSourcesOutput, error)
	DescribeAccountAttributesMock                                       func(arg1 context.Context, arg2 *ec2.DescribeAccountAttributesInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeAccountAttributesOutput, error)
	DescribeAddressesMock                                               func(arg1 context.Context, arg2 *ec2.DescribeAddressesInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeAddressesOutput, error)
	DescribeAddressesAttributeMock                                      func(arg1 context.Context, arg2 *ec2.DescribeAddressesAttributeInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeAddressesAttributeOutput, error)
	DescribeAggregateIdFormatMock                                       func(arg1 context.Context, arg2 *ec2.DescribeAggregateIdFormatInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeAggregateIdFormatOutput, error)
	DescribeAvailabilityZonesMock                                       func(arg1 context.Context, arg2 *ec2.DescribeAvailabilityZonesInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeAvailabilityZonesOutput, error)
	DescribeBundleTasksMock                                             func(arg1 context.Context, arg2 *ec2.DescribeBundleTasksInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeBundleTasksOutput, error)
	DescribeByoipCidrsMock                                              func(arg1 context.Context, arg2 *ec2.DescribeByoipCidrsInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeByoipCidrsOutput, error)
	DescribeCapacityReservationFleetsMock                               func(arg1 context.Context, arg2 *ec2.DescribeCapacityReservationFleetsInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeCapacityReservationFleetsOutput, error)
	DescribeCapacityReservationsMock                                    func(arg1 context.Context, arg2 *ec2.DescribeCapacityReservationsInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeCapacityReservationsOutput, error)
	DescribeCarrierGatewaysMock                                         func(arg1 context.Context, arg2 *ec2.DescribeCarrierGatewaysInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeCarrierGatewaysOutput, error)
	DescribeClassicLinkInstancesMock                                    func(arg1 context.Context, arg2 *ec2.DescribeClassicLinkInstancesInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeClassicLinkInstancesOutput, error)
	DescribeClientVpnAuthorizationRulesMock                             func(arg1 context.Context, arg2 *ec2.DescribeClientVpnAuthorizationRulesInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeClientVpnAuthorizationRulesOutput, error)
	DescribeClientVpnConnectionsMock                                    func(arg1 context.Context, arg2 *ec2.DescribeClientVpnConnectionsInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeClientVpnConnectionsOutput, error)
	DescribeClientVpnEndpointsMock                                      func(arg1 context.Context, arg2 *ec2.DescribeClientVpnEndpointsInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeClientVpnEndpointsOutput, error)
	DescribeClientVpnRoutesMock                                         func(arg1 context.Context, arg2 *ec2.DescribeClientVpnRoutesInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeClientVpnRoutesOutput, error)
	DescribeClientVpnTargetNetworksMock                                 func(arg1 context.Context, arg2 *ec2.DescribeClientVpnTargetNetworksInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeClientVpnTargetNetworksOutput, error)
	DescribeCoipPoolsMock                                               func(arg1 context.Context, arg2 *ec2.DescribeCoipPoolsInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeCoipPoolsOutput, error)
	DescribeConversionTasksMock                                         func(arg1 context.Context, arg2 *ec2.DescribeConversionTasksInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeConversionTasksOutput, error)
	DescribeCustomerGatewaysMock                                        func(arg1 context.Context, arg2 *ec2.DescribeCustomerGatewaysInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeCustomerGatewaysOutput, error)
	DescribeDhcpOptionsMock                                             func(arg1 context.Context, arg2 *ec2.DescribeDhcpOptionsInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeDhcpOptionsOutput, error)
	DescribeEgressOnlyInternetGatewaysMock                              func(arg1 context.Context, arg2 *ec2.DescribeEgressOnlyInternetGatewaysInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeEgressOnlyInternetGatewaysOutput, error)
	DescribeElasticGpusMock                                             func(arg1 context.Context, arg2 *ec2.DescribeElasticGpusInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeElasticGpusOutput, error)
	DescribeExportImageTasksMock                                        func(arg1 context.Context, arg2 *ec2.DescribeExportImageTasksInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeExportImageTasksOutput, error)
	DescribeExportTasksMock                                             func(arg1 context.Context, arg2 *ec2.DescribeExportTasksInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeExportTasksOutput, error)
	DescribeFastLaunchImagesMock                                        func(arg1 context.Context, arg2 *ec2.DescribeFastLaunchImagesInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeFastLaunchImagesOutput, error)
	DescribeFastSnapshotRestoresMock                                    func(arg1 context.Context, arg2 *ec2.DescribeFastSnapshotRestoresInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeFastSnapshotRestoresOutput, error)
	DescribeFleetHistoryMock                                            func(arg1 context.Context, arg2 *ec2.DescribeFleetHistoryInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeFleetHistoryOutput, error)
	DescribeFleetInstancesMock                                          func(arg1 context.Context, arg2 *ec2.DescribeFleetInstancesInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeFleetInstancesOutput, error)
	DescribeFleetsMock                                                  func(arg1 context.Context, arg2 *ec2.DescribeFleetsInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeFleetsOutput, error)
	DescribeFlowLogsMock                                                func(arg1 context.Context, arg2 *ec2.DescribeFlowLogsInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeFlowLogsOutput, error)
	DescribeFpgaImageAttributeMock                                      func(arg1 context.Context, arg2 *ec2.DescribeFpgaImageAttributeInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeFpgaImageAttributeOutput, error)
	DescribeFpgaImagesMock                                              func(arg1 context.Context, arg2 *ec2.DescribeFpgaImagesInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeFpgaImagesOutput, error)
	DescribeHostReservationOfferingsMock                                func(arg1 context.Context, arg2 *ec2.DescribeHostReservationOfferingsInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeHostReservationOfferingsOutput, error)
	DescribeHostReservationsMock                                        func(arg1 context.Context, arg2 *ec2.DescribeHostReservationsInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeHostReservationsOutput, error)
	DescribeHostsMock                                                   func(arg1 context.Context, arg2 *ec2.DescribeHostsInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeHostsOutput, error)
	DescribeIamInstanceProfileAssociationsMock                          func(arg1 context.Context, arg2 *ec2.DescribeIamInstanceProfileAssociationsInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeIamInstanceProfileAssociationsOutput, error)
	DescribeIdFormatMock                                                func(arg1 context.Context, arg2 *ec2.DescribeIdFormatInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeIdFormatOutput, error)
	DescribeIdentityIdFormatMock                                        func(arg1 context.Context, arg2 *ec2.DescribeIdentityIdFormatInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeIdentityIdFormatOutput, error)
	DescribeImageAttributeMock                                          func(arg1 context.Context, arg2 *ec2.DescribeImageAttributeInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeImageAttributeOutput, error)
	DescribeImagesMock                                                  func(arg1 context.Context, arg2 *ec2.DescribeImagesInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeImagesOutput, error)
	DescribeImportImageTasksMock                                        func(arg1 context.Context, arg2 *ec2.DescribeImportImageTasksInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeImportImageTasksOutput, error)
	DescribeImportSnapshotTasksMock                                     func(arg1 context.Context, arg2 *ec2.DescribeImportSnapshotTasksInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeImportSnapshotTasksOutput, error)
	DescribeInstanceAttributeMock                                       func(arg1 context.Context, arg2 *ec2.DescribeInstanceAttributeInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeInstanceAttributeOutput, error)
	DescribeInstanceCreditSpecificationsMock                            func(arg1 context.Context, arg2 *ec2.DescribeInstanceCreditSpecificationsInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeInstanceCreditSpecificationsOutput, error)
	DescribeInstanceEventNotificationAttributesMock                     func(arg1 context.Context, arg2 *ec2.DescribeInstanceEventNotificationAttributesInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeInstanceEventNotificationAttributesOutput, error)
	DescribeInstanceEventWindowsMock                                    func(arg1 context.Context, arg2 *ec2.DescribeInstanceEventWindowsInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeInstanceEventWindowsOutput, error)
	DescribeInstanceStatusMock                                          func(arg1 context.Context, arg2 *ec2.DescribeInstanceStatusInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeInstanceStatusOutput, error)
	DescribeInstanceTypeOfferingsMock                                   func(arg1 context.Context, arg2 *ec2.DescribeInstanceTypeOfferingsInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeInstanceTypeOfferingsOutput, error)
	DescribeInstanceTypesMock                                           func(arg1 context.Context, arg2 *ec2.DescribeInstanceTypesInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeInstanceTypesOutput, error)
	DescribeInstancesMock                                               func(arg1 context.Context, arg2 *ec2.DescribeInstancesInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeInstancesOutput, error)
	DescribeInternetGatewaysMock                                        func(arg1 context.Context, arg2 *ec2.DescribeInternetGatewaysInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeInternetGatewaysOutput, error)
	DescribeIpamPoolsMock                                               func(arg1 context.Context, arg2 *ec2.DescribeIpamPoolsInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeIpamPoolsOutput, error)
	DescribeIpamScopesMock                                              func(arg1 context.Context, arg2 *ec2.DescribeIpamScopesInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeIpamScopesOutput, error)
	DescribeIpamsMock                                                   func(arg1 context.Context, arg2 *ec2.DescribeIpamsInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeIpamsOutput, error)
	DescribeIpv6PoolsMock                                               func(arg1 context.Context, arg2 *ec2.DescribeIpv6PoolsInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeIpv6PoolsOutput, error)
	DescribeKeyPairsMock                                                func(arg1 context.Context, arg2 *ec2.DescribeKeyPairsInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeKeyPairsOutput, error)
	DescribeLaunchTemplateVersionsMock                                  func(arg1 context.Context, arg2 *ec2.DescribeLaunchTemplateVersionsInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeLaunchTemplateVersionsOutput, error)
	DescribeLaunchTemplatesMock                                         func(arg1 context.Context, arg2 *ec2.DescribeLaunchTemplatesInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeLaunchTemplatesOutput, error)
	DescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociationsMock func(arg1 context.Context, arg2 *ec2.DescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociationsInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociationsOutput, error)
	DescribeLocalGatewayRouteTableVpcAssociationsMock                   func(arg1 context.Context, arg2 *ec2.DescribeLocalGatewayRouteTableVpcAssociationsInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeLocalGatewayRouteTableVpcAssociationsOutput, error)
	DescribeLocalGatewayRouteTablesMock                                 func(arg1 context.Context, arg2 *ec2.DescribeLocalGatewayRouteTablesInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeLocalGatewayRouteTablesOutput, error)
	DescribeLocalGatewayVirtualInterfaceGroupsMock                      func(arg1 context.Context, arg2 *ec2.DescribeLocalGatewayVirtualInterfaceGroupsInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeLocalGatewayVirtualInterfaceGroupsOutput, error)
	DescribeLocalGatewayVirtualInterfacesMock                           func(arg1 context.Context, arg2 *ec2.DescribeLocalGatewayVirtualInterfacesInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeLocalGatewayVirtualInterfacesOutput, error)
	DescribeLocalGatewaysMock                                           func(arg1 context.Context, arg2 *ec2.DescribeLocalGatewaysInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeLocalGatewaysOutput, error)
	DescribeManagedPrefixListsMock                                      func(arg1 context.Context, arg2 *ec2.DescribeManagedPrefixListsInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeManagedPrefixListsOutput, error)
	DescribeMovingAddressesMock                                         func(arg1 context.Context, arg2 *ec2.DescribeMovingAddressesInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeMovingAddressesOutput, error)
	DescribeNatGatewaysMock                                             func(arg1 context.Context, arg2 *ec2.DescribeNatGatewaysInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeNatGatewaysOutput, error)
	DescribeNetworkAclsMock                                             func(arg1 context.Context, arg2 *ec2.DescribeNetworkAclsInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeNetworkAclsOutput, error)
	DescribeNetworkInsightsAccessScopeAnalysesMock                      func(arg1 context.Context, arg2 *ec2.DescribeNetworkInsightsAccessScopeAnalysesInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeNetworkInsightsAccessScopeAnalysesOutput, error)
	DescribeNetworkInsightsAccessScopesMock                             func(arg1 context.Context, arg2 *ec2.DescribeNetworkInsightsAccessScopesInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeNetworkInsightsAccessScopesOutput, error)
	DescribeNetworkInsightsAnalysesMock                                 func(arg1 context.Context, arg2 *ec2.DescribeNetworkInsightsAnalysesInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeNetworkInsightsAnalysesOutput, error)
	DescribeNetworkInsightsPathsMock                                    func(arg1 context.Context, arg2 *ec2.DescribeNetworkInsightsPathsInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeNetworkInsightsPathsOutput, error)
	DescribeNetworkInterfaceAttributeMock                               func(arg1 context.Context, arg2 *ec2.DescribeNetworkInterfaceAttributeInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeNetworkInterfaceAttributeOutput, error)
	DescribeNetworkInterfacePermissionsMock                             func(arg1 context.Context, arg2 *ec2.DescribeNetworkInterfacePermissionsInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeNetworkInterfacePermissionsOutput, error)
	DescribeNetworkInterfacesMock                                       func(arg1 context.Context, arg2 *ec2.DescribeNetworkInterfacesInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeNetworkInterfacesOutput, error)
	DescribePlacementGroupsMock                                         func(arg1 context.Context, arg2 *ec2.DescribePlacementGroupsInput, arg3 ...func(*ec2.Options)) (*ec2.DescribePlacementGroupsOutput, error)
	DescribePrefixListsMock                                             func(arg1 context.Context, arg2 *ec2.DescribePrefixListsInput, arg3 ...func(*ec2.Options)) (*ec2.DescribePrefixListsOutput, error)
	DescribePrincipalIdFormatMock                                       func(arg1 context.Context, arg2 *ec2.DescribePrincipalIdFormatInput, arg3 ...func(*ec2.Options)) (*ec2.DescribePrincipalIdFormatOutput, error)
	DescribePublicIpv4PoolsMock                                         func(arg1 context.Context, arg2 *ec2.DescribePublicIpv4PoolsInput, arg3 ...func(*ec2.Options)) (*ec2.DescribePublicIpv4PoolsOutput, error)
	DescribeRegionsMock                                                 func(arg1 context.Context, arg2 *ec2.DescribeRegionsInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeRegionsOutput, error)
	DescribeReplaceRootVolumeTasksMock                                  func(arg1 context.Context, arg2 *ec2.DescribeReplaceRootVolumeTasksInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeReplaceRootVolumeTasksOutput, error)
	DescribeReservedInstancesMock                                       func(arg1 context.Context, arg2 *ec2.DescribeReservedInstancesInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeReservedInstancesOutput, error)
	DescribeReservedInstancesListingsMock                               func(arg1 context.Context, arg2 *ec2.DescribeReservedInstancesListingsInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeReservedInstancesListingsOutput, error)
	DescribeReservedInstancesModificationsMock                          func(arg1 context.Context, arg2 *ec2.DescribeReservedInstancesModificationsInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeReservedInstancesModificationsOutput, error)
	DescribeReservedInstancesOfferingsMock                              func(arg1 context.Context, arg2 *ec2.DescribeReservedInstancesOfferingsInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeReservedInstancesOfferingsOutput, error)
	DescribeRouteTablesMock                                             func(arg1 context.Context, arg2 *ec2.DescribeRouteTablesInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeRouteTablesOutput, error)
	DescribeScheduledInstanceAvailabilityMock                           func(arg1 context.Context, arg2 *ec2.DescribeScheduledInstanceAvailabilityInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeScheduledInstanceAvailabilityOutput, error)
	DescribeScheduledInstancesMock                                      func(arg1 context.Context, arg2 *ec2.DescribeScheduledInstancesInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeScheduledInstancesOutput, error)
	DescribeSecurityGroupReferencesMock                                 func(arg1 context.Context, arg2 *ec2.DescribeSecurityGroupReferencesInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeSecurityGroupReferencesOutput, error)
	DescribeSecurityGroupRulesMock                                      func(arg1 context.Context, arg2 *ec2.DescribeSecurityGroupRulesInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeSecurityGroupRulesOutput, error)
	DescribeSecurityGroupsMock                                          func(arg1 context.Context, arg2 *ec2.DescribeSecurityGroupsInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeSecurityGroupsOutput, error)
	DescribeSnapshotAttributeMock                                       func(arg1 context.Context, arg2 *ec2.DescribeSnapshotAttributeInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeSnapshotAttributeOutput, error)
	DescribeSnapshotTierStatusMock                                      func(arg1 context.Context, arg2 *ec2.DescribeSnapshotTierStatusInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeSnapshotTierStatusOutput, error)
	DescribeSnapshotsMock                                               func(arg1 context.Context, arg2 *ec2.DescribeSnapshotsInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeSnapshotsOutput, error)
	DescribeSpotDatafeedSubscriptionMock                                func(arg1 context.Context, arg2 *ec2.DescribeSpotDatafeedSubscriptionInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeSpotDatafeedSubscriptionOutput, error)
	DescribeSpotFleetInstancesMock                                      func(arg1 context.Context, arg2 *ec2.DescribeSpotFleetInstancesInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeSpotFleetInstancesOutput, error)
	DescribeSpotFleetRequestHistoryMock                                 func(arg1 context.Context, arg2 *ec2.DescribeSpotFleetRequestHistoryInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeSpotFleetRequestHistoryOutput, error)
	DescribeSpotFleetRequestsMock                                       func(arg1 context.Context, arg2 *ec2.DescribeSpotFleetRequestsInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeSpotFleetRequestsOutput, error)
	DescribeSpotInstanceRequestsMock                                    func(arg1 context.Context, arg2 *ec2.DescribeSpotInstanceRequestsInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeSpotInstanceRequestsOutput, error)
	DescribeSpotPriceHistoryMock                                        func(arg1 context.Context, arg2 *ec2.DescribeSpotPriceHistoryInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeSpotPriceHistoryOutput, error)
	DescribeStaleSecurityGroupsMock                                     func(arg1 context.Context, arg2 *ec2.DescribeStaleSecurityGroupsInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeStaleSecurityGroupsOutput, error)
	DescribeStoreImageTasksMock                                         func(arg1 context.Context, arg2 *ec2.DescribeStoreImageTasksInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeStoreImageTasksOutput, error)
	DescribeSubnetsMock                                                 func(arg1 context.Context, arg2 *ec2.DescribeSubnetsInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeSubnetsOutput, error)
	DescribeTagsMock                                                    func(arg1 context.Context, arg2 *ec2.DescribeTagsInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeTagsOutput, error)
	DescribeTrafficMirrorFiltersMock                                    func(arg1 context.Context, arg2 *ec2.DescribeTrafficMirrorFiltersInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeTrafficMirrorFiltersOutput, error)
	DescribeTrafficMirrorSessionsMock                                   func(arg1 context.Context, arg2 *ec2.DescribeTrafficMirrorSessionsInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeTrafficMirrorSessionsOutput, error)
	DescribeTrafficMirrorTargetsMock                                    func(arg1 context.Context, arg2 *ec2.DescribeTrafficMirrorTargetsInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeTrafficMirrorTargetsOutput, error)
	DescribeTransitGatewayAttachmentsMock                               func(arg1 context.Context, arg2 *ec2.DescribeTransitGatewayAttachmentsInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeTransitGatewayAttachmentsOutput, error)
	DescribeTransitGatewayConnectPeersMock                              func(arg1 context.Context, arg2 *ec2.DescribeTransitGatewayConnectPeersInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeTransitGatewayConnectPeersOutput, error)
	DescribeTransitGatewayConnectsMock                                  func(arg1 context.Context, arg2 *ec2.DescribeTransitGatewayConnectsInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeTransitGatewayConnectsOutput, error)
	DescribeTransitGatewayMulticastDomainsMock                          func(arg1 context.Context, arg2 *ec2.DescribeTransitGatewayMulticastDomainsInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeTransitGatewayMulticastDomainsOutput, error)
	DescribeTransitGatewayPeeringAttachmentsMock                        func(arg1 context.Context, arg2 *ec2.DescribeTransitGatewayPeeringAttachmentsInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeTransitGatewayPeeringAttachmentsOutput, error)
	DescribeTransitGatewayPolicyTablesMock                              func(arg1 context.Context, arg2 *ec2.DescribeTransitGatewayPolicyTablesInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeTransitGatewayPolicyTablesOutput, error)
	DescribeTransitGatewayRouteTableAnnouncementsMock                   func(arg1 context.Context, arg2 *ec2.DescribeTransitGatewayRouteTableAnnouncementsInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeTransitGatewayRouteTableAnnouncementsOutput, error)
	DescribeTransitGatewayRouteTablesMock                               func(arg1 context.Context, arg2 *ec2.DescribeTransitGatewayRouteTablesInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeTransitGatewayRouteTablesOutput, error)
	DescribeTransitGatewayVpcAttachmentsMock                            func(arg1 context.Context, arg2 *ec2.DescribeTransitGatewayVpcAttachmentsInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeTransitGatewayVpcAttachmentsOutput, error)
	DescribeTransitGatewaysMock                                         func(arg1 context.Context, arg2 *ec2.DescribeTransitGatewaysInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeTransitGatewaysOutput, error)
	DescribeTrunkInterfaceAssociationsMock                              func(arg1 context.Context, arg2 *ec2.DescribeTrunkInterfaceAssociationsInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeTrunkInterfaceAssociationsOutput, error)
	DescribeVolumeAttributeMock                                         func(arg1 context.Context, arg2 *ec2.DescribeVolumeAttributeInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeVolumeAttributeOutput, error)
	DescribeVolumeStatusMock                                            func(arg1 context.Context, arg2 *ec2.DescribeVolumeStatusInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeVolumeStatusOutput, error)
	DescribeVolumesMock                                                 func(arg1 context.Context, arg2 *ec2.DescribeVolumesInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeVolumesOutput, error)
	DescribeVolumesModificationsMock                                    func(arg1 context.Context, arg2 *ec2.DescribeVolumesModificationsInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeVolumesModificationsOutput, error)
	DescribeVpcAttributeMock                                            func(arg1 context.Context, arg2 *ec2.DescribeVpcAttributeInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeVpcAttributeOutput, error)
	DescribeVpcClassicLinkMock                                          func(arg1 context.Context, arg2 *ec2.DescribeVpcClassicLinkInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeVpcClassicLinkOutput, error)
	DescribeVpcClassicLinkDnsSupportMock                                func(arg1 context.Context, arg2 *ec2.DescribeVpcClassicLinkDnsSupportInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeVpcClassicLinkDnsSupportOutput, error)
	DescribeVpcEndpointConnectionNotificationsMock                      func(arg1 context.Context, arg2 *ec2.DescribeVpcEndpointConnectionNotificationsInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeVpcEndpointConnectionNotificationsOutput, error)
	DescribeVpcEndpointConnectionsMock                                  func(arg1 context.Context, arg2 *ec2.DescribeVpcEndpointConnectionsInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeVpcEndpointConnectionsOutput, error)
	DescribeVpcEndpointServiceConfigurationsMock                        func(arg1 context.Context, arg2 *ec2.DescribeVpcEndpointServiceConfigurationsInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeVpcEndpointServiceConfigurationsOutput, error)
	DescribeVpcEndpointServicePermissionsMock                           func(arg1 context.Context, arg2 *ec2.DescribeVpcEndpointServicePermissionsInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeVpcEndpointServicePermissionsOutput, error)
	DescribeVpcEndpointServicesMock                                     func(arg1 context.Context, arg2 *ec2.DescribeVpcEndpointServicesInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeVpcEndpointServicesOutput, error)
	DescribeVpcEndpointsMock                                            func(arg1 context.Context, arg2 *ec2.DescribeVpcEndpointsInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeVpcEndpointsOutput, error)
	DescribeVpcPeeringConnectionsMock                                   func(arg1 context.Context, arg2 *ec2.DescribeVpcPeeringConnectionsInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeVpcPeeringConnectionsOutput, error)
	DescribeVpcsMock                                                    func(arg1 context.Context, arg2 *ec2.DescribeVpcsInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeVpcsOutput, error)
	DescribeVpnConnectionsMock                                          func(arg1 context.Context, arg2 *ec2.DescribeVpnConnectionsInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeVpnConnectionsOutput, error)
	DescribeVpnGatewaysMock                                             func(arg1 context.Context, arg2 *ec2.DescribeVpnGatewaysInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeVpnGatewaysOutput, error)
	DetachClassicLinkVpcMock                                            func(arg1 context.Context, arg2 *ec2.DetachClassicLinkVpcInput, arg3 ...func(*ec2.Options)) (*ec2.DetachClassicLinkVpcOutput, error)
	DetachInternetGatewayMock                                           func(arg1 context.Context, arg2 *ec2.DetachInternetGatewayInput, arg3 ...func(*ec2.Options)) (*ec2.DetachInternetGatewayOutput, error)
	DetachNetworkInterfaceMock                                          func(arg1 context.Context, arg2 *ec2.DetachNetworkInterfaceInput, arg3 ...func(*ec2.Options)) (*ec2.DetachNetworkInterfaceOutput, error)
	DetachVolumeMock                                                    func(arg1 context.Context, arg2 *ec2.DetachVolumeInput, arg3 ...func(*ec2.Options)) (*ec2.DetachVolumeOutput, error)
	DetachVpnGatewayMock                                                func(arg1 context.Context, arg2 *ec2.DetachVpnGatewayInput, arg3 ...func(*ec2.Options)) (*ec2.DetachVpnGatewayOutput, error)
	DisableEbsEncryptionByDefaultMock                                   func(arg1 context.Context, arg2 *ec2.DisableEbsEncryptionByDefaultInput, arg3 ...func(*ec2.Options)) (*ec2.DisableEbsEncryptionByDefaultOutput, error)
	DisableFastLaunchMock                                               func(arg1 context.Context, arg2 *ec2.DisableFastLaunchInput, arg3 ...func(*ec2.Options)) (*ec2.DisableFastLaunchOutput, error)
	DisableFastSnapshotRestoresMock                                     func(arg1 context.Context, arg2 *ec2.DisableFastSnapshotRestoresInput, arg3 ...func(*ec2.Options)) (*ec2.DisableFastSnapshotRestoresOutput, error)
	DisableImageDeprecationMock                                         func(arg1 context.Context, arg2 *ec2.DisableImageDeprecationInput, arg3 ...func(*ec2.Options)) (*ec2.DisableImageDeprecationOutput, error)
	DisableIpamOrganizationAdminAccountMock                             func(arg1 context.Context, arg2 *ec2.DisableIpamOrganizationAdminAccountInput, arg3 ...func(*ec2.Options)) (*ec2.DisableIpamOrganizationAdminAccountOutput, error)
	DisableSerialConsoleAccessMock                                      func(arg1 context.Context, arg2 *ec2.DisableSerialConsoleAccessInput, arg3 ...func(*ec2.Options)) (*ec2.DisableSerialConsoleAccessOutput, error)
	DisableTransitGatewayRouteTablePropagationMock                      func(arg1 context.Context, arg2 *ec2.DisableTransitGatewayRouteTablePropagationInput, arg3 ...func(*ec2.Options)) (*ec2.DisableTransitGatewayRouteTablePropagationOutput, error)
	DisableVgwRoutePropagationMock                                      func(arg1 context.Context, arg2 *ec2.DisableVgwRoutePropagationInput, arg3 ...func(*ec2.Options)) (*ec2.DisableVgwRoutePropagationOutput, error)
	DisableVpcClassicLinkMock                                           func(arg1 context.Context, arg2 *ec2.DisableVpcClassicLinkInput, arg3 ...func(*ec2.Options)) (*ec2.DisableVpcClassicLinkOutput, error)
	DisableVpcClassicLinkDnsSupportMock                                 func(arg1 context.Context, arg2 *ec2.DisableVpcClassicLinkDnsSupportInput, arg3 ...func(*ec2.Options)) (*ec2.DisableVpcClassicLinkDnsSupportOutput, error)
	DisassociateAddressMock                                             func(arg1 context.Context, arg2 *ec2.DisassociateAddressInput, arg3 ...func(*ec2.Options)) (*ec2.DisassociateAddressOutput, error)
	DisassociateClientVpnTargetNetworkMock                              func(arg1 context.Context, arg2 *ec2.DisassociateClientVpnTargetNetworkInput, arg3 ...func(*ec2.Options)) (*ec2.DisassociateClientVpnTargetNetworkOutput, error)
	DisassociateEnclaveCertificateIamRoleMock                           func(arg1 context.Context, arg2 *ec2.DisassociateEnclaveCertificateIamRoleInput, arg3 ...func(*ec2.Options)) (*ec2.DisassociateEnclaveCertificateIamRoleOutput, error)
	DisassociateIamInstanceProfileMock                                  func(arg1 context.Context, arg2 *ec2.DisassociateIamInstanceProfileInput, arg3 ...func(*ec2.Options)) (*ec2.DisassociateIamInstanceProfileOutput, error)
	DisassociateInstanceEventWindowMock                                 func(arg1 context.Context, arg2 *ec2.DisassociateInstanceEventWindowInput, arg3 ...func(*ec2.Options)) (*ec2.DisassociateInstanceEventWindowOutput, error)
	DisassociateRouteTableMock                                          func(arg1 context.Context, arg2 *ec2.DisassociateRouteTableInput, arg3 ...func(*ec2.Options)) (*ec2.DisassociateRouteTableOutput, error)
	DisassociateSubnetCidrBlockMock                                     func(arg1 context.Context, arg2 *ec2.DisassociateSubnetCidrBlockInput, arg3 ...func(*ec2.Options)) (*ec2.DisassociateSubnetCidrBlockOutput, error)
	DisassociateTransitGatewayMulticastDomainMock                       func(arg1 context.Context, arg2 *ec2.DisassociateTransitGatewayMulticastDomainInput, arg3 ...func(*ec2.Options)) (*ec2.DisassociateTransitGatewayMulticastDomainOutput, error)
	DisassociateTransitGatewayPolicyTableMock                           func(arg1 context.Context, arg2 *ec2.DisassociateTransitGatewayPolicyTableInput, arg3 ...func(*ec2.Options)) (*ec2.DisassociateTransitGatewayPolicyTableOutput, error)
	DisassociateTransitGatewayRouteTableMock                            func(arg1 context.Context, arg2 *ec2.DisassociateTransitGatewayRouteTableInput, arg3 ...func(*ec2.Options)) (*ec2.DisassociateTransitGatewayRouteTableOutput, error)
	DisassociateTrunkInterfaceMock                                      func(arg1 context.Context, arg2 *ec2.DisassociateTrunkInterfaceInput, arg3 ...func(*ec2.Options)) (*ec2.DisassociateTrunkInterfaceOutput, error)
	DisassociateVpcCidrBlockMock                                        func(arg1 context.Context, arg2 *ec2.DisassociateVpcCidrBlockInput, arg3 ...func(*ec2.Options)) (*ec2.DisassociateVpcCidrBlockOutput, error)
	EnableEbsEncryptionByDefaultMock                                    func(arg1 context.Context, arg2 *ec2.EnableEbsEncryptionByDefaultInput, arg3 ...func(*ec2.Options)) (*ec2.EnableEbsEncryptionByDefaultOutput, error)
	EnableFastLaunchMock                                                func(arg1 context.Context, arg2 *ec2.EnableFastLaunchInput, arg3 ...func(*ec2.Options)) (*ec2.EnableFastLaunchOutput, error)
	EnableFastSnapshotRestoresMock                                      func(arg1 context.Context, arg2 *ec2.EnableFastSnapshotRestoresInput, arg3 ...func(*ec2.Options)) (*ec2.EnableFastSnapshotRestoresOutput, error)
	EnableImageDeprecationMock                                          func(arg1 context.Context, arg2 *ec2.EnableImageDeprecationInput, arg3 ...func(*ec2.Options)) (*ec2.EnableImageDeprecationOutput, error)
	EnableIpamOrganizationAdminAccountMock                              func(arg1 context.Context, arg2 *ec2.EnableIpamOrganizationAdminAccountInput, arg3 ...func(*ec2.Options)) (*ec2.EnableIpamOrganizationAdminAccountOutput, error)
	EnableSerialConsoleAccessMock                                       func(arg1 context.Context, arg2 *ec2.EnableSerialConsoleAccessInput, arg3 ...func(*ec2.Options)) (*ec2.EnableSerialConsoleAccessOutput, error)
	EnableTransitGatewayRouteTablePropagationMock                       func(arg1 context.Context, arg2 *ec2.EnableTransitGatewayRouteTablePropagationInput, arg3 ...func(*ec2.Options)) (*ec2.EnableTransitGatewayRouteTablePropagationOutput, error)
	EnableVgwRoutePropagationMock                                       func(arg1 context.Context, arg2 *ec2.EnableVgwRoutePropagationInput, arg3 ...func(*ec2.Options)) (*ec2.EnableVgwRoutePropagationOutput, error)
	EnableVolumeIOMock                                                  func(arg1 context.Context, arg2 *ec2.EnableVolumeIOInput, arg3 ...func(*ec2.Options)) (*ec2.EnableVolumeIOOutput, error)
	EnableVpcClassicLinkMock                                            func(arg1 context.Context, arg2 *ec2.EnableVpcClassicLinkInput, arg3 ...func(*ec2.Options)) (*ec2.EnableVpcClassicLinkOutput, error)
	EnableVpcClassicLinkDnsSupportMock                                  func(arg1 context.Context, arg2 *ec2.EnableVpcClassicLinkDnsSupportInput, arg3 ...func(*ec2.Options)) (*ec2.EnableVpcClassicLinkDnsSupportOutput, error)
	ExportClientVpnClientCertificateRevocationListMock                  func(arg1 context.Context, arg2 *ec2.ExportClientVpnClientCertificateRevocationListInput, arg3 ...func(*ec2.Options)) (*ec2.ExportClientVpnClientCertificateRevocationListOutput, error)
	ExportClientVpnClientConfigurationMock                              func(arg1 context.Context, arg2 *ec2.ExportClientVpnClientConfigurationInput, arg3 ...func(*ec2.Options)) (*ec2.ExportClientVpnClientConfigurationOutput, error)
	ExportImageMock                                                     func(arg1 context.Context, arg2 *ec2.ExportImageInput, arg3 ...func(*ec2.Options)) (*ec2.ExportImageOutput, error)
	ExportTransitGatewayRoutesMock                                      func(arg1 context.Context, arg2 *ec2.ExportTransitGatewayRoutesInput, arg3 ...func(*ec2.Options)) (*ec2.ExportTransitGatewayRoutesOutput, error)
	GetAssociatedEnclaveCertificateIamRolesMock                         func(arg1 context.Context, arg2 *ec2.GetAssociatedEnclaveCertificateIamRolesInput, arg3 ...func(*ec2.Options)) (*ec2.GetAssociatedEnclaveCertificateIamRolesOutput, error)
	GetAssociatedIpv6PoolCidrsMock                                      func(arg1 context.Context, arg2 *ec2.GetAssociatedIpv6PoolCidrsInput, arg3 ...func(*ec2.Options)) (*ec2.GetAssociatedIpv6PoolCidrsOutput, error)
	GetCapacityReservationUsageMock                                     func(arg1 context.Context, arg2 *ec2.GetCapacityReservationUsageInput, arg3 ...func(*ec2.Options)) (*ec2.GetCapacityReservationUsageOutput, error)
	GetCoipPoolUsageMock                                                func(arg1 context.Context, arg2 *ec2.GetCoipPoolUsageInput, arg3 ...func(*ec2.Options)) (*ec2.GetCoipPoolUsageOutput, error)
	GetConsoleOutputMock                                                func(arg1 context.Context, arg2 *ec2.GetConsoleOutputInput, arg3 ...func(*ec2.Options)) (*ec2.GetConsoleOutputOutput, error)
	GetConsoleScreenshotMock                                            func(arg1 context.Context, arg2 *ec2.GetConsoleScreenshotInput, arg3 ...func(*ec2.Options)) (*ec2.GetConsoleScreenshotOutput, error)
	GetDefaultCreditSpecificationMock                                   func(arg1 context.Context, arg2 *ec2.GetDefaultCreditSpecificationInput, arg3 ...func(*ec2.Options)) (*ec2.GetDefaultCreditSpecificationOutput, error)
	GetEbsDefaultKmsKeyIdMock                                           func(arg1 context.Context, arg2 *ec2.GetEbsDefaultKmsKeyIdInput, arg3 ...func(*ec2.Options)) (*ec2.GetEbsDefaultKmsKeyIdOutput, error)
	GetEbsEncryptionByDefaultMock                                       func(arg1 context.Context, arg2 *ec2.GetEbsEncryptionByDefaultInput, arg3 ...func(*ec2.Options)) (*ec2.GetEbsEncryptionByDefaultOutput, error)
	GetFlowLogsIntegrationTemplateMock                                  func(arg1 context.Context, arg2 *ec2.GetFlowLogsIntegrationTemplateInput, arg3 ...func(*ec2.Options)) (*ec2.GetFlowLogsIntegrationTemplateOutput, error)
	GetGroupsForCapacityReservationMock                                 func(arg1 context.Context, arg2 *ec2.GetGroupsForCapacityReservationInput, arg3 ...func(*ec2.Options)) (*ec2.GetGroupsForCapacityReservationOutput, error)
	GetHostReservationPurchasePreviewMock                               func(arg1 context.Context, arg2 *ec2.GetHostReservationPurchasePreviewInput, arg3 ...func(*ec2.Options)) (*ec2.GetHostReservationPurchasePreviewOutput, error)
	GetInstanceTypesFromInstanceRequirementsMock                        func(arg1 context.Context, arg2 *ec2.GetInstanceTypesFromInstanceRequirementsInput, arg3 ...func(*ec2.Options)) (*ec2.GetInstanceTypesFromInstanceRequirementsOutput, error)
	GetInstanceUefiDataMock                                             func(arg1 context.Context, arg2 *ec2.GetInstanceUefiDataInput, arg3 ...func(*ec2.Options)) (*ec2.GetInstanceUefiDataOutput, error)
	GetIpamAddressHistoryMock                                           func(arg1 context.Context, arg2 *ec2.GetIpamAddressHistoryInput, arg3 ...func(*ec2.Options)) (*ec2.GetIpamAddressHistoryOutput, error)
	GetIpamPoolAllocationsMock                                          func(arg1 context.Context, arg2 *ec2.GetIpamPoolAllocationsInput, arg3 ...func(*ec2.Options)) (*ec2.GetIpamPoolAllocationsOutput, error)
	GetIpamPoolCidrsMock                                                func(arg1 context.Context, arg2 *ec2.GetIpamPoolCidrsInput, arg3 ...func(*ec2.Options)) (*ec2.GetIpamPoolCidrsOutput, error)
	GetIpamResourceCidrsMock                                            func(arg1 context.Context, arg2 *ec2.GetIpamResourceCidrsInput, arg3 ...func(*ec2.Options)) (*ec2.GetIpamResourceCidrsOutput, error)
	GetLaunchTemplateDataMock                                           func(arg1 context.Context, arg2 *ec2.GetLaunchTemplateDataInput, arg3 ...func(*ec2.Options)) (*ec2.GetLaunchTemplateDataOutput, error)
	GetManagedPrefixListAssociationsMock                                func(arg1 context.Context, arg2 *ec2.GetManagedPrefixListAssociationsInput, arg3 ...func(*ec2.Options)) (*ec2.GetManagedPrefixListAssociationsOutput, error)
	GetManagedPrefixListEntriesMock                                     func(arg1 context.Context, arg2 *ec2.GetManagedPrefixListEntriesInput, arg3 ...func(*ec2.Options)) (*ec2.GetManagedPrefixListEntriesOutput, error)
	GetNetworkInsightsAccessScopeAnalysisFindingsMock                   func(arg1 context.Context, arg2 *ec2.GetNetworkInsightsAccessScopeAnalysisFindingsInput, arg3 ...func(*ec2.Options)) (*ec2.GetNetworkInsightsAccessScopeAnalysisFindingsOutput, error)
	GetNetworkInsightsAccessScopeContentMock                            func(arg1 context.Context, arg2 *ec2.GetNetworkInsightsAccessScopeContentInput, arg3 ...func(*ec2.Options)) (*ec2.GetNetworkInsightsAccessScopeContentOutput, error)
	GetPasswordDataMock                                                 func(arg1 context.Context, arg2 *ec2.GetPasswordDataInput, arg3 ...func(*ec2.Options)) (*ec2.GetPasswordDataOutput, error)
	GetReservedInstancesExchangeQuoteMock                               func(arg1 context.Context, arg2 *ec2.GetReservedInstancesExchangeQuoteInput, arg3 ...func(*ec2.Options)) (*ec2.GetReservedInstancesExchangeQuoteOutput, error)
	GetSerialConsoleAccessStatusMock                                    func(arg1 context.Context, arg2 *ec2.GetSerialConsoleAccessStatusInput, arg3 ...func(*ec2.Options)) (*ec2.GetSerialConsoleAccessStatusOutput, error)
	GetSpotPlacementScoresMock                                          func(arg1 context.Context, arg2 *ec2.GetSpotPlacementScoresInput, arg3 ...func(*ec2.Options)) (*ec2.GetSpotPlacementScoresOutput, error)
	GetSubnetCidrReservationsMock                                       func(arg1 context.Context, arg2 *ec2.GetSubnetCidrReservationsInput, arg3 ...func(*ec2.Options)) (*ec2.GetSubnetCidrReservationsOutput, error)
	GetTransitGatewayAttachmentPropagationsMock                         func(arg1 context.Context, arg2 *ec2.GetTransitGatewayAttachmentPropagationsInput, arg3 ...func(*ec2.Options)) (*ec2.GetTransitGatewayAttachmentPropagationsOutput, error)
	GetTransitGatewayMulticastDomainAssociationsMock                    func(arg1 context.Context, arg2 *ec2.GetTransitGatewayMulticastDomainAssociationsInput, arg3 ...func(*ec2.Options)) (*ec2.GetTransitGatewayMulticastDomainAssociationsOutput, error)
	GetTransitGatewayPolicyTableAssociationsMock                        func(arg1 context.Context, arg2 *ec2.GetTransitGatewayPolicyTableAssociationsInput, arg3 ...func(*ec2.Options)) (*ec2.GetTransitGatewayPolicyTableAssociationsOutput, error)
	GetTransitGatewayPolicyTableEntriesMock                             func(arg1 context.Context, arg2 *ec2.GetTransitGatewayPolicyTableEntriesInput, arg3 ...func(*ec2.Options)) (*ec2.GetTransitGatewayPolicyTableEntriesOutput, error)
	GetTransitGatewayPrefixListReferencesMock                           func(arg1 context.Context, arg2 *ec2.GetTransitGatewayPrefixListReferencesInput, arg3 ...func(*ec2.Options)) (*ec2.GetTransitGatewayPrefixListReferencesOutput, error)
	GetTransitGatewayRouteTableAssociationsMock                         func(arg1 context.Context, arg2 *ec2.GetTransitGatewayRouteTableAssociationsInput, arg3 ...func(*ec2.Options)) (*ec2.GetTransitGatewayRouteTableAssociationsOutput, error)
	GetTransitGatewayRouteTablePropagationsMock                         func(arg1 context.Context, arg2 *ec2.GetTransitGatewayRouteTablePropagationsInput, arg3 ...func(*ec2.Options)) (*ec2.GetTransitGatewayRouteTablePropagationsOutput, error)
	GetVpnConnectionDeviceSampleConfigurationMock                       func(arg1 context.Context, arg2 *ec2.GetVpnConnectionDeviceSampleConfigurationInput, arg3 ...func(*ec2.Options)) (*ec2.GetVpnConnectionDeviceSampleConfigurationOutput, error)
	GetVpnConnectionDeviceTypesMock                                     func(arg1 context.Context, arg2 *ec2.GetVpnConnectionDeviceTypesInput, arg3 ...func(*ec2.Options)) (*ec2.GetVpnConnectionDeviceTypesOutput, error)
	ImportClientVpnClientCertificateRevocationListMock                  func(arg1 context.Context, arg2 *ec2.ImportClientVpnClientCertificateRevocationListInput, arg3 ...func(*ec2.Options)) (*ec2.ImportClientVpnClientCertificateRevocationListOutput, error)
	ImportImageMock                                                     func(arg1 context.Context, arg2 *ec2.ImportImageInput, arg3 ...func(*ec2.Options)) (*ec2.ImportImageOutput, error)
	ImportInstanceMock                                                  func(arg1 context.Context, arg2 *ec2.ImportInstanceInput, arg3 ...func(*ec2.Options)) (*ec2.ImportInstanceOutput, error)
	ImportKeyPairMock                                                   func(arg1 context.Context, arg2 *ec2.ImportKeyPairInput, arg3 ...func(*ec2.Options)) (*ec2.ImportKeyPairOutput, error)
	ImportSnapshotMock                                                  func(arg1 context.Context, arg2 *ec2.ImportSnapshotInput, arg3 ...func(*ec2.Options)) (*ec2.ImportSnapshotOutput, error)
	ImportVolumeMock                                                    func(arg1 context.Context, arg2 *ec2.ImportVolumeInput, arg3 ...func(*ec2.Options)) (*ec2.ImportVolumeOutput, error)
	ListImagesInRecycleBinMock                                          func(arg1 context.Context, arg2 *ec2.ListImagesInRecycleBinInput, arg3 ...func(*ec2.Options)) (*ec2.ListImagesInRecycleBinOutput, error)
	ListSnapshotsInRecycleBinMock                                       func(arg1 context.Context, arg2 *ec2.ListSnapshotsInRecycleBinInput, arg3 ...func(*ec2.Options)) (*ec2.ListSnapshotsInRecycleBinOutput, error)
	ModifyAddressAttributeMock                                          func(arg1 context.Context, arg2 *ec2.ModifyAddressAttributeInput, arg3 ...func(*ec2.Options)) (*ec2.ModifyAddressAttributeOutput, error)
	ModifyAvailabilityZoneGroupMock                                     func(arg1 context.Context, arg2 *ec2.ModifyAvailabilityZoneGroupInput, arg3 ...func(*ec2.Options)) (*ec2.ModifyAvailabilityZoneGroupOutput, error)
	ModifyCapacityReservationMock                                       func(arg1 context.Context, arg2 *ec2.ModifyCapacityReservationInput, arg3 ...func(*ec2.Options)) (*ec2.ModifyCapacityReservationOutput, error)
	ModifyCapacityReservationFleetMock                                  func(arg1 context.Context, arg2 *ec2.ModifyCapacityReservationFleetInput, arg3 ...func(*ec2.Options)) (*ec2.ModifyCapacityReservationFleetOutput, error)
	ModifyClientVpnEndpointMock                                         func(arg1 context.Context, arg2 *ec2.ModifyClientVpnEndpointInput, arg3 ...func(*ec2.Options)) (*ec2.ModifyClientVpnEndpointOutput, error)
	ModifyDefaultCreditSpecificationMock                                func(arg1 context.Context, arg2 *ec2.ModifyDefaultCreditSpecificationInput, arg3 ...func(*ec2.Options)) (*ec2.ModifyDefaultCreditSpecificationOutput, error)
	ModifyEbsDefaultKmsKeyIdMock                                        func(arg1 context.Context, arg2 *ec2.ModifyEbsDefaultKmsKeyIdInput, arg3 ...func(*ec2.Options)) (*ec2.ModifyEbsDefaultKmsKeyIdOutput, error)
	ModifyFleetMock                                                     func(arg1 context.Context, arg2 *ec2.ModifyFleetInput, arg3 ...func(*ec2.Options)) (*ec2.ModifyFleetOutput, error)
	ModifyFpgaImageAttributeMock                                        func(arg1 context.Context, arg2 *ec2.ModifyFpgaImageAttributeInput, arg3 ...func(*ec2.Options)) (*ec2.ModifyFpgaImageAttributeOutput, error)
	ModifyHostsMock                                                     func(arg1 context.Context, arg2 *ec2.ModifyHostsInput, arg3 ...func(*ec2.Options)) (*ec2.ModifyHostsOutput, error)
	ModifyIdFormatMock                                                  func(arg1 context.Context, arg2 *ec2.ModifyIdFormatInput, arg3 ...func(*ec2.Options)) (*ec2.ModifyIdFormatOutput, error)
	ModifyIdentityIdFormatMock                                          func(arg1 context.Context, arg2 *ec2.ModifyIdentityIdFormatInput, arg3 ...func(*ec2.Options)) (*ec2.ModifyIdentityIdFormatOutput, error)
	ModifyImageAttributeMock                                            func(arg1 context.Context, arg2 *ec2.ModifyImageAttributeInput, arg3 ...func(*ec2.Options)) (*ec2.ModifyImageAttributeOutput, error)
	ModifyInstanceAttributeMock                                         func(arg1 context.Context, arg2 *ec2.ModifyInstanceAttributeInput, arg3 ...func(*ec2.Options)) (*ec2.ModifyInstanceAttributeOutput, error)
	ModifyInstanceCapacityReservationAttributesMock                     func(arg1 context.Context, arg2 *ec2.ModifyInstanceCapacityReservationAttributesInput, arg3 ...func(*ec2.Options)) (*ec2.ModifyInstanceCapacityReservationAttributesOutput, error)
	ModifyInstanceCreditSpecificationMock                               func(arg1 context.Context, arg2 *ec2.ModifyInstanceCreditSpecificationInput, arg3 ...func(*ec2.Options)) (*ec2.ModifyInstanceCreditSpecificationOutput, error)
	ModifyInstanceEventStartTimeMock                                    func(arg1 context.Context, arg2 *ec2.ModifyInstanceEventStartTimeInput, arg3 ...func(*ec2.Options)) (*ec2.ModifyInstanceEventStartTimeOutput, error)
	ModifyInstanceEventWindowMock                                       func(arg1 context.Context, arg2 *ec2.ModifyInstanceEventWindowInput, arg3 ...func(*ec2.Options)) (*ec2.ModifyInstanceEventWindowOutput, error)
	ModifyInstanceMaintenanceOptionsMock                                func(arg1 context.Context, arg2 *ec2.ModifyInstanceMaintenanceOptionsInput, arg3 ...func(*ec2.Options)) (*ec2.ModifyInstanceMaintenanceOptionsOutput, error)
	ModifyInstanceMetadataOptionsMock                                   func(arg1 context.Context, arg2 *ec2.ModifyInstanceMetadataOptionsInput, arg3 ...func(*ec2.Options)) (*ec2.ModifyInstanceMetadataOptionsOutput, error)
	ModifyInstancePlacementMock                                         func(arg1 context.Context, arg2 *ec2.ModifyInstancePlacementInput, arg3 ...func(*ec2.Options)) (*ec2.ModifyInstancePlacementOutput, error)
	ModifyIpamMock                                                      func(arg1 context.Context, arg2 *ec2.ModifyIpamInput, arg3 ...func(*ec2.Options)) (*ec2.ModifyIpamOutput, error)
	ModifyIpamPoolMock                                                  func(arg1 context.Context, arg2 *ec2.ModifyIpamPoolInput, arg3 ...func(*ec2.Options)) (*ec2.ModifyIpamPoolOutput, error)
	ModifyIpamResourceCidrMock                                          func(arg1 context.Context, arg2 *ec2.ModifyIpamResourceCidrInput, arg3 ...func(*ec2.Options)) (*ec2.ModifyIpamResourceCidrOutput, error)
	ModifyIpamScopeMock                                                 func(arg1 context.Context, arg2 *ec2.ModifyIpamScopeInput, arg3 ...func(*ec2.Options)) (*ec2.ModifyIpamScopeOutput, error)
	ModifyLaunchTemplateMock                                            func(arg1 context.Context, arg2 *ec2.ModifyLaunchTemplateInput, arg3 ...func(*ec2.Options)) (*ec2.ModifyLaunchTemplateOutput, error)
	ModifyManagedPrefixListMock                                         func(arg1 context.Context, arg2 *ec2.ModifyManagedPrefixListInput, arg3 ...func(*ec2.Options)) (*ec2.ModifyManagedPrefixListOutput, error)
	ModifyNetworkInterfaceAttributeMock                                 func(arg1 context.Context, arg2 *ec2.ModifyNetworkInterfaceAttributeInput, arg3 ...func(*ec2.Options)) (*ec2.ModifyNetworkInterfaceAttributeOutput, error)
	ModifyPrivateDnsNameOptionsMock                                     func(arg1 context.Context, arg2 *ec2.ModifyPrivateDnsNameOptionsInput, arg3 ...func(*ec2.Options)) (*ec2.ModifyPrivateDnsNameOptionsOutput, error)
	ModifyReservedInstancesMock                                         func(arg1 context.Context, arg2 *ec2.ModifyReservedInstancesInput, arg3 ...func(*ec2.Options)) (*ec2.ModifyReservedInstancesOutput, error)
	ModifySecurityGroupRulesMock                                        func(arg1 context.Context, arg2 *ec2.ModifySecurityGroupRulesInput, arg3 ...func(*ec2.Options)) (*ec2.ModifySecurityGroupRulesOutput, error)
	ModifySnapshotAttributeMock                                         func(arg1 context.Context, arg2 *ec2.ModifySnapshotAttributeInput, arg3 ...func(*ec2.Options)) (*ec2.ModifySnapshotAttributeOutput, error)
	ModifySnapshotTierMock                                              func(arg1 context.Context, arg2 *ec2.ModifySnapshotTierInput, arg3 ...func(*ec2.Options)) (*ec2.ModifySnapshotTierOutput, error)
	ModifySpotFleetRequestMock                                          func(arg1 context.Context, arg2 *ec2.ModifySpotFleetRequestInput, arg3 ...func(*ec2.Options)) (*ec2.ModifySpotFleetRequestOutput, error)
	ModifySubnetAttributeMock                                           func(arg1 context.Context, arg2 *ec2.ModifySubnetAttributeInput, arg3 ...func(*ec2.Options)) (*ec2.ModifySubnetAttributeOutput, error)
	ModifyTrafficMirrorFilterNetworkServicesMock                        func(arg1 context.Context, arg2 *ec2.ModifyTrafficMirrorFilterNetworkServicesInput, arg3 ...func(*ec2.Options)) (*ec2.ModifyTrafficMirrorFilterNetworkServicesOutput, error)
	ModifyTrafficMirrorFilterRuleMock                                   func(arg1 context.Context, arg2 *ec2.ModifyTrafficMirrorFilterRuleInput, arg3 ...func(*ec2.Options)) (*ec2.ModifyTrafficMirrorFilterRuleOutput, error)
	ModifyTrafficMirrorSessionMock                                      func(arg1 context.Context, arg2 *ec2.ModifyTrafficMirrorSessionInput, arg3 ...func(*ec2.Options)) (*ec2.ModifyTrafficMirrorSessionOutput, error)
	ModifyTransitGatewayMock                                            func(arg1 context.Context, arg2 *ec2.ModifyTransitGatewayInput, arg3 ...func(*ec2.Options)) (*ec2.ModifyTransitGatewayOutput, error)
	ModifyTransitGatewayPrefixListReferenceMock                         func(arg1 context.Context, arg2 *ec2.ModifyTransitGatewayPrefixListReferenceInput, arg3 ...func(*ec2.Options)) (*ec2.ModifyTransitGatewayPrefixListReferenceOutput, error)
	ModifyTransitGatewayVpcAttachmentMock                               func(arg1 context.Context, arg2 *ec2.ModifyTransitGatewayVpcAttachmentInput, arg3 ...func(*ec2.Options)) (*ec2.ModifyTransitGatewayVpcAttachmentOutput, error)
	ModifyVolumeMock                                                    func(arg1 context.Context, arg2 *ec2.ModifyVolumeInput, arg3 ...func(*ec2.Options)) (*ec2.ModifyVolumeOutput, error)
	ModifyVolumeAttributeMock                                           func(arg1 context.Context, arg2 *ec2.ModifyVolumeAttributeInput, arg3 ...func(*ec2.Options)) (*ec2.ModifyVolumeAttributeOutput, error)
	ModifyVpcAttributeMock                                              func(arg1 context.Context, arg2 *ec2.ModifyVpcAttributeInput, arg3 ...func(*ec2.Options)) (*ec2.ModifyVpcAttributeOutput, error)
	ModifyVpcEndpointMock                                               func(arg1 context.Context, arg2 *ec2.ModifyVpcEndpointInput, arg3 ...func(*ec2.Options)) (*ec2.ModifyVpcEndpointOutput, error)
	ModifyVpcEndpointConnectionNotificationMock                         func(arg1 context.Context, arg2 *ec2.ModifyVpcEndpointConnectionNotificationInput, arg3 ...func(*ec2.Options)) (*ec2.ModifyVpcEndpointConnectionNotificationOutput, error)
	ModifyVpcEndpointServiceConfigurationMock                           func(arg1 context.Context, arg2 *ec2.ModifyVpcEndpointServiceConfigurationInput, arg3 ...func(*ec2.Options)) (*ec2.ModifyVpcEndpointServiceConfigurationOutput, error)
	ModifyVpcEndpointServicePayerResponsibilityMock                     func(arg1 context.Context, arg2 *ec2.ModifyVpcEndpointServicePayerResponsibilityInput, arg3 ...func(*ec2.Options)) (*ec2.ModifyVpcEndpointServicePayerResponsibilityOutput, error)
	ModifyVpcEndpointServicePermissionsMock                             func(arg1 context.Context, arg2 *ec2.ModifyVpcEndpointServicePermissionsInput, arg3 ...func(*ec2.Options)) (*ec2.ModifyVpcEndpointServicePermissionsOutput, error)
	ModifyVpcPeeringConnectionOptionsMock                               func(arg1 context.Context, arg2 *ec2.ModifyVpcPeeringConnectionOptionsInput, arg3 ...func(*ec2.Options)) (*ec2.ModifyVpcPeeringConnectionOptionsOutput, error)
	ModifyVpcTenancyMock                                                func(arg1 context.Context, arg2 *ec2.ModifyVpcTenancyInput, arg3 ...func(*ec2.Options)) (*ec2.ModifyVpcTenancyOutput, error)
	ModifyVpnConnectionMock                                             func(arg1 context.Context, arg2 *ec2.ModifyVpnConnectionInput, arg3 ...func(*ec2.Options)) (*ec2.ModifyVpnConnectionOutput, error)
	ModifyVpnConnectionOptionsMock                                      func(arg1 context.Context, arg2 *ec2.ModifyVpnConnectionOptionsInput, arg3 ...func(*ec2.Options)) (*ec2.ModifyVpnConnectionOptionsOutput, error)
	ModifyVpnTunnelCertificateMock                                      func(arg1 context.Context, arg2 *ec2.ModifyVpnTunnelCertificateInput, arg3 ...func(*ec2.Options)) (*ec2.ModifyVpnTunnelCertificateOutput, error)
	ModifyVpnTunnelOptionsMock                                          func(arg1 context.Context, arg2 *ec2.ModifyVpnTunnelOptionsInput, arg3 ...func(*ec2.Options)) (*ec2.ModifyVpnTunnelOptionsOutput, error)
	MonitorInstancesMock                                                func(arg1 context.Context, arg2 *ec2.MonitorInstancesInput, arg3 ...func(*ec2.Options)) (*ec2.MonitorInstancesOutput, error)
	MoveAddressToVpcMock                                                func(arg1 context.Context, arg2 *ec2.MoveAddressToVpcInput, arg3 ...func(*ec2.Options)) (*ec2.MoveAddressToVpcOutput, error)
	MoveByoipCidrToIpamMock                                             func(arg1 context.Context, arg2 *ec2.MoveByoipCidrToIpamInput, arg3 ...func(*ec2.Options)) (*ec2.MoveByoipCidrToIpamOutput, error)
	ProvisionByoipCidrMock                                              func(arg1 context.Context, arg2 *ec2.ProvisionByoipCidrInput, arg3 ...func(*ec2.Options)) (*ec2.ProvisionByoipCidrOutput, error)
	ProvisionIpamPoolCidrMock                                           func(arg1 context.Context, arg2 *ec2.ProvisionIpamPoolCidrInput, arg3 ...func(*ec2.Options)) (*ec2.ProvisionIpamPoolCidrOutput, error)
	ProvisionPublicIpv4PoolCidrMock                                     func(arg1 context.Context, arg2 *ec2.ProvisionPublicIpv4PoolCidrInput, arg3 ...func(*ec2.Options)) (*ec2.ProvisionPublicIpv4PoolCidrOutput, error)
	PurchaseHostReservationMock                                         func(arg1 context.Context, arg2 *ec2.PurchaseHostReservationInput, arg3 ...func(*ec2.Options)) (*ec2.PurchaseHostReservationOutput, error)
	PurchaseReservedInstancesOfferingMock                               func(arg1 context.Context, arg2 *ec2.PurchaseReservedInstancesOfferingInput, arg3 ...func(*ec2.Options)) (*ec2.PurchaseReservedInstancesOfferingOutput, error)
	PurchaseScheduledInstancesMock                                      func(arg1 context.Context, arg2 *ec2.PurchaseScheduledInstancesInput, arg3 ...func(*ec2.Options)) (*ec2.PurchaseScheduledInstancesOutput, error)
	RebootInstancesMock                                                 func(arg1 context.Context, arg2 *ec2.RebootInstancesInput, arg3 ...func(*ec2.Options)) (*ec2.RebootInstancesOutput, error)
	RegisterImageMock                                                   func(arg1 context.Context, arg2 *ec2.RegisterImageInput, arg3 ...func(*ec2.Options)) (*ec2.RegisterImageOutput, error)
	RegisterInstanceEventNotificationAttributesMock                     func(arg1 context.Context, arg2 *ec2.RegisterInstanceEventNotificationAttributesInput, arg3 ...func(*ec2.Options)) (*ec2.RegisterInstanceEventNotificationAttributesOutput, error)
	RegisterTransitGatewayMulticastGroupMembersMock                     func(arg1 context.Context, arg2 *ec2.RegisterTransitGatewayMulticastGroupMembersInput, arg3 ...func(*ec2.Options)) (*ec2.RegisterTransitGatewayMulticastGroupMembersOutput, error)
	RegisterTransitGatewayMulticastGroupSourcesMock                     func(arg1 context.Context, arg2 *ec2.RegisterTransitGatewayMulticastGroupSourcesInput, arg3 ...func(*ec2.Options)) (*ec2.RegisterTransitGatewayMulticastGroupSourcesOutput, error)
	RejectTransitGatewayMulticastDomainAssociationsMock                 func(arg1 context.Context, arg2 *ec2.RejectTransitGatewayMulticastDomainAssociationsInput, arg3 ...func(*ec2.Options)) (*ec2.RejectTransitGatewayMulticastDomainAssociationsOutput, error)
	RejectTransitGatewayPeeringAttachmentMock                           func(arg1 context.Context, arg2 *ec2.RejectTransitGatewayPeeringAttachmentInput, arg3 ...func(*ec2.Options)) (*ec2.RejectTransitGatewayPeeringAttachmentOutput, error)
	RejectTransitGatewayVpcAttachmentMock                               func(arg1 context.Context, arg2 *ec2.RejectTransitGatewayVpcAttachmentInput, arg3 ...func(*ec2.Options)) (*ec2.RejectTransitGatewayVpcAttachmentOutput, error)
	RejectVpcEndpointConnectionsMock                                    func(arg1 context.Context, arg2 *ec2.RejectVpcEndpointConnectionsInput, arg3 ...func(*ec2.Options)) (*ec2.RejectVpcEndpointConnectionsOutput, error)
	RejectVpcPeeringConnectionMock                                      func(arg1 context.Context, arg2 *ec2.RejectVpcPeeringConnectionInput, arg3 ...func(*ec2.Options)) (*ec2.RejectVpcPeeringConnectionOutput, error)
	ReleaseAddressMock                                                  func(arg1 context.Context, arg2 *ec2.ReleaseAddressInput, arg3 ...func(*ec2.Options)) (*ec2.ReleaseAddressOutput, error)
	ReleaseHostsMock                                                    func(arg1 context.Context, arg2 *ec2.ReleaseHostsInput, arg3 ...func(*ec2.Options)) (*ec2.ReleaseHostsOutput, error)
	ReleaseIpamPoolAllocationMock                                       func(arg1 context.Context, arg2 *ec2.ReleaseIpamPoolAllocationInput, arg3 ...func(*ec2.Options)) (*ec2.ReleaseIpamPoolAllocationOutput, error)
	ReplaceIamInstanceProfileAssociationMock                            func(arg1 context.Context, arg2 *ec2.ReplaceIamInstanceProfileAssociationInput, arg3 ...func(*ec2.Options)) (*ec2.ReplaceIamInstanceProfileAssociationOutput, error)
	ReplaceNetworkAclAssociationMock                                    func(arg1 context.Context, arg2 *ec2.ReplaceNetworkAclAssociationInput, arg3 ...func(*ec2.Options)) (*ec2.ReplaceNetworkAclAssociationOutput, error)
	ReplaceNetworkAclEntryMock                                          func(arg1 context.Context, arg2 *ec2.ReplaceNetworkAclEntryInput, arg3 ...func(*ec2.Options)) (*ec2.ReplaceNetworkAclEntryOutput, error)
	ReplaceRouteMock                                                    func(arg1 context.Context, arg2 *ec2.ReplaceRouteInput, arg3 ...func(*ec2.Options)) (*ec2.ReplaceRouteOutput, error)
	ReplaceRouteTableAssociationMock                                    func(arg1 context.Context, arg2 *ec2.ReplaceRouteTableAssociationInput, arg3 ...func(*ec2.Options)) (*ec2.ReplaceRouteTableAssociationOutput, error)
	ReplaceTransitGatewayRouteMock                                      func(arg1 context.Context, arg2 *ec2.ReplaceTransitGatewayRouteInput, arg3 ...func(*ec2.Options)) (*ec2.ReplaceTransitGatewayRouteOutput, error)
	ReportInstanceStatusMock                                            func(arg1 context.Context, arg2 *ec2.ReportInstanceStatusInput, arg3 ...func(*ec2.Options)) (*ec2.ReportInstanceStatusOutput, error)
	RequestSpotFleetMock                                                func(arg1 context.Context, arg2 *ec2.RequestSpotFleetInput, arg3 ...func(*ec2.Options)) (*ec2.RequestSpotFleetOutput, error)
	RequestSpotInstancesMock                                            func(arg1 context.Context, arg2 *ec2.RequestSpotInstancesInput, arg3 ...func(*ec2.Options)) (*ec2.RequestSpotInstancesOutput, error)
	ResetAddressAttributeMock                                           func(arg1 context.Context, arg2 *ec2.ResetAddressAttributeInput, arg3 ...func(*ec2.Options)) (*ec2.ResetAddressAttributeOutput, error)
	ResetEbsDefaultKmsKeyIdMock                                         func(arg1 context.Context, arg2 *ec2.ResetEbsDefaultKmsKeyIdInput, arg3 ...func(*ec2.Options)) (*ec2.ResetEbsDefaultKmsKeyIdOutput, error)
	ResetFpgaImageAttributeMock                                         func(arg1 context.Context, arg2 *ec2.ResetFpgaImageAttributeInput, arg3 ...func(*ec2.Options)) (*ec2.ResetFpgaImageAttributeOutput, error)
	ResetImageAttributeMock                                             func(arg1 context.Context, arg2 *ec2.ResetImageAttributeInput, arg3 ...func(*ec2.Options)) (*ec2.ResetImageAttributeOutput, error)
	ResetInstanceAttributeMock                                          func(arg1 context.Context, arg2 *ec2.ResetInstanceAttributeInput, arg3 ...func(*ec2.Options)) (*ec2.ResetInstanceAttributeOutput, error)
	ResetNetworkInterfaceAttributeMock                                  func(arg1 context.Context, arg2 *ec2.ResetNetworkInterfaceAttributeInput, arg3 ...func(*ec2.Options)) (*ec2.ResetNetworkInterfaceAttributeOutput, error)
	ResetSnapshotAttributeMock                                          func(arg1 context.Context, arg2 *ec2.ResetSnapshotAttributeInput, arg3 ...func(*ec2.Options)) (*ec2.ResetSnapshotAttributeOutput, error)
	RestoreAddressToClassicMock                                         func(arg1 context.Context, arg2 *ec2.RestoreAddressToClassicInput, arg3 ...func(*ec2.Options)) (*ec2.RestoreAddressToClassicOutput, error)
	RestoreImageFromRecycleBinMock                                      func(arg1 context.Context, arg2 *ec2.RestoreImageFromRecycleBinInput, arg3 ...func(*ec2.Options)) (*ec2.RestoreImageFromRecycleBinOutput, error)
	RestoreManagedPrefixListVersionMock                                 func(arg1 context.Context, arg2 *ec2.RestoreManagedPrefixListVersionInput, arg3 ...func(*ec2.Options)) (*ec2.RestoreManagedPrefixListVersionOutput, error)
	RestoreSnapshotFromRecycleBinMock                                   func(arg1 context.Context, arg2 *ec2.RestoreSnapshotFromRecycleBinInput, arg3 ...func(*ec2.Options)) (*ec2.RestoreSnapshotFromRecycleBinOutput, error)
	RestoreSnapshotTierMock                                             func(arg1 context.Context, arg2 *ec2.RestoreSnapshotTierInput, arg3 ...func(*ec2.Options)) (*ec2.RestoreSnapshotTierOutput, error)
	RevokeClientVpnIngressMock                                          func(arg1 context.Context, arg2 *ec2.RevokeClientVpnIngressInput, arg3 ...func(*ec2.Options)) (*ec2.RevokeClientVpnIngressOutput, error)
	RevokeSecurityGroupEgressMock                                       func(arg1 context.Context, arg2 *ec2.RevokeSecurityGroupEgressInput, arg3 ...func(*ec2.Options)) (*ec2.RevokeSecurityGroupEgressOutput, error)
	RevokeSecurityGroupIngressMock                                      func(arg1 context.Context, arg2 *ec2.RevokeSecurityGroupIngressInput, arg3 ...func(*ec2.Options)) (*ec2.RevokeSecurityGroupIngressOutput, error)
	RunInstancesMock                                                    func(arg1 context.Context, arg2 *ec2.RunInstancesInput, arg3 ...func(*ec2.Options)) (*ec2.RunInstancesOutput, error)
	RunScheduledInstancesMock                                           func(arg1 context.Context, arg2 *ec2.RunScheduledInstancesInput, arg3 ...func(*ec2.Options)) (*ec2.RunScheduledInstancesOutput, error)
	SearchLocalGatewayRoutesMock                                        func(arg1 context.Context, arg2 *ec2.SearchLocalGatewayRoutesInput, arg3 ...func(*ec2.Options)) (*ec2.SearchLocalGatewayRoutesOutput, error)
	SearchTransitGatewayMulticastGroupsMock                             func(arg1 context.Context, arg2 *ec2.SearchTransitGatewayMulticastGroupsInput, arg3 ...func(*ec2.Options)) (*ec2.SearchTransitGatewayMulticastGroupsOutput, error)
	SearchTransitGatewayRoutesMock                                      func(arg1 context.Context, arg2 *ec2.SearchTransitGatewayRoutesInput, arg3 ...func(*ec2.Options)) (*ec2.SearchTransitGatewayRoutesOutput, error)
	SendDiagnosticInterruptMock                                         func(arg1 context.Context, arg2 *ec2.SendDiagnosticInterruptInput, arg3 ...func(*ec2.Options)) (*ec2.SendDiagnosticInterruptOutput, error)
	StartInstancesMock                                                  func(arg1 context.Context, arg2 *ec2.StartInstancesInput, arg3 ...func(*ec2.Options)) (*ec2.StartInstancesOutput, error)
	StartNetworkInsightsAccessScopeAnalysisMock                         func(arg1 context.Context, arg2 *ec2.StartNetworkInsightsAccessScopeAnalysisInput, arg3 ...func(*ec2.Options)) (*ec2.StartNetworkInsightsAccessScopeAnalysisOutput, error)
	StartNetworkInsightsAnalysisMock                                    func(arg1 context.Context, arg2 *ec2.StartNetworkInsightsAnalysisInput, arg3 ...func(*ec2.Options)) (*ec2.StartNetworkInsightsAnalysisOutput, error)
	StartVpcEndpointServicePrivateDnsVerificationMock                   func(arg1 context.Context, arg2 *ec2.StartVpcEndpointServicePrivateDnsVerificationInput, arg3 ...func(*ec2.Options)) (*ec2.StartVpcEndpointServicePrivateDnsVerificationOutput, error)
	StopInstancesMock                                                   func(arg1 context.Context, arg2 *ec2.StopInstancesInput, arg3 ...func(*ec2.Options)) (*ec2.StopInstancesOutput, error)
	TerminateClientVpnConnectionsMock                                   func(arg1 context.Context, arg2 *ec2.TerminateClientVpnConnectionsInput, arg3 ...func(*ec2.Options)) (*ec2.TerminateClientVpnConnectionsOutput, error)
	TerminateInstancesMock                                              func(arg1 context.Context, arg2 *ec2.TerminateInstancesInput, arg3 ...func(*ec2.Options)) (*ec2.TerminateInstancesOutput, error)
	UnassignIpv6AddressesMock                                           func(arg1 context.Context, arg2 *ec2.UnassignIpv6AddressesInput, arg3 ...func(*ec2.Options)) (*ec2.UnassignIpv6AddressesOutput, error)
	UnassignPrivateIpAddressesMock                                      func(arg1 context.Context, arg2 *ec2.UnassignPrivateIpAddressesInput, arg3 ...func(*ec2.Options)) (*ec2.UnassignPrivateIpAddressesOutput, error)
	UnmonitorInstancesMock                                              func(arg1 context.Context, arg2 *ec2.UnmonitorInstancesInput, arg3 ...func(*ec2.Options)) (*ec2.UnmonitorInstancesOutput, error)
	UpdateSecurityGroupRuleDescriptionsEgressMock                       func(arg1 context.Context, arg2 *ec2.UpdateSecurityGroupRuleDescriptionsEgressInput, arg3 ...func(*ec2.Options)) (*ec2.UpdateSecurityGroupRuleDescriptionsEgressOutput, error)
	UpdateSecurityGroupRuleDescriptionsIngressMock                      func(arg1 context.Context, arg2 *ec2.UpdateSecurityGroupRuleDescriptionsIngressInput, arg3 ...func(*ec2.Options)) (*ec2.UpdateSecurityGroupRuleDescriptionsIngressOutput, error)
	WithdrawByoipCidrMock                                               func(arg1 context.Context, arg2 *ec2.WithdrawByoipCidrInput, arg3 ...func(*ec2.Options)) (*ec2.WithdrawByoipCidrOutput, error)
}

// AcceptReservedInstancesExchangeQuote mocked function
func (m EC2ClientMock) AcceptReservedInstancesExchangeQuote(arg1 context.Context, arg2 *ec2.AcceptReservedInstancesExchangeQuoteInput, arg3 ...func(*ec2.Options)) (*ec2.AcceptReservedInstancesExchangeQuoteOutput, error) {
	return m.AcceptReservedInstancesExchangeQuoteMock(arg1, arg2, arg3...)
}

// AcceptTransitGatewayMulticastDomainAssociations mocked function
func (m EC2ClientMock) AcceptTransitGatewayMulticastDomainAssociations(arg1 context.Context, arg2 *ec2.AcceptTransitGatewayMulticastDomainAssociationsInput, arg3 ...func(*ec2.Options)) (*ec2.AcceptTransitGatewayMulticastDomainAssociationsOutput, error) {
	return m.AcceptTransitGatewayMulticastDomainAssociationsMock(arg1, arg2, arg3...)
}

// AcceptTransitGatewayPeeringAttachment mocked function
func (m EC2ClientMock) AcceptTransitGatewayPeeringAttachment(arg1 context.Context, arg2 *ec2.AcceptTransitGatewayPeeringAttachmentInput, arg3 ...func(*ec2.Options)) (*ec2.AcceptTransitGatewayPeeringAttachmentOutput, error) {
	return m.AcceptTransitGatewayPeeringAttachmentMock(arg1, arg2, arg3...)
}

// AcceptTransitGatewayVpcAttachment mocked function
func (m EC2ClientMock) AcceptTransitGatewayVpcAttachment(arg1 context.Context, arg2 *ec2.AcceptTransitGatewayVpcAttachmentInput, arg3 ...func(*ec2.Options)) (*ec2.AcceptTransitGatewayVpcAttachmentOutput, error) {
	return m.AcceptTransitGatewayVpcAttachmentMock(arg1, arg2, arg3...)
}

// AcceptVpcEndpointConnections mocked function
func (m EC2ClientMock) AcceptVpcEndpointConnections(arg1 context.Context, arg2 *ec2.AcceptVpcEndpointConnectionsInput, arg3 ...func(*ec2.Options)) (*ec2.AcceptVpcEndpointConnectionsOutput, error) {
	return m.AcceptVpcEndpointConnectionsMock(arg1, arg2, arg3...)
}

// AcceptVpcPeeringConnection mocked function
func (m EC2ClientMock) AcceptVpcPeeringConnection(arg1 context.Context, arg2 *ec2.AcceptVpcPeeringConnectionInput, arg3 ...func(*ec2.Options)) (*ec2.AcceptVpcPeeringConnectionOutput, error) {
	return m.AcceptVpcPeeringConnectionMock(arg1, arg2, arg3...)
}

// AdvertiseByoipCidr mocked function
func (m EC2ClientMock) AdvertiseByoipCidr(arg1 context.Context, arg2 *ec2.AdvertiseByoipCidrInput, arg3 ...func(*ec2.Options)) (*ec2.AdvertiseByoipCidrOutput, error) {
	return m.AdvertiseByoipCidrMock(arg1, arg2, arg3...)
}

// AllocateAddress mocked function
func (m EC2ClientMock) AllocateAddress(arg1 context.Context, arg2 *ec2.AllocateAddressInput, arg3 ...func(*ec2.Options)) (*ec2.AllocateAddressOutput, error) {
	return m.AllocateAddressMock(arg1, arg2, arg3...)
}

// AllocateHosts mocked function
func (m EC2ClientMock) AllocateHosts(arg1 context.Context, arg2 *ec2.AllocateHostsInput, arg3 ...func(*ec2.Options)) (*ec2.AllocateHostsOutput, error) {
	return m.AllocateHostsMock(arg1, arg2, arg3...)
}

// AllocateIpamPoolCidr mocked function
func (m EC2ClientMock) AllocateIpamPoolCidr(arg1 context.Context, arg2 *ec2.AllocateIpamPoolCidrInput, arg3 ...func(*ec2.Options)) (*ec2.AllocateIpamPoolCidrOutput, error) {
	return m.AllocateIpamPoolCidrMock(arg1, arg2, arg3...)
}

// ApplySecurityGroupsToClientVpnTargetNetwork mocked function
func (m EC2ClientMock) ApplySecurityGroupsToClientVpnTargetNetwork(arg1 context.Context, arg2 *ec2.ApplySecurityGroupsToClientVpnTargetNetworkInput, arg3 ...func(*ec2.Options)) (*ec2.ApplySecurityGroupsToClientVpnTargetNetworkOutput, error) {
	return m.ApplySecurityGroupsToClientVpnTargetNetworkMock(arg1, arg2, arg3...)
}

// AssignIpv6Addresses mocked function
func (m EC2ClientMock) AssignIpv6Addresses(arg1 context.Context, arg2 *ec2.AssignIpv6AddressesInput, arg3 ...func(*ec2.Options)) (*ec2.AssignIpv6AddressesOutput, error) {
	return m.AssignIpv6AddressesMock(arg1, arg2, arg3...)
}

// AssignPrivateIpAddresses mocked function
func (m EC2ClientMock) AssignPrivateIpAddresses(arg1 context.Context, arg2 *ec2.AssignPrivateIpAddressesInput, arg3 ...func(*ec2.Options)) (*ec2.AssignPrivateIpAddressesOutput, error) {
	return m.AssignPrivateIpAddressesMock(arg1, arg2, arg3...)
}

// AssociateAddress mocked function
func (m EC2ClientMock) AssociateAddress(arg1 context.Context, arg2 *ec2.AssociateAddressInput, arg3 ...func(*ec2.Options)) (*ec2.AssociateAddressOutput, error) {
	return m.AssociateAddressMock(arg1, arg2, arg3...)
}

// AssociateClientVpnTargetNetwork mocked function
func (m EC2ClientMock) AssociateClientVpnTargetNetwork(arg1 context.Context, arg2 *ec2.AssociateClientVpnTargetNetworkInput, arg3 ...func(*ec2.Options)) (*ec2.AssociateClientVpnTargetNetworkOutput, error) {
	return m.AssociateClientVpnTargetNetworkMock(arg1, arg2, arg3...)
}

// AssociateDhcpOptions mocked function
func (m EC2ClientMock) AssociateDhcpOptions(arg1 context.Context, arg2 *ec2.AssociateDhcpOptionsInput, arg3 ...func(*ec2.Options)) (*ec2.AssociateDhcpOptionsOutput, error) {
	return m.AssociateDhcpOptionsMock(arg1, arg2, arg3...)
}

// AssociateEnclaveCertificateIamRole mocked function
func (m EC2ClientMock) AssociateEnclaveCertificateIamRole(arg1 context.Context, arg2 *ec2.AssociateEnclaveCertificateIamRoleInput, arg3 ...func(*ec2.Options)) (*ec2.AssociateEnclaveCertificateIamRoleOutput, error) {
	return m.AssociateEnclaveCertificateIamRoleMock(arg1, arg2, arg3...)
}

// AssociateIamInstanceProfile mocked function
func (m EC2ClientMock) AssociateIamInstanceProfile(arg1 context.Context, arg2 *ec2.AssociateIamInstanceProfileInput, arg3 ...func(*ec2.Options)) (*ec2.AssociateIamInstanceProfileOutput, error) {
	return m.AssociateIamInstanceProfileMock(arg1, arg2, arg3...)
}

// AssociateInstanceEventWindow mocked function
func (m EC2ClientMock) AssociateInstanceEventWindow(arg1 context.Context, arg2 *ec2.AssociateInstanceEventWindowInput, arg3 ...func(*ec2.Options)) (*ec2.AssociateInstanceEventWindowOutput, error) {
	return m.AssociateInstanceEventWindowMock(arg1, arg2, arg3...)
}

// AssociateRouteTable mocked function
func (m EC2ClientMock) AssociateRouteTable(arg1 context.Context, arg2 *ec2.AssociateRouteTableInput, arg3 ...func(*ec2.Options)) (*ec2.AssociateRouteTableOutput, error) {
	return m.AssociateRouteTableMock(arg1, arg2, arg3...)
}

// AssociateSubnetCidrBlock mocked function
func (m EC2ClientMock) AssociateSubnetCidrBlock(arg1 context.Context, arg2 *ec2.AssociateSubnetCidrBlockInput, arg3 ...func(*ec2.Options)) (*ec2.AssociateSubnetCidrBlockOutput, error) {
	return m.AssociateSubnetCidrBlockMock(arg1, arg2, arg3...)
}

// AssociateTransitGatewayMulticastDomain mocked function
func (m EC2ClientMock) AssociateTransitGatewayMulticastDomain(arg1 context.Context, arg2 *ec2.AssociateTransitGatewayMulticastDomainInput, arg3 ...func(*ec2.Options)) (*ec2.AssociateTransitGatewayMulticastDomainOutput, error) {
	return m.AssociateTransitGatewayMulticastDomainMock(arg1, arg2, arg3...)
}

// AssociateTransitGatewayPolicyTable mocked function
func (m EC2ClientMock) AssociateTransitGatewayPolicyTable(arg1 context.Context, arg2 *ec2.AssociateTransitGatewayPolicyTableInput, arg3 ...func(*ec2.Options)) (*ec2.AssociateTransitGatewayPolicyTableOutput, error) {
	return m.AssociateTransitGatewayPolicyTableMock(arg1, arg2, arg3...)
}

// AssociateTransitGatewayRouteTable mocked function
func (m EC2ClientMock) AssociateTransitGatewayRouteTable(arg1 context.Context, arg2 *ec2.AssociateTransitGatewayRouteTableInput, arg3 ...func(*ec2.Options)) (*ec2.AssociateTransitGatewayRouteTableOutput, error) {
	return m.AssociateTransitGatewayRouteTableMock(arg1, arg2, arg3...)
}

// AssociateTrunkInterface mocked function
func (m EC2ClientMock) AssociateTrunkInterface(arg1 context.Context, arg2 *ec2.AssociateTrunkInterfaceInput, arg3 ...func(*ec2.Options)) (*ec2.AssociateTrunkInterfaceOutput, error) {
	return m.AssociateTrunkInterfaceMock(arg1, arg2, arg3...)
}

// AssociateVpcCidrBlock mocked function
func (m EC2ClientMock) AssociateVpcCidrBlock(arg1 context.Context, arg2 *ec2.AssociateVpcCidrBlockInput, arg3 ...func(*ec2.Options)) (*ec2.AssociateVpcCidrBlockOutput, error) {
	return m.AssociateVpcCidrBlockMock(arg1, arg2, arg3...)
}

// AttachClassicLinkVpc mocked function
func (m EC2ClientMock) AttachClassicLinkVpc(arg1 context.Context, arg2 *ec2.AttachClassicLinkVpcInput, arg3 ...func(*ec2.Options)) (*ec2.AttachClassicLinkVpcOutput, error) {
	return m.AttachClassicLinkVpcMock(arg1, arg2, arg3...)
}

// AttachInternetGateway mocked function
func (m EC2ClientMock) AttachInternetGateway(arg1 context.Context, arg2 *ec2.AttachInternetGatewayInput, arg3 ...func(*ec2.Options)) (*ec2.AttachInternetGatewayOutput, error) {
	return m.AttachInternetGatewayMock(arg1, arg2, arg3...)
}

// AttachNetworkInterface mocked function
func (m EC2ClientMock) AttachNetworkInterface(arg1 context.Context, arg2 *ec2.AttachNetworkInterfaceInput, arg3 ...func(*ec2.Options)) (*ec2.AttachNetworkInterfaceOutput, error) {
	return m.AttachNetworkInterfaceMock(arg1, arg2, arg3...)
}

// AttachVolume mocked function
func (m EC2ClientMock) AttachVolume(arg1 context.Context, arg2 *ec2.AttachVolumeInput, arg3 ...func(*ec2.Options)) (*ec2.AttachVolumeOutput, error) {
	return m.AttachVolumeMock(arg1, arg2, arg3...)
}

// AttachVpnGateway mocked function
func (m EC2ClientMock) AttachVpnGateway(arg1 context.Context, arg2 *ec2.AttachVpnGatewayInput, arg3 ...func(*ec2.Options)) (*ec2.AttachVpnGatewayOutput, error) {
	return m.AttachVpnGatewayMock(arg1, arg2, arg3...)
}

// AuthorizeClientVpnIngress mocked function
func (m EC2ClientMock) AuthorizeClientVpnIngress(arg1 context.Context, arg2 *ec2.AuthorizeClientVpnIngressInput, arg3 ...func(*ec2.Options)) (*ec2.AuthorizeClientVpnIngressOutput, error) {
	return m.AuthorizeClientVpnIngressMock(arg1, arg2, arg3...)
}

// AuthorizeSecurityGroupEgress mocked function
func (m EC2ClientMock) AuthorizeSecurityGroupEgress(arg1 context.Context, arg2 *ec2.AuthorizeSecurityGroupEgressInput, arg3 ...func(*ec2.Options)) (*ec2.AuthorizeSecurityGroupEgressOutput, error) {
	return m.AuthorizeSecurityGroupEgressMock(arg1, arg2, arg3...)
}

// AuthorizeSecurityGroupIngress mocked function
func (m EC2ClientMock) AuthorizeSecurityGroupIngress(arg1 context.Context, arg2 *ec2.AuthorizeSecurityGroupIngressInput, arg3 ...func(*ec2.Options)) (*ec2.AuthorizeSecurityGroupIngressOutput, error) {
	return m.AuthorizeSecurityGroupIngressMock(arg1, arg2, arg3...)
}

// BundleInstance mocked function
func (m EC2ClientMock) BundleInstance(arg1 context.Context, arg2 *ec2.BundleInstanceInput, arg3 ...func(*ec2.Options)) (*ec2.BundleInstanceOutput, error) {
	return m.BundleInstanceMock(arg1, arg2, arg3...)
}

// CancelBundleTask mocked function
func (m EC2ClientMock) CancelBundleTask(arg1 context.Context, arg2 *ec2.CancelBundleTaskInput, arg3 ...func(*ec2.Options)) (*ec2.CancelBundleTaskOutput, error) {
	return m.CancelBundleTaskMock(arg1, arg2, arg3...)
}

// CancelCapacityReservation mocked function
func (m EC2ClientMock) CancelCapacityReservation(arg1 context.Context, arg2 *ec2.CancelCapacityReservationInput, arg3 ...func(*ec2.Options)) (*ec2.CancelCapacityReservationOutput, error) {
	return m.CancelCapacityReservationMock(arg1, arg2, arg3...)
}

// CancelCapacityReservationFleets mocked function
func (m EC2ClientMock) CancelCapacityReservationFleets(arg1 context.Context, arg2 *ec2.CancelCapacityReservationFleetsInput, arg3 ...func(*ec2.Options)) (*ec2.CancelCapacityReservationFleetsOutput, error) {
	return m.CancelCapacityReservationFleetsMock(arg1, arg2, arg3...)
}

// CancelConversionTask mocked function
func (m EC2ClientMock) CancelConversionTask(arg1 context.Context, arg2 *ec2.CancelConversionTaskInput, arg3 ...func(*ec2.Options)) (*ec2.CancelConversionTaskOutput, error) {
	return m.CancelConversionTaskMock(arg1, arg2, arg3...)
}

// CancelExportTask mocked function
func (m EC2ClientMock) CancelExportTask(arg1 context.Context, arg2 *ec2.CancelExportTaskInput, arg3 ...func(*ec2.Options)) (*ec2.CancelExportTaskOutput, error) {
	return m.CancelExportTaskMock(arg1, arg2, arg3...)
}

// CancelImportTask mocked function
func (m EC2ClientMock) CancelImportTask(arg1 context.Context, arg2 *ec2.CancelImportTaskInput, arg3 ...func(*ec2.Options)) (*ec2.CancelImportTaskOutput, error) {
	return m.CancelImportTaskMock(arg1, arg2, arg3...)
}

// CancelReservedInstancesListing mocked function
func (m EC2ClientMock) CancelReservedInstancesListing(arg1 context.Context, arg2 *ec2.CancelReservedInstancesListingInput, arg3 ...func(*ec2.Options)) (*ec2.CancelReservedInstancesListingOutput, error) {
	return m.CancelReservedInstancesListingMock(arg1, arg2, arg3...)
}

// CancelSpotFleetRequests mocked function
func (m EC2ClientMock) CancelSpotFleetRequests(arg1 context.Context, arg2 *ec2.CancelSpotFleetRequestsInput, arg3 ...func(*ec2.Options)) (*ec2.CancelSpotFleetRequestsOutput, error) {
	return m.CancelSpotFleetRequestsMock(arg1, arg2, arg3...)
}

// CancelSpotInstanceRequests mocked function
func (m EC2ClientMock) CancelSpotInstanceRequests(arg1 context.Context, arg2 *ec2.CancelSpotInstanceRequestsInput, arg3 ...func(*ec2.Options)) (*ec2.CancelSpotInstanceRequestsOutput, error) {
	return m.CancelSpotInstanceRequestsMock(arg1, arg2, arg3...)
}

// ConfirmProductInstance mocked function
func (m EC2ClientMock) ConfirmProductInstance(arg1 context.Context, arg2 *ec2.ConfirmProductInstanceInput, arg3 ...func(*ec2.Options)) (*ec2.ConfirmProductInstanceOutput, error) {
	return m.ConfirmProductInstanceMock(arg1, arg2, arg3...)
}

// CopyFpgaImage mocked function
func (m EC2ClientMock) CopyFpgaImage(arg1 context.Context, arg2 *ec2.CopyFpgaImageInput, arg3 ...func(*ec2.Options)) (*ec2.CopyFpgaImageOutput, error) {
	return m.CopyFpgaImageMock(arg1, arg2, arg3...)
}

// CopyImage mocked function
func (m EC2ClientMock) CopyImage(arg1 context.Context, arg2 *ec2.CopyImageInput, arg3 ...func(*ec2.Options)) (*ec2.CopyImageOutput, error) {
	return m.CopyImageMock(arg1, arg2, arg3...)
}

// CopySnapshot mocked function
func (m EC2ClientMock) CopySnapshot(arg1 context.Context, arg2 *ec2.CopySnapshotInput, arg3 ...func(*ec2.Options)) (*ec2.CopySnapshotOutput, error) {
	return m.CopySnapshotMock(arg1, arg2, arg3...)
}

// CreateCapacityReservation mocked function
func (m EC2ClientMock) CreateCapacityReservation(arg1 context.Context, arg2 *ec2.CreateCapacityReservationInput, arg3 ...func(*ec2.Options)) (*ec2.CreateCapacityReservationOutput, error) {
	return m.CreateCapacityReservationMock(arg1, arg2, arg3...)
}

// CreateCapacityReservationFleet mocked function
func (m EC2ClientMock) CreateCapacityReservationFleet(arg1 context.Context, arg2 *ec2.CreateCapacityReservationFleetInput, arg3 ...func(*ec2.Options)) (*ec2.CreateCapacityReservationFleetOutput, error) {
	return m.CreateCapacityReservationFleetMock(arg1, arg2, arg3...)
}

// CreateCarrierGateway mocked function
func (m EC2ClientMock) CreateCarrierGateway(arg1 context.Context, arg2 *ec2.CreateCarrierGatewayInput, arg3 ...func(*ec2.Options)) (*ec2.CreateCarrierGatewayOutput, error) {
	return m.CreateCarrierGatewayMock(arg1, arg2, arg3...)
}

// CreateClientVpnEndpoint mocked function
func (m EC2ClientMock) CreateClientVpnEndpoint(arg1 context.Context, arg2 *ec2.CreateClientVpnEndpointInput, arg3 ...func(*ec2.Options)) (*ec2.CreateClientVpnEndpointOutput, error) {
	return m.CreateClientVpnEndpointMock(arg1, arg2, arg3...)
}

// CreateClientVpnRoute mocked function
func (m EC2ClientMock) CreateClientVpnRoute(arg1 context.Context, arg2 *ec2.CreateClientVpnRouteInput, arg3 ...func(*ec2.Options)) (*ec2.CreateClientVpnRouteOutput, error) {
	return m.CreateClientVpnRouteMock(arg1, arg2, arg3...)
}

// CreateCustomerGateway mocked function
func (m EC2ClientMock) CreateCustomerGateway(arg1 context.Context, arg2 *ec2.CreateCustomerGatewayInput, arg3 ...func(*ec2.Options)) (*ec2.CreateCustomerGatewayOutput, error) {
	return m.CreateCustomerGatewayMock(arg1, arg2, arg3...)
}

// CreateDefaultSubnet mocked function
func (m EC2ClientMock) CreateDefaultSubnet(arg1 context.Context, arg2 *ec2.CreateDefaultSubnetInput, arg3 ...func(*ec2.Options)) (*ec2.CreateDefaultSubnetOutput, error) {
	return m.CreateDefaultSubnetMock(arg1, arg2, arg3...)
}

// CreateDefaultVpc mocked function
func (m EC2ClientMock) CreateDefaultVpc(arg1 context.Context, arg2 *ec2.CreateDefaultVpcInput, arg3 ...func(*ec2.Options)) (*ec2.CreateDefaultVpcOutput, error) {
	return m.CreateDefaultVpcMock(arg1, arg2, arg3...)
}

// CreateDhcpOptions mocked function
func (m EC2ClientMock) CreateDhcpOptions(arg1 context.Context, arg2 *ec2.CreateDhcpOptionsInput, arg3 ...func(*ec2.Options)) (*ec2.CreateDhcpOptionsOutput, error) {
	return m.CreateDhcpOptionsMock(arg1, arg2, arg3...)
}

// CreateEgressOnlyInternetGateway mocked function
func (m EC2ClientMock) CreateEgressOnlyInternetGateway(arg1 context.Context, arg2 *ec2.CreateEgressOnlyInternetGatewayInput, arg3 ...func(*ec2.Options)) (*ec2.CreateEgressOnlyInternetGatewayOutput, error) {
	return m.CreateEgressOnlyInternetGatewayMock(arg1, arg2, arg3...)
}

// CreateFleet mocked function
func (m EC2ClientMock) CreateFleet(arg1 context.Context, arg2 *ec2.CreateFleetInput, arg3 ...func(*ec2.Options)) (*ec2.CreateFleetOutput, error) {
	return m.CreateFleetMock(arg1, arg2, arg3...)
}

// CreateFlowLogs mocked function
func (m EC2ClientMock) CreateFlowLogs(arg1 context.Context, arg2 *ec2.CreateFlowLogsInput, arg3 ...func(*ec2.Options)) (*ec2.CreateFlowLogsOutput, error) {
	return m.CreateFlowLogsMock(arg1, arg2, arg3...)
}

// CreateFpgaImage mocked function
func (m EC2ClientMock) CreateFpgaImage(arg1 context.Context, arg2 *ec2.CreateFpgaImageInput, arg3 ...func(*ec2.Options)) (*ec2.CreateFpgaImageOutput, error) {
	return m.CreateFpgaImageMock(arg1, arg2, arg3...)
}

// CreateImage mocked function
func (m EC2ClientMock) CreateImage(arg1 context.Context, arg2 *ec2.CreateImageInput, arg3 ...func(*ec2.Options)) (*ec2.CreateImageOutput, error) {
	return m.CreateImageMock(arg1, arg2, arg3...)
}

// CreateInstanceEventWindow mocked function
func (m EC2ClientMock) CreateInstanceEventWindow(arg1 context.Context, arg2 *ec2.CreateInstanceEventWindowInput, arg3 ...func(*ec2.Options)) (*ec2.CreateInstanceEventWindowOutput, error) {
	return m.CreateInstanceEventWindowMock(arg1, arg2, arg3...)
}

// CreateInstanceExportTask mocked function
func (m EC2ClientMock) CreateInstanceExportTask(arg1 context.Context, arg2 *ec2.CreateInstanceExportTaskInput, arg3 ...func(*ec2.Options)) (*ec2.CreateInstanceExportTaskOutput, error) {
	return m.CreateInstanceExportTaskMock(arg1, arg2, arg3...)
}

// CreateInternetGateway mocked function
func (m EC2ClientMock) CreateInternetGateway(arg1 context.Context, arg2 *ec2.CreateInternetGatewayInput, arg3 ...func(*ec2.Options)) (*ec2.CreateInternetGatewayOutput, error) {
	return m.CreateInternetGatewayMock(arg1, arg2, arg3...)
}

// CreateIpam mocked function
func (m EC2ClientMock) CreateIpam(arg1 context.Context, arg2 *ec2.CreateIpamInput, arg3 ...func(*ec2.Options)) (*ec2.CreateIpamOutput, error) {
	return m.CreateIpamMock(arg1, arg2, arg3...)
}

// CreateIpamPool mocked function
func (m EC2ClientMock) CreateIpamPool(arg1 context.Context, arg2 *ec2.CreateIpamPoolInput, arg3 ...func(*ec2.Options)) (*ec2.CreateIpamPoolOutput, error) {
	return m.CreateIpamPoolMock(arg1, arg2, arg3...)
}

// CreateIpamScope mocked function
func (m EC2ClientMock) CreateIpamScope(arg1 context.Context, arg2 *ec2.CreateIpamScopeInput, arg3 ...func(*ec2.Options)) (*ec2.CreateIpamScopeOutput, error) {
	return m.CreateIpamScopeMock(arg1, arg2, arg3...)
}

// CreateKeyPair mocked function
func (m EC2ClientMock) CreateKeyPair(arg1 context.Context, arg2 *ec2.CreateKeyPairInput, arg3 ...func(*ec2.Options)) (*ec2.CreateKeyPairOutput, error) {
	return m.CreateKeyPairMock(arg1, arg2, arg3...)
}

// CreateLaunchTemplate mocked function
func (m EC2ClientMock) CreateLaunchTemplate(arg1 context.Context, arg2 *ec2.CreateLaunchTemplateInput, arg3 ...func(*ec2.Options)) (*ec2.CreateLaunchTemplateOutput, error) {
	return m.CreateLaunchTemplateMock(arg1, arg2, arg3...)
}

// CreateLaunchTemplateVersion mocked function
func (m EC2ClientMock) CreateLaunchTemplateVersion(arg1 context.Context, arg2 *ec2.CreateLaunchTemplateVersionInput, arg3 ...func(*ec2.Options)) (*ec2.CreateLaunchTemplateVersionOutput, error) {
	return m.CreateLaunchTemplateVersionMock(arg1, arg2, arg3...)
}

// CreateLocalGatewayRoute mocked function
func (m EC2ClientMock) CreateLocalGatewayRoute(arg1 context.Context, arg2 *ec2.CreateLocalGatewayRouteInput, arg3 ...func(*ec2.Options)) (*ec2.CreateLocalGatewayRouteOutput, error) {
	return m.CreateLocalGatewayRouteMock(arg1, arg2, arg3...)
}

// CreateLocalGatewayRouteTableVpcAssociation mocked function
func (m EC2ClientMock) CreateLocalGatewayRouteTableVpcAssociation(arg1 context.Context, arg2 *ec2.CreateLocalGatewayRouteTableVpcAssociationInput, arg3 ...func(*ec2.Options)) (*ec2.CreateLocalGatewayRouteTableVpcAssociationOutput, error) {
	return m.CreateLocalGatewayRouteTableVpcAssociationMock(arg1, arg2, arg3...)
}

// CreateManagedPrefixList mocked function
func (m EC2ClientMock) CreateManagedPrefixList(arg1 context.Context, arg2 *ec2.CreateManagedPrefixListInput, arg3 ...func(*ec2.Options)) (*ec2.CreateManagedPrefixListOutput, error) {
	return m.CreateManagedPrefixListMock(arg1, arg2, arg3...)
}

// CreateNatGateway mocked function
func (m EC2ClientMock) CreateNatGateway(arg1 context.Context, arg2 *ec2.CreateNatGatewayInput, arg3 ...func(*ec2.Options)) (*ec2.CreateNatGatewayOutput, error) {
	return m.CreateNatGatewayMock(arg1, arg2, arg3...)
}

// CreateNetworkAcl mocked function
func (m EC2ClientMock) CreateNetworkAcl(arg1 context.Context, arg2 *ec2.CreateNetworkAclInput, arg3 ...func(*ec2.Options)) (*ec2.CreateNetworkAclOutput, error) {
	return m.CreateNetworkAclMock(arg1, arg2, arg3...)
}

// CreateNetworkAclEntry mocked function
func (m EC2ClientMock) CreateNetworkAclEntry(arg1 context.Context, arg2 *ec2.CreateNetworkAclEntryInput, arg3 ...func(*ec2.Options)) (*ec2.CreateNetworkAclEntryOutput, error) {
	return m.CreateNetworkAclEntryMock(arg1, arg2, arg3...)
}

// CreateNetworkInsightsAccessScope mocked function
func (m EC2ClientMock) CreateNetworkInsightsAccessScope(arg1 context.Context, arg2 *ec2.CreateNetworkInsightsAccessScopeInput, arg3 ...func(*ec2.Options)) (*ec2.CreateNetworkInsightsAccessScopeOutput, error) {
	return m.CreateNetworkInsightsAccessScopeMock(arg1, arg2, arg3...)
}

// CreateNetworkInsightsPath mocked function
func (m EC2ClientMock) CreateNetworkInsightsPath(arg1 context.Context, arg2 *ec2.CreateNetworkInsightsPathInput, arg3 ...func(*ec2.Options)) (*ec2.CreateNetworkInsightsPathOutput, error) {
	return m.CreateNetworkInsightsPathMock(arg1, arg2, arg3...)
}

// CreateNetworkInterface mocked function
func (m EC2ClientMock) CreateNetworkInterface(arg1 context.Context, arg2 *ec2.CreateNetworkInterfaceInput, arg3 ...func(*ec2.Options)) (*ec2.CreateNetworkInterfaceOutput, error) {
	return m.CreateNetworkInterfaceMock(arg1, arg2, arg3...)
}

// CreateNetworkInterfacePermission mocked function
func (m EC2ClientMock) CreateNetworkInterfacePermission(arg1 context.Context, arg2 *ec2.CreateNetworkInterfacePermissionInput, arg3 ...func(*ec2.Options)) (*ec2.CreateNetworkInterfacePermissionOutput, error) {
	return m.CreateNetworkInterfacePermissionMock(arg1, arg2, arg3...)
}

// CreatePlacementGroup mocked function
func (m EC2ClientMock) CreatePlacementGroup(arg1 context.Context, arg2 *ec2.CreatePlacementGroupInput, arg3 ...func(*ec2.Options)) (*ec2.CreatePlacementGroupOutput, error) {
	return m.CreatePlacementGroupMock(arg1, arg2, arg3...)
}

// CreatePublicIpv4Pool mocked function
func (m EC2ClientMock) CreatePublicIpv4Pool(arg1 context.Context, arg2 *ec2.CreatePublicIpv4PoolInput, arg3 ...func(*ec2.Options)) (*ec2.CreatePublicIpv4PoolOutput, error) {
	return m.CreatePublicIpv4PoolMock(arg1, arg2, arg3...)
}

// CreateReplaceRootVolumeTask mocked function
func (m EC2ClientMock) CreateReplaceRootVolumeTask(arg1 context.Context, arg2 *ec2.CreateReplaceRootVolumeTaskInput, arg3 ...func(*ec2.Options)) (*ec2.CreateReplaceRootVolumeTaskOutput, error) {
	return m.CreateReplaceRootVolumeTaskMock(arg1, arg2, arg3...)
}

// CreateReservedInstancesListing mocked function
func (m EC2ClientMock) CreateReservedInstancesListing(arg1 context.Context, arg2 *ec2.CreateReservedInstancesListingInput, arg3 ...func(*ec2.Options)) (*ec2.CreateReservedInstancesListingOutput, error) {
	return m.CreateReservedInstancesListingMock(arg1, arg2, arg3...)
}

// CreateRestoreImageTask mocked function
func (m EC2ClientMock) CreateRestoreImageTask(arg1 context.Context, arg2 *ec2.CreateRestoreImageTaskInput, arg3 ...func(*ec2.Options)) (*ec2.CreateRestoreImageTaskOutput, error) {
	return m.CreateRestoreImageTaskMock(arg1, arg2, arg3...)
}

// CreateRoute mocked function
func (m EC2ClientMock) CreateRoute(arg1 context.Context, arg2 *ec2.CreateRouteInput, arg3 ...func(*ec2.Options)) (*ec2.CreateRouteOutput, error) {
	return m.CreateRouteMock(arg1, arg2, arg3...)
}

// CreateRouteTable mocked function
func (m EC2ClientMock) CreateRouteTable(arg1 context.Context, arg2 *ec2.CreateRouteTableInput, arg3 ...func(*ec2.Options)) (*ec2.CreateRouteTableOutput, error) {
	return m.CreateRouteTableMock(arg1, arg2, arg3...)
}

// CreateSecurityGroup mocked function
func (m EC2ClientMock) CreateSecurityGroup(arg1 context.Context, arg2 *ec2.CreateSecurityGroupInput, arg3 ...func(*ec2.Options)) (*ec2.CreateSecurityGroupOutput, error) {
	return m.CreateSecurityGroupMock(arg1, arg2, arg3...)
}

// CreateSnapshot mocked function
func (m EC2ClientMock) CreateSnapshot(arg1 context.Context, arg2 *ec2.CreateSnapshotInput, arg3 ...func(*ec2.Options)) (*ec2.CreateSnapshotOutput, error) {
	return m.CreateSnapshotMock(arg1, arg2, arg3...)
}

// CreateSnapshots mocked function
func (m EC2ClientMock) CreateSnapshots(arg1 context.Context, arg2 *ec2.CreateSnapshotsInput, arg3 ...func(*ec2.Options)) (*ec2.CreateSnapshotsOutput, error) {
	return m.CreateSnapshotsMock(arg1, arg2, arg3...)
}

// CreateSpotDatafeedSubscription mocked function
func (m EC2ClientMock) CreateSpotDatafeedSubscription(arg1 context.Context, arg2 *ec2.CreateSpotDatafeedSubscriptionInput, arg3 ...func(*ec2.Options)) (*ec2.CreateSpotDatafeedSubscriptionOutput, error) {
	return m.CreateSpotDatafeedSubscriptionMock(arg1, arg2, arg3...)
}

// CreateStoreImageTask mocked function
func (m EC2ClientMock) CreateStoreImageTask(arg1 context.Context, arg2 *ec2.CreateStoreImageTaskInput, arg3 ...func(*ec2.Options)) (*ec2.CreateStoreImageTaskOutput, error) {
	return m.CreateStoreImageTaskMock(arg1, arg2, arg3...)
}

// CreateSubnet mocked function
func (m EC2ClientMock) CreateSubnet(arg1 context.Context, arg2 *ec2.CreateSubnetInput, arg3 ...func(*ec2.Options)) (*ec2.CreateSubnetOutput, error) {
	return m.CreateSubnetMock(arg1, arg2, arg3...)
}

// CreateSubnetCidrReservation mocked function
func (m EC2ClientMock) CreateSubnetCidrReservation(arg1 context.Context, arg2 *ec2.CreateSubnetCidrReservationInput, arg3 ...func(*ec2.Options)) (*ec2.CreateSubnetCidrReservationOutput, error) {
	return m.CreateSubnetCidrReservationMock(arg1, arg2, arg3...)
}

// CreateTags mocked function
func (m EC2ClientMock) CreateTags(arg1 context.Context, arg2 *ec2.CreateTagsInput, arg3 ...func(*ec2.Options)) (*ec2.CreateTagsOutput, error) {
	return m.CreateTagsMock(arg1, arg2, arg3...)
}

// CreateTrafficMirrorFilter mocked function
func (m EC2ClientMock) CreateTrafficMirrorFilter(arg1 context.Context, arg2 *ec2.CreateTrafficMirrorFilterInput, arg3 ...func(*ec2.Options)) (*ec2.CreateTrafficMirrorFilterOutput, error) {
	return m.CreateTrafficMirrorFilterMock(arg1, arg2, arg3...)
}

// CreateTrafficMirrorFilterRule mocked function
func (m EC2ClientMock) CreateTrafficMirrorFilterRule(arg1 context.Context, arg2 *ec2.CreateTrafficMirrorFilterRuleInput, arg3 ...func(*ec2.Options)) (*ec2.CreateTrafficMirrorFilterRuleOutput, error) {
	return m.CreateTrafficMirrorFilterRuleMock(arg1, arg2, arg3...)
}

// CreateTrafficMirrorSession mocked function
func (m EC2ClientMock) CreateTrafficMirrorSession(arg1 context.Context, arg2 *ec2.CreateTrafficMirrorSessionInput, arg3 ...func(*ec2.Options)) (*ec2.CreateTrafficMirrorSessionOutput, error) {
	return m.CreateTrafficMirrorSessionMock(arg1, arg2, arg3...)
}

// CreateTrafficMirrorTarget mocked function
func (m EC2ClientMock) CreateTrafficMirrorTarget(arg1 context.Context, arg2 *ec2.CreateTrafficMirrorTargetInput, arg3 ...func(*ec2.Options)) (*ec2.CreateTrafficMirrorTargetOutput, error) {
	return m.CreateTrafficMirrorTargetMock(arg1, arg2, arg3...)
}

// CreateTransitGateway mocked function
func (m EC2ClientMock) CreateTransitGateway(arg1 context.Context, arg2 *ec2.CreateTransitGatewayInput, arg3 ...func(*ec2.Options)) (*ec2.CreateTransitGatewayOutput, error) {
	return m.CreateTransitGatewayMock(arg1, arg2, arg3...)
}

// CreateTransitGatewayConnect mocked function
func (m EC2ClientMock) CreateTransitGatewayConnect(arg1 context.Context, arg2 *ec2.CreateTransitGatewayConnectInput, arg3 ...func(*ec2.Options)) (*ec2.CreateTransitGatewayConnectOutput, error) {
	return m.CreateTransitGatewayConnectMock(arg1, arg2, arg3...)
}

// CreateTransitGatewayConnectPeer mocked function
func (m EC2ClientMock) CreateTransitGatewayConnectPeer(arg1 context.Context, arg2 *ec2.CreateTransitGatewayConnectPeerInput, arg3 ...func(*ec2.Options)) (*ec2.CreateTransitGatewayConnectPeerOutput, error) {
	return m.CreateTransitGatewayConnectPeerMock(arg1, arg2, arg3...)
}

// CreateTransitGatewayMulticastDomain mocked function
func (m EC2ClientMock) CreateTransitGatewayMulticastDomain(arg1 context.Context, arg2 *ec2.CreateTransitGatewayMulticastDomainInput, arg3 ...func(*ec2.Options)) (*ec2.CreateTransitGatewayMulticastDomainOutput, error) {
	return m.CreateTransitGatewayMulticastDomainMock(arg1, arg2, arg3...)
}

// CreateTransitGatewayPeeringAttachment mocked function
func (m EC2ClientMock) CreateTransitGatewayPeeringAttachment(arg1 context.Context, arg2 *ec2.CreateTransitGatewayPeeringAttachmentInput, arg3 ...func(*ec2.Options)) (*ec2.CreateTransitGatewayPeeringAttachmentOutput, error) {
	return m.CreateTransitGatewayPeeringAttachmentMock(arg1, arg2, arg3...)
}

// CreateTransitGatewayPolicyTable mocked function
func (m EC2ClientMock) CreateTransitGatewayPolicyTable(arg1 context.Context, arg2 *ec2.CreateTransitGatewayPolicyTableInput, arg3 ...func(*ec2.Options)) (*ec2.CreateTransitGatewayPolicyTableOutput, error) {
	return m.CreateTransitGatewayPolicyTableMock(arg1, arg2, arg3...)
}

// CreateTransitGatewayPrefixListReference mocked function
func (m EC2ClientMock) CreateTransitGatewayPrefixListReference(arg1 context.Context, arg2 *ec2.CreateTransitGatewayPrefixListReferenceInput, arg3 ...func(*ec2.Options)) (*ec2.CreateTransitGatewayPrefixListReferenceOutput, error) {
	return m.CreateTransitGatewayPrefixListReferenceMock(arg1, arg2, arg3...)
}

// CreateTransitGatewayRoute mocked function
func (m EC2ClientMock) CreateTransitGatewayRoute(arg1 context.Context, arg2 *ec2.CreateTransitGatewayRouteInput, arg3 ...func(*ec2.Options)) (*ec2.CreateTransitGatewayRouteOutput, error) {
	return m.CreateTransitGatewayRouteMock(arg1, arg2, arg3...)
}

// CreateTransitGatewayRouteTable mocked function
func (m EC2ClientMock) CreateTransitGatewayRouteTable(arg1 context.Context, arg2 *ec2.CreateTransitGatewayRouteTableInput, arg3 ...func(*ec2.Options)) (*ec2.CreateTransitGatewayRouteTableOutput, error) {
	return m.CreateTransitGatewayRouteTableMock(arg1, arg2, arg3...)
}

// CreateTransitGatewayRouteTableAnnouncement mocked function
func (m EC2ClientMock) CreateTransitGatewayRouteTableAnnouncement(arg1 context.Context, arg2 *ec2.CreateTransitGatewayRouteTableAnnouncementInput, arg3 ...func(*ec2.Options)) (*ec2.CreateTransitGatewayRouteTableAnnouncementOutput, error) {
	return m.CreateTransitGatewayRouteTableAnnouncementMock(arg1, arg2, arg3...)
}

// CreateTransitGatewayVpcAttachment mocked function
func (m EC2ClientMock) CreateTransitGatewayVpcAttachment(arg1 context.Context, arg2 *ec2.CreateTransitGatewayVpcAttachmentInput, arg3 ...func(*ec2.Options)) (*ec2.CreateTransitGatewayVpcAttachmentOutput, error) {
	return m.CreateTransitGatewayVpcAttachmentMock(arg1, arg2, arg3...)
}

// CreateVolume mocked function
func (m EC2ClientMock) CreateVolume(arg1 context.Context, arg2 *ec2.CreateVolumeInput, arg3 ...func(*ec2.Options)) (*ec2.CreateVolumeOutput, error) {
	return m.CreateVolumeMock(arg1, arg2, arg3...)
}

// CreateVpc mocked function
func (m EC2ClientMock) CreateVpc(arg1 context.Context, arg2 *ec2.CreateVpcInput, arg3 ...func(*ec2.Options)) (*ec2.CreateVpcOutput, error) {
	return m.CreateVpcMock(arg1, arg2, arg3...)
}

// CreateVpcEndpoint mocked function
func (m EC2ClientMock) CreateVpcEndpoint(arg1 context.Context, arg2 *ec2.CreateVpcEndpointInput, arg3 ...func(*ec2.Options)) (*ec2.CreateVpcEndpointOutput, error) {
	return m.CreateVpcEndpointMock(arg1, arg2, arg3...)
}

// CreateVpcEndpointConnectionNotification mocked function
func (m EC2ClientMock) CreateVpcEndpointConnectionNotification(arg1 context.Context, arg2 *ec2.CreateVpcEndpointConnectionNotificationInput, arg3 ...func(*ec2.Options)) (*ec2.CreateVpcEndpointConnectionNotificationOutput, error) {
	return m.CreateVpcEndpointConnectionNotificationMock(arg1, arg2, arg3...)
}

// CreateVpcEndpointServiceConfiguration mocked function
func (m EC2ClientMock) CreateVpcEndpointServiceConfiguration(arg1 context.Context, arg2 *ec2.CreateVpcEndpointServiceConfigurationInput, arg3 ...func(*ec2.Options)) (*ec2.CreateVpcEndpointServiceConfigurationOutput, error) {
	return m.CreateVpcEndpointServiceConfigurationMock(arg1, arg2, arg3...)
}

// CreateVpcPeeringConnection mocked function
func (m EC2ClientMock) CreateVpcPeeringConnection(arg1 context.Context, arg2 *ec2.CreateVpcPeeringConnectionInput, arg3 ...func(*ec2.Options)) (*ec2.CreateVpcPeeringConnectionOutput, error) {
	return m.CreateVpcPeeringConnectionMock(arg1, arg2, arg3...)
}

// CreateVpnConnection mocked function
func (m EC2ClientMock) CreateVpnConnection(arg1 context.Context, arg2 *ec2.CreateVpnConnectionInput, arg3 ...func(*ec2.Options)) (*ec2.CreateVpnConnectionOutput, error) {
	return m.CreateVpnConnectionMock(arg1, arg2, arg3...)
}

// CreateVpnConnectionRoute mocked function
func (m EC2ClientMock) CreateVpnConnectionRoute(arg1 context.Context, arg2 *ec2.CreateVpnConnectionRouteInput, arg3 ...func(*ec2.Options)) (*ec2.CreateVpnConnectionRouteOutput, error) {
	return m.CreateVpnConnectionRouteMock(arg1, arg2, arg3...)
}

// CreateVpnGateway mocked function
func (m EC2ClientMock) CreateVpnGateway(arg1 context.Context, arg2 *ec2.CreateVpnGatewayInput, arg3 ...func(*ec2.Options)) (*ec2.CreateVpnGatewayOutput, error) {
	return m.CreateVpnGatewayMock(arg1, arg2, arg3...)
}

// DeleteCarrierGateway mocked function
func (m EC2ClientMock) DeleteCarrierGateway(arg1 context.Context, arg2 *ec2.DeleteCarrierGatewayInput, arg3 ...func(*ec2.Options)) (*ec2.DeleteCarrierGatewayOutput, error) {
	return m.DeleteCarrierGatewayMock(arg1, arg2, arg3...)
}

// DeleteClientVpnEndpoint mocked function
func (m EC2ClientMock) DeleteClientVpnEndpoint(arg1 context.Context, arg2 *ec2.DeleteClientVpnEndpointInput, arg3 ...func(*ec2.Options)) (*ec2.DeleteClientVpnEndpointOutput, error) {
	return m.DeleteClientVpnEndpointMock(arg1, arg2, arg3...)
}

// DeleteClientVpnRoute mocked function
func (m EC2ClientMock) DeleteClientVpnRoute(arg1 context.Context, arg2 *ec2.DeleteClientVpnRouteInput, arg3 ...func(*ec2.Options)) (*ec2.DeleteClientVpnRouteOutput, error) {
	return m.DeleteClientVpnRouteMock(arg1, arg2, arg3...)
}

// DeleteCustomerGateway mocked function
func (m EC2ClientMock) DeleteCustomerGateway(arg1 context.Context, arg2 *ec2.DeleteCustomerGatewayInput, arg3 ...func(*ec2.Options)) (*ec2.DeleteCustomerGatewayOutput, error) {
	return m.DeleteCustomerGatewayMock(arg1, arg2, arg3...)
}

// DeleteDhcpOptions mocked function
func (m EC2ClientMock) DeleteDhcpOptions(arg1 context.Context, arg2 *ec2.DeleteDhcpOptionsInput, arg3 ...func(*ec2.Options)) (*ec2.DeleteDhcpOptionsOutput, error) {
	return m.DeleteDhcpOptionsMock(arg1, arg2, arg3...)
}

// DeleteEgressOnlyInternetGateway mocked function
func (m EC2ClientMock) DeleteEgressOnlyInternetGateway(arg1 context.Context, arg2 *ec2.DeleteEgressOnlyInternetGatewayInput, arg3 ...func(*ec2.Options)) (*ec2.DeleteEgressOnlyInternetGatewayOutput, error) {
	return m.DeleteEgressOnlyInternetGatewayMock(arg1, arg2, arg3...)
}

// DeleteFleets mocked function
func (m EC2ClientMock) DeleteFleets(arg1 context.Context, arg2 *ec2.DeleteFleetsInput, arg3 ...func(*ec2.Options)) (*ec2.DeleteFleetsOutput, error) {
	return m.DeleteFleetsMock(arg1, arg2, arg3...)
}

// DeleteFlowLogs mocked function
func (m EC2ClientMock) DeleteFlowLogs(arg1 context.Context, arg2 *ec2.DeleteFlowLogsInput, arg3 ...func(*ec2.Options)) (*ec2.DeleteFlowLogsOutput, error) {
	return m.DeleteFlowLogsMock(arg1, arg2, arg3...)
}

// DeleteFpgaImage mocked function
func (m EC2ClientMock) DeleteFpgaImage(arg1 context.Context, arg2 *ec2.DeleteFpgaImageInput, arg3 ...func(*ec2.Options)) (*ec2.DeleteFpgaImageOutput, error) {
	return m.DeleteFpgaImageMock(arg1, arg2, arg3...)
}

// DeleteInstanceEventWindow mocked function
func (m EC2ClientMock) DeleteInstanceEventWindow(arg1 context.Context, arg2 *ec2.DeleteInstanceEventWindowInput, arg3 ...func(*ec2.Options)) (*ec2.DeleteInstanceEventWindowOutput, error) {
	return m.DeleteInstanceEventWindowMock(arg1, arg2, arg3...)
}

// DeleteInternetGateway mocked function
func (m EC2ClientMock) DeleteInternetGateway(arg1 context.Context, arg2 *ec2.DeleteInternetGatewayInput, arg3 ...func(*ec2.Options)) (*ec2.DeleteInternetGatewayOutput, error) {
	return m.DeleteInternetGatewayMock(arg1, arg2, arg3...)
}

// DeleteIpam mocked function
func (m EC2ClientMock) DeleteIpam(arg1 context.Context, arg2 *ec2.DeleteIpamInput, arg3 ...func(*ec2.Options)) (*ec2.DeleteIpamOutput, error) {
	return m.DeleteIpamMock(arg1, arg2, arg3...)
}

// DeleteIpamPool mocked function
func (m EC2ClientMock) DeleteIpamPool(arg1 context.Context, arg2 *ec2.DeleteIpamPoolInput, arg3 ...func(*ec2.Options)) (*ec2.DeleteIpamPoolOutput, error) {
	return m.DeleteIpamPoolMock(arg1, arg2, arg3...)
}

// DeleteIpamScope mocked function
func (m EC2ClientMock) DeleteIpamScope(arg1 context.Context, arg2 *ec2.DeleteIpamScopeInput, arg3 ...func(*ec2.Options)) (*ec2.DeleteIpamScopeOutput, error) {
	return m.DeleteIpamScopeMock(arg1, arg2, arg3...)
}

// DeleteKeyPair mocked function
func (m EC2ClientMock) DeleteKeyPair(arg1 context.Context, arg2 *ec2.DeleteKeyPairInput, arg3 ...func(*ec2.Options)) (*ec2.DeleteKeyPairOutput, error) {
	return m.DeleteKeyPairMock(arg1, arg2, arg3...)
}

// DeleteLaunchTemplate mocked function
func (m EC2ClientMock) DeleteLaunchTemplate(arg1 context.Context, arg2 *ec2.DeleteLaunchTemplateInput, arg3 ...func(*ec2.Options)) (*ec2.DeleteLaunchTemplateOutput, error) {
	return m.DeleteLaunchTemplateMock(arg1, arg2, arg3...)
}

// DeleteLaunchTemplateVersions mocked function
func (m EC2ClientMock) DeleteLaunchTemplateVersions(arg1 context.Context, arg2 *ec2.DeleteLaunchTemplateVersionsInput, arg3 ...func(*ec2.Options)) (*ec2.DeleteLaunchTemplateVersionsOutput, error) {
	return m.DeleteLaunchTemplateVersionsMock(arg1, arg2, arg3...)
}

// DeleteLocalGatewayRoute mocked function
func (m EC2ClientMock) DeleteLocalGatewayRoute(arg1 context.Context, arg2 *ec2.DeleteLocalGatewayRouteInput, arg3 ...func(*ec2.Options)) (*ec2.DeleteLocalGatewayRouteOutput, error) {
	return m.DeleteLocalGatewayRouteMock(arg1, arg2, arg3...)
}

// DeleteLocalGatewayRouteTableVpcAssociation mocked function
func (m EC2ClientMock) DeleteLocalGatewayRouteTableVpcAssociation(arg1 context.Context, arg2 *ec2.DeleteLocalGatewayRouteTableVpcAssociationInput, arg3 ...func(*ec2.Options)) (*ec2.DeleteLocalGatewayRouteTableVpcAssociationOutput, error) {
	return m.DeleteLocalGatewayRouteTableVpcAssociationMock(arg1, arg2, arg3...)
}

// DeleteManagedPrefixList mocked function
func (m EC2ClientMock) DeleteManagedPrefixList(arg1 context.Context, arg2 *ec2.DeleteManagedPrefixListInput, arg3 ...func(*ec2.Options)) (*ec2.DeleteManagedPrefixListOutput, error) {
	return m.DeleteManagedPrefixListMock(arg1, arg2, arg3...)
}

// DeleteNatGateway mocked function
func (m EC2ClientMock) DeleteNatGateway(arg1 context.Context, arg2 *ec2.DeleteNatGatewayInput, arg3 ...func(*ec2.Options)) (*ec2.DeleteNatGatewayOutput, error) {
	return m.DeleteNatGatewayMock(arg1, arg2, arg3...)
}

// DeleteNetworkAcl mocked function
func (m EC2ClientMock) DeleteNetworkAcl(arg1 context.Context, arg2 *ec2.DeleteNetworkAclInput, arg3 ...func(*ec2.Options)) (*ec2.DeleteNetworkAclOutput, error) {
	return m.DeleteNetworkAclMock(arg1, arg2, arg3...)
}

// DeleteNetworkAclEntry mocked function
func (m EC2ClientMock) DeleteNetworkAclEntry(arg1 context.Context, arg2 *ec2.DeleteNetworkAclEntryInput, arg3 ...func(*ec2.Options)) (*ec2.DeleteNetworkAclEntryOutput, error) {
	return m.DeleteNetworkAclEntryMock(arg1, arg2, arg3...)
}

// DeleteNetworkInsightsAccessScope mocked function
func (m EC2ClientMock) DeleteNetworkInsightsAccessScope(arg1 context.Context, arg2 *ec2.DeleteNetworkInsightsAccessScopeInput, arg3 ...func(*ec2.Options)) (*ec2.DeleteNetworkInsightsAccessScopeOutput, error) {
	return m.DeleteNetworkInsightsAccessScopeMock(arg1, arg2, arg3...)
}

// DeleteNetworkInsightsAccessScopeAnalysis mocked function
func (m EC2ClientMock) DeleteNetworkInsightsAccessScopeAnalysis(arg1 context.Context, arg2 *ec2.DeleteNetworkInsightsAccessScopeAnalysisInput, arg3 ...func(*ec2.Options)) (*ec2.DeleteNetworkInsightsAccessScopeAnalysisOutput, error) {
	return m.DeleteNetworkInsightsAccessScopeAnalysisMock(arg1, arg2, arg3...)
}

// DeleteNetworkInsightsAnalysis mocked function
func (m EC2ClientMock) DeleteNetworkInsightsAnalysis(arg1 context.Context, arg2 *ec2.DeleteNetworkInsightsAnalysisInput, arg3 ...func(*ec2.Options)) (*ec2.DeleteNetworkInsightsAnalysisOutput, error) {
	return m.DeleteNetworkInsightsAnalysisMock(arg1, arg2, arg3...)
}

// DeleteNetworkInsightsPath mocked function
func (m EC2ClientMock) DeleteNetworkInsightsPath(arg1 context.Context, arg2 *ec2.DeleteNetworkInsightsPathInput, arg3 ...func(*ec2.Options)) (*ec2.DeleteNetworkInsightsPathOutput, error) {
	return m.DeleteNetworkInsightsPathMock(arg1, arg2, arg3...)
}

// DeleteNetworkInterface mocked function
func (m EC2ClientMock) DeleteNetworkInterface(arg1 context.Context, arg2 *ec2.DeleteNetworkInterfaceInput, arg3 ...func(*ec2.Options)) (*ec2.DeleteNetworkInterfaceOutput, error) {
	return m.DeleteNetworkInterfaceMock(arg1, arg2, arg3...)
}

// DeleteNetworkInterfacePermission mocked function
func (m EC2ClientMock) DeleteNetworkInterfacePermission(arg1 context.Context, arg2 *ec2.DeleteNetworkInterfacePermissionInput, arg3 ...func(*ec2.Options)) (*ec2.DeleteNetworkInterfacePermissionOutput, error) {
	return m.DeleteNetworkInterfacePermissionMock(arg1, arg2, arg3...)
}

// DeletePlacementGroup mocked function
func (m EC2ClientMock) DeletePlacementGroup(arg1 context.Context, arg2 *ec2.DeletePlacementGroupInput, arg3 ...func(*ec2.Options)) (*ec2.DeletePlacementGroupOutput, error) {
	return m.DeletePlacementGroupMock(arg1, arg2, arg3...)
}

// DeletePublicIpv4Pool mocked function
func (m EC2ClientMock) DeletePublicIpv4Pool(arg1 context.Context, arg2 *ec2.DeletePublicIpv4PoolInput, arg3 ...func(*ec2.Options)) (*ec2.DeletePublicIpv4PoolOutput, error) {
	return m.DeletePublicIpv4PoolMock(arg1, arg2, arg3...)
}

// DeleteQueuedReservedInstances mocked function
func (m EC2ClientMock) DeleteQueuedReservedInstances(arg1 context.Context, arg2 *ec2.DeleteQueuedReservedInstancesInput, arg3 ...func(*ec2.Options)) (*ec2.DeleteQueuedReservedInstancesOutput, error) {
	return m.DeleteQueuedReservedInstancesMock(arg1, arg2, arg3...)
}

// DeleteRoute mocked function
func (m EC2ClientMock) DeleteRoute(arg1 context.Context, arg2 *ec2.DeleteRouteInput, arg3 ...func(*ec2.Options)) (*ec2.DeleteRouteOutput, error) {
	return m.DeleteRouteMock(arg1, arg2, arg3...)
}

// DeleteRouteTable mocked function
func (m EC2ClientMock) DeleteRouteTable(arg1 context.Context, arg2 *ec2.DeleteRouteTableInput, arg3 ...func(*ec2.Options)) (*ec2.DeleteRouteTableOutput, error) {
	return m.DeleteRouteTableMock(arg1, arg2, arg3...)
}

// DeleteSecurityGroup mocked function
func (m EC2ClientMock) DeleteSecurityGroup(arg1 context.Context, arg2 *ec2.DeleteSecurityGroupInput, arg3 ...func(*ec2.Options)) (*ec2.DeleteSecurityGroupOutput, error) {
	return m.DeleteSecurityGroupMock(arg1, arg2, arg3...)
}

// DeleteSnapshot mocked function
func (m EC2ClientMock) DeleteSnapshot(arg1 context.Context, arg2 *ec2.DeleteSnapshotInput, arg3 ...func(*ec2.Options)) (*ec2.DeleteSnapshotOutput, error) {
	return m.DeleteSnapshotMock(arg1, arg2, arg3...)
}

// DeleteSpotDatafeedSubscription mocked function
func (m EC2ClientMock) DeleteSpotDatafeedSubscription(arg1 context.Context, arg2 *ec2.DeleteSpotDatafeedSubscriptionInput, arg3 ...func(*ec2.Options)) (*ec2.DeleteSpotDatafeedSubscriptionOutput, error) {
	return m.DeleteSpotDatafeedSubscriptionMock(arg1, arg2, arg3...)
}

// DeleteSubnet mocked function
func (m EC2ClientMock) DeleteSubnet(arg1 context.Context, arg2 *ec2.DeleteSubnetInput, arg3 ...func(*ec2.Options)) (*ec2.DeleteSubnetOutput, error) {
	return m.DeleteSubnetMock(arg1, arg2, arg3...)
}

// DeleteSubnetCidrReservation mocked function
func (m EC2ClientMock) DeleteSubnetCidrReservation(arg1 context.Context, arg2 *ec2.DeleteSubnetCidrReservationInput, arg3 ...func(*ec2.Options)) (*ec2.DeleteSubnetCidrReservationOutput, error) {
	return m.DeleteSubnetCidrReservationMock(arg1, arg2, arg3...)
}

// DeleteTags mocked function
func (m EC2ClientMock) DeleteTags(arg1 context.Context, arg2 *ec2.DeleteTagsInput, arg3 ...func(*ec2.Options)) (*ec2.DeleteTagsOutput, error) {
	return m.DeleteTagsMock(arg1, arg2, arg3...)
}

// DeleteTrafficMirrorFilter mocked function
func (m EC2ClientMock) DeleteTrafficMirrorFilter(arg1 context.Context, arg2 *ec2.DeleteTrafficMirrorFilterInput, arg3 ...func(*ec2.Options)) (*ec2.DeleteTrafficMirrorFilterOutput, error) {
	return m.DeleteTrafficMirrorFilterMock(arg1, arg2, arg3...)
}

// DeleteTrafficMirrorFilterRule mocked function
func (m EC2ClientMock) DeleteTrafficMirrorFilterRule(arg1 context.Context, arg2 *ec2.DeleteTrafficMirrorFilterRuleInput, arg3 ...func(*ec2.Options)) (*ec2.DeleteTrafficMirrorFilterRuleOutput, error) {
	return m.DeleteTrafficMirrorFilterRuleMock(arg1, arg2, arg3...)
}

// DeleteTrafficMirrorSession mocked function
func (m EC2ClientMock) DeleteTrafficMirrorSession(arg1 context.Context, arg2 *ec2.DeleteTrafficMirrorSessionInput, arg3 ...func(*ec2.Options)) (*ec2.DeleteTrafficMirrorSessionOutput, error) {
	return m.DeleteTrafficMirrorSessionMock(arg1, arg2, arg3...)
}

// DeleteTrafficMirrorTarget mocked function
func (m EC2ClientMock) DeleteTrafficMirrorTarget(arg1 context.Context, arg2 *ec2.DeleteTrafficMirrorTargetInput, arg3 ...func(*ec2.Options)) (*ec2.DeleteTrafficMirrorTargetOutput, error) {
	return m.DeleteTrafficMirrorTargetMock(arg1, arg2, arg3...)
}

// DeleteTransitGateway mocked function
func (m EC2ClientMock) DeleteTransitGateway(arg1 context.Context, arg2 *ec2.DeleteTransitGatewayInput, arg3 ...func(*ec2.Options)) (*ec2.DeleteTransitGatewayOutput, error) {
	return m.DeleteTransitGatewayMock(arg1, arg2, arg3...)
}

// DeleteTransitGatewayConnect mocked function
func (m EC2ClientMock) DeleteTransitGatewayConnect(arg1 context.Context, arg2 *ec2.DeleteTransitGatewayConnectInput, arg3 ...func(*ec2.Options)) (*ec2.DeleteTransitGatewayConnectOutput, error) {
	return m.DeleteTransitGatewayConnectMock(arg1, arg2, arg3...)
}

// DeleteTransitGatewayConnectPeer mocked function
func (m EC2ClientMock) DeleteTransitGatewayConnectPeer(arg1 context.Context, arg2 *ec2.DeleteTransitGatewayConnectPeerInput, arg3 ...func(*ec2.Options)) (*ec2.DeleteTransitGatewayConnectPeerOutput, error) {
	return m.DeleteTransitGatewayConnectPeerMock(arg1, arg2, arg3...)
}

// DeleteTransitGatewayMulticastDomain mocked function
func (m EC2ClientMock) DeleteTransitGatewayMulticastDomain(arg1 context.Context, arg2 *ec2.DeleteTransitGatewayMulticastDomainInput, arg3 ...func(*ec2.Options)) (*ec2.DeleteTransitGatewayMulticastDomainOutput, error) {
	return m.DeleteTransitGatewayMulticastDomainMock(arg1, arg2, arg3...)
}

// DeleteTransitGatewayPeeringAttachment mocked function
func (m EC2ClientMock) DeleteTransitGatewayPeeringAttachment(arg1 context.Context, arg2 *ec2.DeleteTransitGatewayPeeringAttachmentInput, arg3 ...func(*ec2.Options)) (*ec2.DeleteTransitGatewayPeeringAttachmentOutput, error) {
	return m.DeleteTransitGatewayPeeringAttachmentMock(arg1, arg2, arg3...)
}

// DeleteTransitGatewayPolicyTable mocked function
func (m EC2ClientMock) DeleteTransitGatewayPolicyTable(arg1 context.Context, arg2 *ec2.DeleteTransitGatewayPolicyTableInput, arg3 ...func(*ec2.Options)) (*ec2.DeleteTransitGatewayPolicyTableOutput, error) {
	return m.DeleteTransitGatewayPolicyTableMock(arg1, arg2, arg3...)
}

// DeleteTransitGatewayPrefixListReference mocked function
func (m EC2ClientMock) DeleteTransitGatewayPrefixListReference(arg1 context.Context, arg2 *ec2.DeleteTransitGatewayPrefixListReferenceInput, arg3 ...func(*ec2.Options)) (*ec2.DeleteTransitGatewayPrefixListReferenceOutput, error) {
	return m.DeleteTransitGatewayPrefixListReferenceMock(arg1, arg2, arg3...)
}

// DeleteTransitGatewayRoute mocked function
func (m EC2ClientMock) DeleteTransitGatewayRoute(arg1 context.Context, arg2 *ec2.DeleteTransitGatewayRouteInput, arg3 ...func(*ec2.Options)) (*ec2.DeleteTransitGatewayRouteOutput, error) {
	return m.DeleteTransitGatewayRouteMock(arg1, arg2, arg3...)
}

// DeleteTransitGatewayRouteTable mocked function
func (m EC2ClientMock) DeleteTransitGatewayRouteTable(arg1 context.Context, arg2 *ec2.DeleteTransitGatewayRouteTableInput, arg3 ...func(*ec2.Options)) (*ec2.DeleteTransitGatewayRouteTableOutput, error) {
	return m.DeleteTransitGatewayRouteTableMock(arg1, arg2, arg3...)
}

// DeleteTransitGatewayRouteTableAnnouncement mocked function
func (m EC2ClientMock) DeleteTransitGatewayRouteTableAnnouncement(arg1 context.Context, arg2 *ec2.DeleteTransitGatewayRouteTableAnnouncementInput, arg3 ...func(*ec2.Options)) (*ec2.DeleteTransitGatewayRouteTableAnnouncementOutput, error) {
	return m.DeleteTransitGatewayRouteTableAnnouncementMock(arg1, arg2, arg3...)
}

// DeleteTransitGatewayVpcAttachment mocked function
func (m EC2ClientMock) DeleteTransitGatewayVpcAttachment(arg1 context.Context, arg2 *ec2.DeleteTransitGatewayVpcAttachmentInput, arg3 ...func(*ec2.Options)) (*ec2.DeleteTransitGatewayVpcAttachmentOutput, error) {
	return m.DeleteTransitGatewayVpcAttachmentMock(arg1, arg2, arg3...)
}

// DeleteVolume mocked function
func (m EC2ClientMock) DeleteVolume(arg1 context.Context, arg2 *ec2.DeleteVolumeInput, arg3 ...func(*ec2.Options)) (*ec2.DeleteVolumeOutput, error) {
	return m.DeleteVolumeMock(arg1, arg2, arg3...)
}

// DeleteVpc mocked function
func (m EC2ClientMock) DeleteVpc(arg1 context.Context, arg2 *ec2.DeleteVpcInput, arg3 ...func(*ec2.Options)) (*ec2.DeleteVpcOutput, error) {
	return m.DeleteVpcMock(arg1, arg2, arg3...)
}

// DeleteVpcEndpointConnectionNotifications mocked function
func (m EC2ClientMock) DeleteVpcEndpointConnectionNotifications(arg1 context.Context, arg2 *ec2.DeleteVpcEndpointConnectionNotificationsInput, arg3 ...func(*ec2.Options)) (*ec2.DeleteVpcEndpointConnectionNotificationsOutput, error) {
	return m.DeleteVpcEndpointConnectionNotificationsMock(arg1, arg2, arg3...)
}

// DeleteVpcEndpointServiceConfigurations mocked function
func (m EC2ClientMock) DeleteVpcEndpointServiceConfigurations(arg1 context.Context, arg2 *ec2.DeleteVpcEndpointServiceConfigurationsInput, arg3 ...func(*ec2.Options)) (*ec2.DeleteVpcEndpointServiceConfigurationsOutput, error) {
	return m.DeleteVpcEndpointServiceConfigurationsMock(arg1, arg2, arg3...)
}

// DeleteVpcEndpoints mocked function
func (m EC2ClientMock) DeleteVpcEndpoints(arg1 context.Context, arg2 *ec2.DeleteVpcEndpointsInput, arg3 ...func(*ec2.Options)) (*ec2.DeleteVpcEndpointsOutput, error) {
	return m.DeleteVpcEndpointsMock(arg1, arg2, arg3...)
}

// DeleteVpcPeeringConnection mocked function
func (m EC2ClientMock) DeleteVpcPeeringConnection(arg1 context.Context, arg2 *ec2.DeleteVpcPeeringConnectionInput, arg3 ...func(*ec2.Options)) (*ec2.DeleteVpcPeeringConnectionOutput, error) {
	return m.DeleteVpcPeeringConnectionMock(arg1, arg2, arg3...)
}

// DeleteVpnConnection mocked function
func (m EC2ClientMock) DeleteVpnConnection(arg1 context.Context, arg2 *ec2.DeleteVpnConnectionInput, arg3 ...func(*ec2.Options)) (*ec2.DeleteVpnConnectionOutput, error) {
	return m.DeleteVpnConnectionMock(arg1, arg2, arg3...)
}

// DeleteVpnConnectionRoute mocked function
func (m EC2ClientMock) DeleteVpnConnectionRoute(arg1 context.Context, arg2 *ec2.DeleteVpnConnectionRouteInput, arg3 ...func(*ec2.Options)) (*ec2.DeleteVpnConnectionRouteOutput, error) {
	return m.DeleteVpnConnectionRouteMock(arg1, arg2, arg3...)
}

// DeleteVpnGateway mocked function
func (m EC2ClientMock) DeleteVpnGateway(arg1 context.Context, arg2 *ec2.DeleteVpnGatewayInput, arg3 ...func(*ec2.Options)) (*ec2.DeleteVpnGatewayOutput, error) {
	return m.DeleteVpnGatewayMock(arg1, arg2, arg3...)
}

// DeprovisionByoipCidr mocked function
func (m EC2ClientMock) DeprovisionByoipCidr(arg1 context.Context, arg2 *ec2.DeprovisionByoipCidrInput, arg3 ...func(*ec2.Options)) (*ec2.DeprovisionByoipCidrOutput, error) {
	return m.DeprovisionByoipCidrMock(arg1, arg2, arg3...)
}

// DeprovisionIpamPoolCidr mocked function
func (m EC2ClientMock) DeprovisionIpamPoolCidr(arg1 context.Context, arg2 *ec2.DeprovisionIpamPoolCidrInput, arg3 ...func(*ec2.Options)) (*ec2.DeprovisionIpamPoolCidrOutput, error) {
	return m.DeprovisionIpamPoolCidrMock(arg1, arg2, arg3...)
}

// DeprovisionPublicIpv4PoolCidr mocked function
func (m EC2ClientMock) DeprovisionPublicIpv4PoolCidr(arg1 context.Context, arg2 *ec2.DeprovisionPublicIpv4PoolCidrInput, arg3 ...func(*ec2.Options)) (*ec2.DeprovisionPublicIpv4PoolCidrOutput, error) {
	return m.DeprovisionPublicIpv4PoolCidrMock(arg1, arg2, arg3...)
}

// DeregisterImage mocked function
func (m EC2ClientMock) DeregisterImage(arg1 context.Context, arg2 *ec2.DeregisterImageInput, arg3 ...func(*ec2.Options)) (*ec2.DeregisterImageOutput, error) {
	return m.DeregisterImageMock(arg1, arg2, arg3...)
}

// DeregisterInstanceEventNotificationAttributes mocked function
func (m EC2ClientMock) DeregisterInstanceEventNotificationAttributes(arg1 context.Context, arg2 *ec2.DeregisterInstanceEventNotificationAttributesInput, arg3 ...func(*ec2.Options)) (*ec2.DeregisterInstanceEventNotificationAttributesOutput, error) {
	return m.DeregisterInstanceEventNotificationAttributesMock(arg1, arg2, arg3...)
}

// DeregisterTransitGatewayMulticastGroupMembers mocked function
func (m EC2ClientMock) DeregisterTransitGatewayMulticastGroupMembers(arg1 context.Context, arg2 *ec2.DeregisterTransitGatewayMulticastGroupMembersInput, arg3 ...func(*ec2.Options)) (*ec2.DeregisterTransitGatewayMulticastGroupMembersOutput, error) {
	return m.DeregisterTransitGatewayMulticastGroupMembersMock(arg1, arg2, arg3...)
}

// DeregisterTransitGatewayMulticastGroupSources mocked function
func (m EC2ClientMock) DeregisterTransitGatewayMulticastGroupSources(arg1 context.Context, arg2 *ec2.DeregisterTransitGatewayMulticastGroupSourcesInput, arg3 ...func(*ec2.Options)) (*ec2.DeregisterTransitGatewayMulticastGroupSourcesOutput, error) {
	return m.DeregisterTransitGatewayMulticastGroupSourcesMock(arg1, arg2, arg3...)
}

// DescribeAccountAttributes mocked function
func (m EC2ClientMock) DescribeAccountAttributes(arg1 context.Context, arg2 *ec2.DescribeAccountAttributesInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeAccountAttributesOutput, error) {
	return m.DescribeAccountAttributesMock(arg1, arg2, arg3...)
}

// DescribeAddresses mocked function
func (m EC2ClientMock) DescribeAddresses(arg1 context.Context, arg2 *ec2.DescribeAddressesInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeAddressesOutput, error) {
	return m.DescribeAddressesMock(arg1, arg2, arg3...)
}

// DescribeAddressesAttribute mocked function
func (m EC2ClientMock) DescribeAddressesAttribute(arg1 context.Context, arg2 *ec2.DescribeAddressesAttributeInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeAddressesAttributeOutput, error) {
	return m.DescribeAddressesAttributeMock(arg1, arg2, arg3...)
}

// DescribeAggregateIdFormat mocked function
func (m EC2ClientMock) DescribeAggregateIdFormat(arg1 context.Context, arg2 *ec2.DescribeAggregateIdFormatInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeAggregateIdFormatOutput, error) {
	return m.DescribeAggregateIdFormatMock(arg1, arg2, arg3...)
}

// DescribeAvailabilityZones mocked function
func (m EC2ClientMock) DescribeAvailabilityZones(arg1 context.Context, arg2 *ec2.DescribeAvailabilityZonesInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeAvailabilityZonesOutput, error) {
	return m.DescribeAvailabilityZonesMock(arg1, arg2, arg3...)
}

// DescribeBundleTasks mocked function
func (m EC2ClientMock) DescribeBundleTasks(arg1 context.Context, arg2 *ec2.DescribeBundleTasksInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeBundleTasksOutput, error) {
	return m.DescribeBundleTasksMock(arg1, arg2, arg3...)
}

// DescribeByoipCidrs mocked function
func (m EC2ClientMock) DescribeByoipCidrs(arg1 context.Context, arg2 *ec2.DescribeByoipCidrsInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeByoipCidrsOutput, error) {
	return m.DescribeByoipCidrsMock(arg1, arg2, arg3...)
}

// DescribeCapacityReservationFleets mocked function
func (m EC2ClientMock) DescribeCapacityReservationFleets(arg1 context.Context, arg2 *ec2.DescribeCapacityReservationFleetsInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeCapacityReservationFleetsOutput, error) {
	return m.DescribeCapacityReservationFleetsMock(arg1, arg2, arg3...)
}

// DescribeCapacityReservations mocked function
func (m EC2ClientMock) DescribeCapacityReservations(arg1 context.Context, arg2 *ec2.DescribeCapacityReservationsInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeCapacityReservationsOutput, error) {
	return m.DescribeCapacityReservationsMock(arg1, arg2, arg3...)
}

// DescribeCarrierGateways mocked function
func (m EC2ClientMock) DescribeCarrierGateways(arg1 context.Context, arg2 *ec2.DescribeCarrierGatewaysInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeCarrierGatewaysOutput, error) {
	return m.DescribeCarrierGatewaysMock(arg1, arg2, arg3...)
}

// DescribeClassicLinkInstances mocked function
func (m EC2ClientMock) DescribeClassicLinkInstances(arg1 context.Context, arg2 *ec2.DescribeClassicLinkInstancesInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeClassicLinkInstancesOutput, error) {
	return m.DescribeClassicLinkInstancesMock(arg1, arg2, arg3...)
}

// DescribeClientVpnAuthorizationRules mocked function
func (m EC2ClientMock) DescribeClientVpnAuthorizationRules(arg1 context.Context, arg2 *ec2.DescribeClientVpnAuthorizationRulesInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeClientVpnAuthorizationRulesOutput, error) {
	return m.DescribeClientVpnAuthorizationRulesMock(arg1, arg2, arg3...)
}

// DescribeClientVpnConnections mocked function
func (m EC2ClientMock) DescribeClientVpnConnections(arg1 context.Context, arg2 *ec2.DescribeClientVpnConnectionsInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeClientVpnConnectionsOutput, error) {
	return m.DescribeClientVpnConnectionsMock(arg1, arg2, arg3...)
}

// DescribeClientVpnEndpoints mocked function
func (m EC2ClientMock) DescribeClientVpnEndpoints(arg1 context.Context, arg2 *ec2.DescribeClientVpnEndpointsInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeClientVpnEndpointsOutput, error) {
	return m.DescribeClientVpnEndpointsMock(arg1, arg2, arg3...)
}

// DescribeClientVpnRoutes mocked function
func (m EC2ClientMock) DescribeClientVpnRoutes(arg1 context.Context, arg2 *ec2.DescribeClientVpnRoutesInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeClientVpnRoutesOutput, error) {
	return m.DescribeClientVpnRoutesMock(arg1, arg2, arg3...)
}

// DescribeClientVpnTargetNetworks mocked function
func (m EC2ClientMock) DescribeClientVpnTargetNetworks(arg1 context.Context, arg2 *ec2.DescribeClientVpnTargetNetworksInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeClientVpnTargetNetworksOutput, error) {
	return m.DescribeClientVpnTargetNetworksMock(arg1, arg2, arg3...)
}

// DescribeCoipPools mocked function
func (m EC2ClientMock) DescribeCoipPools(arg1 context.Context, arg2 *ec2.DescribeCoipPoolsInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeCoipPoolsOutput, error) {
	return m.DescribeCoipPoolsMock(arg1, arg2, arg3...)
}

// DescribeConversionTasks mocked function
func (m EC2ClientMock) DescribeConversionTasks(arg1 context.Context, arg2 *ec2.DescribeConversionTasksInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeConversionTasksOutput, error) {
	return m.DescribeConversionTasksMock(arg1, arg2, arg3...)
}

// DescribeCustomerGateways mocked function
func (m EC2ClientMock) DescribeCustomerGateways(arg1 context.Context, arg2 *ec2.DescribeCustomerGatewaysInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeCustomerGatewaysOutput, error) {
	return m.DescribeCustomerGatewaysMock(arg1, arg2, arg3...)
}

// DescribeDhcpOptions mocked function
func (m EC2ClientMock) DescribeDhcpOptions(arg1 context.Context, arg2 *ec2.DescribeDhcpOptionsInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeDhcpOptionsOutput, error) {
	return m.DescribeDhcpOptionsMock(arg1, arg2, arg3...)
}

// DescribeEgressOnlyInternetGateways mocked function
func (m EC2ClientMock) DescribeEgressOnlyInternetGateways(arg1 context.Context, arg2 *ec2.DescribeEgressOnlyInternetGatewaysInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeEgressOnlyInternetGatewaysOutput, error) {
	return m.DescribeEgressOnlyInternetGatewaysMock(arg1, arg2, arg3...)
}

// DescribeElasticGpus mocked function
func (m EC2ClientMock) DescribeElasticGpus(arg1 context.Context, arg2 *ec2.DescribeElasticGpusInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeElasticGpusOutput, error) {
	return m.DescribeElasticGpusMock(arg1, arg2, arg3...)
}

// DescribeExportImageTasks mocked function
func (m EC2ClientMock) DescribeExportImageTasks(arg1 context.Context, arg2 *ec2.DescribeExportImageTasksInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeExportImageTasksOutput, error) {
	return m.DescribeExportImageTasksMock(arg1, arg2, arg3...)
}

// DescribeExportTasks mocked function
func (m EC2ClientMock) DescribeExportTasks(arg1 context.Context, arg2 *ec2.DescribeExportTasksInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeExportTasksOutput, error) {
	return m.DescribeExportTasksMock(arg1, arg2, arg3...)
}

// DescribeFastLaunchImages mocked function
func (m EC2ClientMock) DescribeFastLaunchImages(arg1 context.Context, arg2 *ec2.DescribeFastLaunchImagesInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeFastLaunchImagesOutput, error) {
	return m.DescribeFastLaunchImagesMock(arg1, arg2, arg3...)
}

// DescribeFastSnapshotRestores mocked function
func (m EC2ClientMock) DescribeFastSnapshotRestores(arg1 context.Context, arg2 *ec2.DescribeFastSnapshotRestoresInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeFastSnapshotRestoresOutput, error) {
	return m.DescribeFastSnapshotRestoresMock(arg1, arg2, arg3...)
}

// DescribeFleetHistory mocked function
func (m EC2ClientMock) DescribeFleetHistory(arg1 context.Context, arg2 *ec2.DescribeFleetHistoryInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeFleetHistoryOutput, error) {
	return m.DescribeFleetHistoryMock(arg1, arg2, arg3...)
}

// DescribeFleetInstances mocked function
func (m EC2ClientMock) DescribeFleetInstances(arg1 context.Context, arg2 *ec2.DescribeFleetInstancesInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeFleetInstancesOutput, error) {
	return m.DescribeFleetInstancesMock(arg1, arg2, arg3...)
}

// DescribeFleets mocked function
func (m EC2ClientMock) DescribeFleets(arg1 context.Context, arg2 *ec2.DescribeFleetsInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeFleetsOutput, error) {
	return m.DescribeFleetsMock(arg1, arg2, arg3...)
}

// DescribeFlowLogs mocked function
func (m EC2ClientMock) DescribeFlowLogs(arg1 context.Context, arg2 *ec2.DescribeFlowLogsInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeFlowLogsOutput, error) {
	return m.DescribeFlowLogsMock(arg1, arg2, arg3...)
}

// DescribeFpgaImageAttribute mocked function
func (m EC2ClientMock) DescribeFpgaImageAttribute(arg1 context.Context, arg2 *ec2.DescribeFpgaImageAttributeInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeFpgaImageAttributeOutput, error) {
	return m.DescribeFpgaImageAttributeMock(arg1, arg2, arg3...)
}

// DescribeFpgaImages mocked function
func (m EC2ClientMock) DescribeFpgaImages(arg1 context.Context, arg2 *ec2.DescribeFpgaImagesInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeFpgaImagesOutput, error) {
	return m.DescribeFpgaImagesMock(arg1, arg2, arg3...)
}

// DescribeHostReservationOfferings mocked function
func (m EC2ClientMock) DescribeHostReservationOfferings(arg1 context.Context, arg2 *ec2.DescribeHostReservationOfferingsInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeHostReservationOfferingsOutput, error) {
	return m.DescribeHostReservationOfferingsMock(arg1, arg2, arg3...)
}

// DescribeHostReservations mocked function
func (m EC2ClientMock) DescribeHostReservations(arg1 context.Context, arg2 *ec2.DescribeHostReservationsInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeHostReservationsOutput, error) {
	return m.DescribeHostReservationsMock(arg1, arg2, arg3...)
}

// DescribeHosts mocked function
func (m EC2ClientMock) DescribeHosts(arg1 context.Context, arg2 *ec2.DescribeHostsInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeHostsOutput, error) {
	return m.DescribeHostsMock(arg1, arg2, arg3...)
}

// DescribeIamInstanceProfileAssociations mocked function
func (m EC2ClientMock) DescribeIamInstanceProfileAssociations(arg1 context.Context, arg2 *ec2.DescribeIamInstanceProfileAssociationsInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeIamInstanceProfileAssociationsOutput, error) {
	return m.DescribeIamInstanceProfileAssociationsMock(arg1, arg2, arg3...)
}

// DescribeIdFormat mocked function
func (m EC2ClientMock) DescribeIdFormat(arg1 context.Context, arg2 *ec2.DescribeIdFormatInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeIdFormatOutput, error) {
	return m.DescribeIdFormatMock(arg1, arg2, arg3...)
}

// DescribeIdentityIdFormat mocked function
func (m EC2ClientMock) DescribeIdentityIdFormat(arg1 context.Context, arg2 *ec2.DescribeIdentityIdFormatInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeIdentityIdFormatOutput, error) {
	return m.DescribeIdentityIdFormatMock(arg1, arg2, arg3...)
}

// DescribeImageAttribute mocked function
func (m EC2ClientMock) DescribeImageAttribute(arg1 context.Context, arg2 *ec2.DescribeImageAttributeInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeImageAttributeOutput, error) {
	return m.DescribeImageAttributeMock(arg1, arg2, arg3...)
}

// DescribeImages mocked function
func (m EC2ClientMock) DescribeImages(arg1 context.Context, arg2 *ec2.DescribeImagesInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeImagesOutput, error) {
	return m.DescribeImagesMock(arg1, arg2, arg3...)
}

// DescribeImportImageTasks mocked function
func (m EC2ClientMock) DescribeImportImageTasks(arg1 context.Context, arg2 *ec2.DescribeImportImageTasksInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeImportImageTasksOutput, error) {
	return m.DescribeImportImageTasksMock(arg1, arg2, arg3...)
}

// DescribeImportSnapshotTasks mocked function
func (m EC2ClientMock) DescribeImportSnapshotTasks(arg1 context.Context, arg2 *ec2.DescribeImportSnapshotTasksInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeImportSnapshotTasksOutput, error) {
	return m.DescribeImportSnapshotTasksMock(arg1, arg2, arg3...)
}

// DescribeInstanceAttribute mocked function
func (m EC2ClientMock) DescribeInstanceAttribute(arg1 context.Context, arg2 *ec2.DescribeInstanceAttributeInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeInstanceAttributeOutput, error) {
	return m.DescribeInstanceAttributeMock(arg1, arg2, arg3...)
}

// DescribeInstanceCreditSpecifications mocked function
func (m EC2ClientMock) DescribeInstanceCreditSpecifications(arg1 context.Context, arg2 *ec2.DescribeInstanceCreditSpecificationsInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeInstanceCreditSpecificationsOutput, error) {
	return m.DescribeInstanceCreditSpecificationsMock(arg1, arg2, arg3...)
}

// DescribeInstanceEventNotificationAttributes mocked function
func (m EC2ClientMock) DescribeInstanceEventNotificationAttributes(arg1 context.Context, arg2 *ec2.DescribeInstanceEventNotificationAttributesInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeInstanceEventNotificationAttributesOutput, error) {
	return m.DescribeInstanceEventNotificationAttributesMock(arg1, arg2, arg3...)
}

// DescribeInstanceEventWindows mocked function
func (m EC2ClientMock) DescribeInstanceEventWindows(arg1 context.Context, arg2 *ec2.DescribeInstanceEventWindowsInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeInstanceEventWindowsOutput, error) {
	return m.DescribeInstanceEventWindowsMock(arg1, arg2, arg3...)
}

// DescribeInstanceStatus mocked function
func (m EC2ClientMock) DescribeInstanceStatus(arg1 context.Context, arg2 *ec2.DescribeInstanceStatusInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeInstanceStatusOutput, error) {
	return m.DescribeInstanceStatusMock(arg1, arg2, arg3...)
}

// DescribeInstanceTypeOfferings mocked function
func (m EC2ClientMock) DescribeInstanceTypeOfferings(arg1 context.Context, arg2 *ec2.DescribeInstanceTypeOfferingsInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeInstanceTypeOfferingsOutput, error) {
	return m.DescribeInstanceTypeOfferingsMock(arg1, arg2, arg3...)
}

// DescribeInstanceTypes mocked function
func (m EC2ClientMock) DescribeInstanceTypes(arg1 context.Context, arg2 *ec2.DescribeInstanceTypesInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeInstanceTypesOutput, error) {
	return m.DescribeInstanceTypesMock(arg1, arg2, arg3...)
}

// DescribeInstances mocked function
func (m EC2ClientMock) DescribeInstances(arg1 context.Context, arg2 *ec2.DescribeInstancesInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeInstancesOutput, error) {
	return m.DescribeInstancesMock(arg1, arg2, arg3...)
}

// DescribeInternetGateways mocked function
func (m EC2ClientMock) DescribeInternetGateways(arg1 context.Context, arg2 *ec2.DescribeInternetGatewaysInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeInternetGatewaysOutput, error) {
	return m.DescribeInternetGatewaysMock(arg1, arg2, arg3...)
}

// DescribeIpamPools mocked function
func (m EC2ClientMock) DescribeIpamPools(arg1 context.Context, arg2 *ec2.DescribeIpamPoolsInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeIpamPoolsOutput, error) {
	return m.DescribeIpamPoolsMock(arg1, arg2, arg3...)
}

// DescribeIpamScopes mocked function
func (m EC2ClientMock) DescribeIpamScopes(arg1 context.Context, arg2 *ec2.DescribeIpamScopesInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeIpamScopesOutput, error) {
	return m.DescribeIpamScopesMock(arg1, arg2, arg3...)
}

// DescribeIpams mocked function
func (m EC2ClientMock) DescribeIpams(arg1 context.Context, arg2 *ec2.DescribeIpamsInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeIpamsOutput, error) {
	return m.DescribeIpamsMock(arg1, arg2, arg3...)
}

// DescribeIpv6Pools mocked function
func (m EC2ClientMock) DescribeIpv6Pools(arg1 context.Context, arg2 *ec2.DescribeIpv6PoolsInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeIpv6PoolsOutput, error) {
	return m.DescribeIpv6PoolsMock(arg1, arg2, arg3...)
}

// DescribeKeyPairs mocked function
func (m EC2ClientMock) DescribeKeyPairs(arg1 context.Context, arg2 *ec2.DescribeKeyPairsInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeKeyPairsOutput, error) {
	return m.DescribeKeyPairsMock(arg1, arg2, arg3...)
}

// DescribeLaunchTemplateVersions mocked function
func (m EC2ClientMock) DescribeLaunchTemplateVersions(arg1 context.Context, arg2 *ec2.DescribeLaunchTemplateVersionsInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeLaunchTemplateVersionsOutput, error) {
	return m.DescribeLaunchTemplateVersionsMock(arg1, arg2, arg3...)
}

// DescribeLaunchTemplates mocked function
func (m EC2ClientMock) DescribeLaunchTemplates(arg1 context.Context, arg2 *ec2.DescribeLaunchTemplatesInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeLaunchTemplatesOutput, error) {
	return m.DescribeLaunchTemplatesMock(arg1, arg2, arg3...)
}

// DescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociations mocked function
func (m EC2ClientMock) DescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociations(arg1 context.Context, arg2 *ec2.DescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociationsInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociationsOutput, error) {
	return m.DescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociationsMock(arg1, arg2, arg3...)
}

// DescribeLocalGatewayRouteTableVpcAssociations mocked function
func (m EC2ClientMock) DescribeLocalGatewayRouteTableVpcAssociations(arg1 context.Context, arg2 *ec2.DescribeLocalGatewayRouteTableVpcAssociationsInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeLocalGatewayRouteTableVpcAssociationsOutput, error) {
	return m.DescribeLocalGatewayRouteTableVpcAssociationsMock(arg1, arg2, arg3...)
}

// DescribeLocalGatewayRouteTables mocked function
func (m EC2ClientMock) DescribeLocalGatewayRouteTables(arg1 context.Context, arg2 *ec2.DescribeLocalGatewayRouteTablesInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeLocalGatewayRouteTablesOutput, error) {
	return m.DescribeLocalGatewayRouteTablesMock(arg1, arg2, arg3...)
}

// DescribeLocalGatewayVirtualInterfaceGroups mocked function
func (m EC2ClientMock) DescribeLocalGatewayVirtualInterfaceGroups(arg1 context.Context, arg2 *ec2.DescribeLocalGatewayVirtualInterfaceGroupsInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeLocalGatewayVirtualInterfaceGroupsOutput, error) {
	return m.DescribeLocalGatewayVirtualInterfaceGroupsMock(arg1, arg2, arg3...)
}

// DescribeLocalGatewayVirtualInterfaces mocked function
func (m EC2ClientMock) DescribeLocalGatewayVirtualInterfaces(arg1 context.Context, arg2 *ec2.DescribeLocalGatewayVirtualInterfacesInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeLocalGatewayVirtualInterfacesOutput, error) {
	return m.DescribeLocalGatewayVirtualInterfacesMock(arg1, arg2, arg3...)
}

// DescribeLocalGateways mocked function
func (m EC2ClientMock) DescribeLocalGateways(arg1 context.Context, arg2 *ec2.DescribeLocalGatewaysInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeLocalGatewaysOutput, error) {
	return m.DescribeLocalGatewaysMock(arg1, arg2, arg3...)
}

// DescribeManagedPrefixLists mocked function
func (m EC2ClientMock) DescribeManagedPrefixLists(arg1 context.Context, arg2 *ec2.DescribeManagedPrefixListsInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeManagedPrefixListsOutput, error) {
	return m.DescribeManagedPrefixListsMock(arg1, arg2, arg3...)
}

// DescribeMovingAddresses mocked function
func (m EC2ClientMock) DescribeMovingAddresses(arg1 context.Context, arg2 *ec2.DescribeMovingAddressesInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeMovingAddressesOutput, error) {
	return m.DescribeMovingAddressesMock(arg1, arg2, arg3...)
}

// DescribeNatGateways mocked function
func (m EC2ClientMock) DescribeNatGateways(arg1 context.Context, arg2 *ec2.DescribeNatGatewaysInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeNatGatewaysOutput, error) {
	return m.DescribeNatGatewaysMock(arg1, arg2, arg3...)
}

// DescribeNetworkAcls mocked function
func (m EC2ClientMock) DescribeNetworkAcls(arg1 context.Context, arg2 *ec2.DescribeNetworkAclsInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeNetworkAclsOutput, error) {
	return m.DescribeNetworkAclsMock(arg1, arg2, arg3...)
}

// DescribeNetworkInsightsAccessScopeAnalyses mocked function
func (m EC2ClientMock) DescribeNetworkInsightsAccessScopeAnalyses(arg1 context.Context, arg2 *ec2.DescribeNetworkInsightsAccessScopeAnalysesInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeNetworkInsightsAccessScopeAnalysesOutput, error) {
	return m.DescribeNetworkInsightsAccessScopeAnalysesMock(arg1, arg2, arg3...)
}

// DescribeNetworkInsightsAccessScopes mocked function
func (m EC2ClientMock) DescribeNetworkInsightsAccessScopes(arg1 context.Context, arg2 *ec2.DescribeNetworkInsightsAccessScopesInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeNetworkInsightsAccessScopesOutput, error) {
	return m.DescribeNetworkInsightsAccessScopesMock(arg1, arg2, arg3...)
}

// DescribeNetworkInsightsAnalyses mocked function
func (m EC2ClientMock) DescribeNetworkInsightsAnalyses(arg1 context.Context, arg2 *ec2.DescribeNetworkInsightsAnalysesInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeNetworkInsightsAnalysesOutput, error) {
	return m.DescribeNetworkInsightsAnalysesMock(arg1, arg2, arg3...)
}

// DescribeNetworkInsightsPaths mocked function
func (m EC2ClientMock) DescribeNetworkInsightsPaths(arg1 context.Context, arg2 *ec2.DescribeNetworkInsightsPathsInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeNetworkInsightsPathsOutput, error) {
	return m.DescribeNetworkInsightsPathsMock(arg1, arg2, arg3...)
}

// DescribeNetworkInterfaceAttribute mocked function
func (m EC2ClientMock) DescribeNetworkInterfaceAttribute(arg1 context.Context, arg2 *ec2.DescribeNetworkInterfaceAttributeInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeNetworkInterfaceAttributeOutput, error) {
	return m.DescribeNetworkInterfaceAttributeMock(arg1, arg2, arg3...)
}

// DescribeNetworkInterfacePermissions mocked function
func (m EC2ClientMock) DescribeNetworkInterfacePermissions(arg1 context.Context, arg2 *ec2.DescribeNetworkInterfacePermissionsInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeNetworkInterfacePermissionsOutput, error) {
	return m.DescribeNetworkInterfacePermissionsMock(arg1, arg2, arg3...)
}

// DescribeNetworkInterfaces mocked function
func (m EC2ClientMock) DescribeNetworkInterfaces(arg1 context.Context, arg2 *ec2.DescribeNetworkInterfacesInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeNetworkInterfacesOutput, error) {
	return m.DescribeNetworkInterfacesMock(arg1, arg2, arg3...)
}

// DescribePlacementGroups mocked function
func (m EC2ClientMock) DescribePlacementGroups(arg1 context.Context, arg2 *ec2.DescribePlacementGroupsInput, arg3 ...func(*ec2.Options)) (*ec2.DescribePlacementGroupsOutput, error) {
	return m.DescribePlacementGroupsMock(arg1, arg2, arg3...)
}

// DescribePrefixLists mocked function
func (m EC2ClientMock) DescribePrefixLists(arg1 context.Context, arg2 *ec2.DescribePrefixListsInput, arg3 ...func(*ec2.Options)) (*ec2.DescribePrefixListsOutput, error) {
	return m.DescribePrefixListsMock(arg1, arg2, arg3...)
}

// DescribePrincipalIdFormat mocked function
func (m EC2ClientMock) DescribePrincipalIdFormat(arg1 context.Context, arg2 *ec2.DescribePrincipalIdFormatInput, arg3 ...func(*ec2.Options)) (*ec2.DescribePrincipalIdFormatOutput, error) {
	return m.DescribePrincipalIdFormatMock(arg1, arg2, arg3...)
}

// DescribePublicIpv4Pools mocked function
func (m EC2ClientMock) DescribePublicIpv4Pools(arg1 context.Context, arg2 *ec2.DescribePublicIpv4PoolsInput, arg3 ...func(*ec2.Options)) (*ec2.DescribePublicIpv4PoolsOutput, error) {
	return m.DescribePublicIpv4PoolsMock(arg1, arg2, arg3...)
}

// DescribeRegions mocked function
func (m EC2ClientMock) DescribeRegions(arg1 context.Context, arg2 *ec2.DescribeRegionsInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeRegionsOutput, error) {
	return m.DescribeRegionsMock(arg1, arg2, arg3...)
}

// DescribeReplaceRootVolumeTasks mocked function
func (m EC2ClientMock) DescribeReplaceRootVolumeTasks(arg1 context.Context, arg2 *ec2.DescribeReplaceRootVolumeTasksInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeReplaceRootVolumeTasksOutput, error) {
	return m.DescribeReplaceRootVolumeTasksMock(arg1, arg2, arg3...)
}

// DescribeReservedInstances mocked function
func (m EC2ClientMock) DescribeReservedInstances(arg1 context.Context, arg2 *ec2.DescribeReservedInstancesInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeReservedInstancesOutput, error) {
	return m.DescribeReservedInstancesMock(arg1, arg2, arg3...)
}

// DescribeReservedInstancesListings mocked function
func (m EC2ClientMock) DescribeReservedInstancesListings(arg1 context.Context, arg2 *ec2.DescribeReservedInstancesListingsInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeReservedInstancesListingsOutput, error) {
	return m.DescribeReservedInstancesListingsMock(arg1, arg2, arg3...)
}

// DescribeReservedInstancesModifications mocked function
func (m EC2ClientMock) DescribeReservedInstancesModifications(arg1 context.Context, arg2 *ec2.DescribeReservedInstancesModificationsInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeReservedInstancesModificationsOutput, error) {
	return m.DescribeReservedInstancesModificationsMock(arg1, arg2, arg3...)
}

// DescribeReservedInstancesOfferings mocked function
func (m EC2ClientMock) DescribeReservedInstancesOfferings(arg1 context.Context, arg2 *ec2.DescribeReservedInstancesOfferingsInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeReservedInstancesOfferingsOutput, error) {
	return m.DescribeReservedInstancesOfferingsMock(arg1, arg2, arg3...)
}

// DescribeRouteTables mocked function
func (m EC2ClientMock) DescribeRouteTables(arg1 context.Context, arg2 *ec2.DescribeRouteTablesInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeRouteTablesOutput, error) {
	return m.DescribeRouteTablesMock(arg1, arg2, arg3...)
}

// DescribeScheduledInstanceAvailability mocked function
func (m EC2ClientMock) DescribeScheduledInstanceAvailability(arg1 context.Context, arg2 *ec2.DescribeScheduledInstanceAvailabilityInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeScheduledInstanceAvailabilityOutput, error) {
	return m.DescribeScheduledInstanceAvailabilityMock(arg1, arg2, arg3...)
}

// DescribeScheduledInstances mocked function
func (m EC2ClientMock) DescribeScheduledInstances(arg1 context.Context, arg2 *ec2.DescribeScheduledInstancesInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeScheduledInstancesOutput, error) {
	return m.DescribeScheduledInstancesMock(arg1, arg2, arg3...)
}

// DescribeSecurityGroupReferences mocked function
func (m EC2ClientMock) DescribeSecurityGroupReferences(arg1 context.Context, arg2 *ec2.DescribeSecurityGroupReferencesInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeSecurityGroupReferencesOutput, error) {
	return m.DescribeSecurityGroupReferencesMock(arg1, arg2, arg3...)
}

// DescribeSecurityGroupRules mocked function
func (m EC2ClientMock) DescribeSecurityGroupRules(arg1 context.Context, arg2 *ec2.DescribeSecurityGroupRulesInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeSecurityGroupRulesOutput, error) {
	return m.DescribeSecurityGroupRulesMock(arg1, arg2, arg3...)
}

// DescribeSecurityGroups mocked function
func (m EC2ClientMock) DescribeSecurityGroups(arg1 context.Context, arg2 *ec2.DescribeSecurityGroupsInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeSecurityGroupsOutput, error) {
	return m.DescribeSecurityGroupsMock(arg1, arg2, arg3...)
}

// DescribeSnapshotAttribute mocked function
func (m EC2ClientMock) DescribeSnapshotAttribute(arg1 context.Context, arg2 *ec2.DescribeSnapshotAttributeInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeSnapshotAttributeOutput, error) {
	return m.DescribeSnapshotAttributeMock(arg1, arg2, arg3...)
}

// DescribeSnapshotTierStatus mocked function
func (m EC2ClientMock) DescribeSnapshotTierStatus(arg1 context.Context, arg2 *ec2.DescribeSnapshotTierStatusInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeSnapshotTierStatusOutput, error) {
	return m.DescribeSnapshotTierStatusMock(arg1, arg2, arg3...)
}

// DescribeSnapshots mocked function
func (m EC2ClientMock) DescribeSnapshots(arg1 context.Context, arg2 *ec2.DescribeSnapshotsInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeSnapshotsOutput, error) {
	return m.DescribeSnapshotsMock(arg1, arg2, arg3...)
}

// DescribeSpotDatafeedSubscription mocked function
func (m EC2ClientMock) DescribeSpotDatafeedSubscription(arg1 context.Context, arg2 *ec2.DescribeSpotDatafeedSubscriptionInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeSpotDatafeedSubscriptionOutput, error) {
	return m.DescribeSpotDatafeedSubscriptionMock(arg1, arg2, arg3...)
}

// DescribeSpotFleetInstances mocked function
func (m EC2ClientMock) DescribeSpotFleetInstances(arg1 context.Context, arg2 *ec2.DescribeSpotFleetInstancesInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeSpotFleetInstancesOutput, error) {
	return m.DescribeSpotFleetInstancesMock(arg1, arg2, arg3...)
}

// DescribeSpotFleetRequestHistory mocked function
func (m EC2ClientMock) DescribeSpotFleetRequestHistory(arg1 context.Context, arg2 *ec2.DescribeSpotFleetRequestHistoryInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeSpotFleetRequestHistoryOutput, error) {
	return m.DescribeSpotFleetRequestHistoryMock(arg1, arg2, arg3...)
}

// DescribeSpotFleetRequests mocked function
func (m EC2ClientMock) DescribeSpotFleetRequests(arg1 context.Context, arg2 *ec2.DescribeSpotFleetRequestsInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeSpotFleetRequestsOutput, error) {
	return m.DescribeSpotFleetRequestsMock(arg1, arg2, arg3...)
}

// DescribeSpotInstanceRequests mocked function
func (m EC2ClientMock) DescribeSpotInstanceRequests(arg1 context.Context, arg2 *ec2.DescribeSpotInstanceRequestsInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeSpotInstanceRequestsOutput, error) {
	return m.DescribeSpotInstanceRequestsMock(arg1, arg2, arg3...)
}

// DescribeSpotPriceHistory mocked function
func (m EC2ClientMock) DescribeSpotPriceHistory(arg1 context.Context, arg2 *ec2.DescribeSpotPriceHistoryInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeSpotPriceHistoryOutput, error) {
	return m.DescribeSpotPriceHistoryMock(arg1, arg2, arg3...)
}

// DescribeStaleSecurityGroups mocked function
func (m EC2ClientMock) DescribeStaleSecurityGroups(arg1 context.Context, arg2 *ec2.DescribeStaleSecurityGroupsInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeStaleSecurityGroupsOutput, error) {
	return m.DescribeStaleSecurityGroupsMock(arg1, arg2, arg3...)
}

// DescribeStoreImageTasks mocked function
func (m EC2ClientMock) DescribeStoreImageTasks(arg1 context.Context, arg2 *ec2.DescribeStoreImageTasksInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeStoreImageTasksOutput, error) {
	return m.DescribeStoreImageTasksMock(arg1, arg2, arg3...)
}

// DescribeSubnets mocked function
func (m EC2ClientMock) DescribeSubnets(arg1 context.Context, arg2 *ec2.DescribeSubnetsInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeSubnetsOutput, error) {
	return m.DescribeSubnetsMock(arg1, arg2, arg3...)
}

// DescribeTags mocked function
func (m EC2ClientMock) DescribeTags(arg1 context.Context, arg2 *ec2.DescribeTagsInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeTagsOutput, error) {
	return m.DescribeTagsMock(arg1, arg2, arg3...)
}

// DescribeTrafficMirrorFilters mocked function
func (m EC2ClientMock) DescribeTrafficMirrorFilters(arg1 context.Context, arg2 *ec2.DescribeTrafficMirrorFiltersInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeTrafficMirrorFiltersOutput, error) {
	return m.DescribeTrafficMirrorFiltersMock(arg1, arg2, arg3...)
}

// DescribeTrafficMirrorSessions mocked function
func (m EC2ClientMock) DescribeTrafficMirrorSessions(arg1 context.Context, arg2 *ec2.DescribeTrafficMirrorSessionsInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeTrafficMirrorSessionsOutput, error) {
	return m.DescribeTrafficMirrorSessionsMock(arg1, arg2, arg3...)
}

// DescribeTrafficMirrorTargets mocked function
func (m EC2ClientMock) DescribeTrafficMirrorTargets(arg1 context.Context, arg2 *ec2.DescribeTrafficMirrorTargetsInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeTrafficMirrorTargetsOutput, error) {
	return m.DescribeTrafficMirrorTargetsMock(arg1, arg2, arg3...)
}

// DescribeTransitGatewayAttachments mocked function
func (m EC2ClientMock) DescribeTransitGatewayAttachments(arg1 context.Context, arg2 *ec2.DescribeTransitGatewayAttachmentsInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeTransitGatewayAttachmentsOutput, error) {
	return m.DescribeTransitGatewayAttachmentsMock(arg1, arg2, arg3...)
}

// DescribeTransitGatewayConnectPeers mocked function
func (m EC2ClientMock) DescribeTransitGatewayConnectPeers(arg1 context.Context, arg2 *ec2.DescribeTransitGatewayConnectPeersInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeTransitGatewayConnectPeersOutput, error) {
	return m.DescribeTransitGatewayConnectPeersMock(arg1, arg2, arg3...)
}

// DescribeTransitGatewayConnects mocked function
func (m EC2ClientMock) DescribeTransitGatewayConnects(arg1 context.Context, arg2 *ec2.DescribeTransitGatewayConnectsInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeTransitGatewayConnectsOutput, error) {
	return m.DescribeTransitGatewayConnectsMock(arg1, arg2, arg3...)
}

// DescribeTransitGatewayMulticastDomains mocked function
func (m EC2ClientMock) DescribeTransitGatewayMulticastDomains(arg1 context.Context, arg2 *ec2.DescribeTransitGatewayMulticastDomainsInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeTransitGatewayMulticastDomainsOutput, error) {
	return m.DescribeTransitGatewayMulticastDomainsMock(arg1, arg2, arg3...)
}

// DescribeTransitGatewayPeeringAttachments mocked function
func (m EC2ClientMock) DescribeTransitGatewayPeeringAttachments(arg1 context.Context, arg2 *ec2.DescribeTransitGatewayPeeringAttachmentsInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeTransitGatewayPeeringAttachmentsOutput, error) {
	return m.DescribeTransitGatewayPeeringAttachmentsMock(arg1, arg2, arg3...)
}

// DescribeTransitGatewayPolicyTables mocked function
func (m EC2ClientMock) DescribeTransitGatewayPolicyTables(arg1 context.Context, arg2 *ec2.DescribeTransitGatewayPolicyTablesInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeTransitGatewayPolicyTablesOutput, error) {
	return m.DescribeTransitGatewayPolicyTablesMock(arg1, arg2, arg3...)
}

// DescribeTransitGatewayRouteTableAnnouncements mocked function
func (m EC2ClientMock) DescribeTransitGatewayRouteTableAnnouncements(arg1 context.Context, arg2 *ec2.DescribeTransitGatewayRouteTableAnnouncementsInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeTransitGatewayRouteTableAnnouncementsOutput, error) {
	return m.DescribeTransitGatewayRouteTableAnnouncementsMock(arg1, arg2, arg3...)
}

// DescribeTransitGatewayRouteTables mocked function
func (m EC2ClientMock) DescribeTransitGatewayRouteTables(arg1 context.Context, arg2 *ec2.DescribeTransitGatewayRouteTablesInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeTransitGatewayRouteTablesOutput, error) {
	return m.DescribeTransitGatewayRouteTablesMock(arg1, arg2, arg3...)
}

// DescribeTransitGatewayVpcAttachments mocked function
func (m EC2ClientMock) DescribeTransitGatewayVpcAttachments(arg1 context.Context, arg2 *ec2.DescribeTransitGatewayVpcAttachmentsInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeTransitGatewayVpcAttachmentsOutput, error) {
	return m.DescribeTransitGatewayVpcAttachmentsMock(arg1, arg2, arg3...)
}

// DescribeTransitGateways mocked function
func (m EC2ClientMock) DescribeTransitGateways(arg1 context.Context, arg2 *ec2.DescribeTransitGatewaysInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeTransitGatewaysOutput, error) {
	return m.DescribeTransitGatewaysMock(arg1, arg2, arg3...)
}

// DescribeTrunkInterfaceAssociations mocked function
func (m EC2ClientMock) DescribeTrunkInterfaceAssociations(arg1 context.Context, arg2 *ec2.DescribeTrunkInterfaceAssociationsInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeTrunkInterfaceAssociationsOutput, error) {
	return m.DescribeTrunkInterfaceAssociationsMock(arg1, arg2, arg3...)
}

// DescribeVolumeAttribute mocked function
func (m EC2ClientMock) DescribeVolumeAttribute(arg1 context.Context, arg2 *ec2.DescribeVolumeAttributeInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeVolumeAttributeOutput, error) {
	return m.DescribeVolumeAttributeMock(arg1, arg2, arg3...)
}

// DescribeVolumeStatus mocked function
func (m EC2ClientMock) DescribeVolumeStatus(arg1 context.Context, arg2 *ec2.DescribeVolumeStatusInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeVolumeStatusOutput, error) {
	return m.DescribeVolumeStatusMock(arg1, arg2, arg3...)
}

// DescribeVolumes mocked function
func (m EC2ClientMock) DescribeVolumes(arg1 context.Context, arg2 *ec2.DescribeVolumesInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeVolumesOutput, error) {
	return m.DescribeVolumesMock(arg1, arg2, arg3...)
}

// DescribeVolumesModifications mocked function
func (m EC2ClientMock) DescribeVolumesModifications(arg1 context.Context, arg2 *ec2.DescribeVolumesModificationsInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeVolumesModificationsOutput, error) {
	return m.DescribeVolumesModificationsMock(arg1, arg2, arg3...)
}

// DescribeVpcAttribute mocked function
func (m EC2ClientMock) DescribeVpcAttribute(arg1 context.Context, arg2 *ec2.DescribeVpcAttributeInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeVpcAttributeOutput, error) {
	return m.DescribeVpcAttributeMock(arg1, arg2, arg3...)
}

// DescribeVpcClassicLink mocked function
func (m EC2ClientMock) DescribeVpcClassicLink(arg1 context.Context, arg2 *ec2.DescribeVpcClassicLinkInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeVpcClassicLinkOutput, error) {
	return m.DescribeVpcClassicLinkMock(arg1, arg2, arg3...)
}

// DescribeVpcClassicLinkDnsSupport mocked function
func (m EC2ClientMock) DescribeVpcClassicLinkDnsSupport(arg1 context.Context, arg2 *ec2.DescribeVpcClassicLinkDnsSupportInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeVpcClassicLinkDnsSupportOutput, error) {
	return m.DescribeVpcClassicLinkDnsSupportMock(arg1, arg2, arg3...)
}

// DescribeVpcEndpointConnectionNotifications mocked function
func (m EC2ClientMock) DescribeVpcEndpointConnectionNotifications(arg1 context.Context, arg2 *ec2.DescribeVpcEndpointConnectionNotificationsInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeVpcEndpointConnectionNotificationsOutput, error) {
	return m.DescribeVpcEndpointConnectionNotificationsMock(arg1, arg2, arg3...)
}

// DescribeVpcEndpointConnections mocked function
func (m EC2ClientMock) DescribeVpcEndpointConnections(arg1 context.Context, arg2 *ec2.DescribeVpcEndpointConnectionsInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeVpcEndpointConnectionsOutput, error) {
	return m.DescribeVpcEndpointConnectionsMock(arg1, arg2, arg3...)
}

// DescribeVpcEndpointServiceConfigurations mocked function
func (m EC2ClientMock) DescribeVpcEndpointServiceConfigurations(arg1 context.Context, arg2 *ec2.DescribeVpcEndpointServiceConfigurationsInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeVpcEndpointServiceConfigurationsOutput, error) {
	return m.DescribeVpcEndpointServiceConfigurationsMock(arg1, arg2, arg3...)
}

// DescribeVpcEndpointServicePermissions mocked function
func (m EC2ClientMock) DescribeVpcEndpointServicePermissions(arg1 context.Context, arg2 *ec2.DescribeVpcEndpointServicePermissionsInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeVpcEndpointServicePermissionsOutput, error) {
	return m.DescribeVpcEndpointServicePermissionsMock(arg1, arg2, arg3...)
}

// DescribeVpcEndpointServices mocked function
func (m EC2ClientMock) DescribeVpcEndpointServices(arg1 context.Context, arg2 *ec2.DescribeVpcEndpointServicesInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeVpcEndpointServicesOutput, error) {
	return m.DescribeVpcEndpointServicesMock(arg1, arg2, arg3...)
}

// DescribeVpcEndpoints mocked function
func (m EC2ClientMock) DescribeVpcEndpoints(arg1 context.Context, arg2 *ec2.DescribeVpcEndpointsInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeVpcEndpointsOutput, error) {
	return m.DescribeVpcEndpointsMock(arg1, arg2, arg3...)
}

// DescribeVpcPeeringConnections mocked function
func (m EC2ClientMock) DescribeVpcPeeringConnections(arg1 context.Context, arg2 *ec2.DescribeVpcPeeringConnectionsInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeVpcPeeringConnectionsOutput, error) {
	return m.DescribeVpcPeeringConnectionsMock(arg1, arg2, arg3...)
}

// DescribeVpcs mocked function
func (m EC2ClientMock) DescribeVpcs(arg1 context.Context, arg2 *ec2.DescribeVpcsInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeVpcsOutput, error) {
	return m.DescribeVpcsMock(arg1, arg2, arg3...)
}

// DescribeVpnConnections mocked function
func (m EC2ClientMock) DescribeVpnConnections(arg1 context.Context, arg2 *ec2.DescribeVpnConnectionsInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeVpnConnectionsOutput, error) {
	return m.DescribeVpnConnectionsMock(arg1, arg2, arg3...)
}

// DescribeVpnGateways mocked function
func (m EC2ClientMock) DescribeVpnGateways(arg1 context.Context, arg2 *ec2.DescribeVpnGatewaysInput, arg3 ...func(*ec2.Options)) (*ec2.DescribeVpnGatewaysOutput, error) {
	return m.DescribeVpnGatewaysMock(arg1, arg2, arg3...)
}

// DetachClassicLinkVpc mocked function
func (m EC2ClientMock) DetachClassicLinkVpc(arg1 context.Context, arg2 *ec2.DetachClassicLinkVpcInput, arg3 ...func(*ec2.Options)) (*ec2.DetachClassicLinkVpcOutput, error) {
	return m.DetachClassicLinkVpcMock(arg1, arg2, arg3...)
}

// DetachInternetGateway mocked function
func (m EC2ClientMock) DetachInternetGateway(arg1 context.Context, arg2 *ec2.DetachInternetGatewayInput, arg3 ...func(*ec2.Options)) (*ec2.DetachInternetGatewayOutput, error) {
	return m.DetachInternetGatewayMock(arg1, arg2, arg3...)
}

// DetachNetworkInterface mocked function
func (m EC2ClientMock) DetachNetworkInterface(arg1 context.Context, arg2 *ec2.DetachNetworkInterfaceInput, arg3 ...func(*ec2.Options)) (*ec2.DetachNetworkInterfaceOutput, error) {
	return m.DetachNetworkInterfaceMock(arg1, arg2, arg3...)
}

// DetachVolume mocked function
func (m EC2ClientMock) DetachVolume(arg1 context.Context, arg2 *ec2.DetachVolumeInput, arg3 ...func(*ec2.Options)) (*ec2.DetachVolumeOutput, error) {
	return m.DetachVolumeMock(arg1, arg2, arg3...)
}

// DetachVpnGateway mocked function
func (m EC2ClientMock) DetachVpnGateway(arg1 context.Context, arg2 *ec2.DetachVpnGatewayInput, arg3 ...func(*ec2.Options)) (*ec2.DetachVpnGatewayOutput, error) {
	return m.DetachVpnGatewayMock(arg1, arg2, arg3...)
}

// DisableEbsEncryptionByDefault mocked function
func (m EC2ClientMock) DisableEbsEncryptionByDefault(arg1 context.Context, arg2 *ec2.DisableEbsEncryptionByDefaultInput, arg3 ...func(*ec2.Options)) (*ec2.DisableEbsEncryptionByDefaultOutput, error) {
	return m.DisableEbsEncryptionByDefaultMock(arg1, arg2, arg3...)
}

// DisableFastLaunch mocked function
func (m EC2ClientMock) DisableFastLaunch(arg1 context.Context, arg2 *ec2.DisableFastLaunchInput, arg3 ...func(*ec2.Options)) (*ec2.DisableFastLaunchOutput, error) {
	return m.DisableFastLaunchMock(arg1, arg2, arg3...)
}

// DisableFastSnapshotRestores mocked function
func (m EC2ClientMock) DisableFastSnapshotRestores(arg1 context.Context, arg2 *ec2.DisableFastSnapshotRestoresInput, arg3 ...func(*ec2.Options)) (*ec2.DisableFastSnapshotRestoresOutput, error) {
	return m.DisableFastSnapshotRestoresMock(arg1, arg2, arg3...)
}

// DisableImageDeprecation mocked function
func (m EC2ClientMock) DisableImageDeprecation(arg1 context.Context, arg2 *ec2.DisableImageDeprecationInput, arg3 ...func(*ec2.Options)) (*ec2.DisableImageDeprecationOutput, error) {
	return m.DisableImageDeprecationMock(arg1, arg2, arg3...)
}

// DisableIpamOrganizationAdminAccount mocked function
func (m EC2ClientMock) DisableIpamOrganizationAdminAccount(arg1 context.Context, arg2 *ec2.DisableIpamOrganizationAdminAccountInput, arg3 ...func(*ec2.Options)) (*ec2.DisableIpamOrganizationAdminAccountOutput, error) {
	return m.DisableIpamOrganizationAdminAccountMock(arg1, arg2, arg3...)
}

// DisableSerialConsoleAccess mocked function
func (m EC2ClientMock) DisableSerialConsoleAccess(arg1 context.Context, arg2 *ec2.DisableSerialConsoleAccessInput, arg3 ...func(*ec2.Options)) (*ec2.DisableSerialConsoleAccessOutput, error) {
	return m.DisableSerialConsoleAccessMock(arg1, arg2, arg3...)
}

// DisableTransitGatewayRouteTablePropagation mocked function
func (m EC2ClientMock) DisableTransitGatewayRouteTablePropagation(arg1 context.Context, arg2 *ec2.DisableTransitGatewayRouteTablePropagationInput, arg3 ...func(*ec2.Options)) (*ec2.DisableTransitGatewayRouteTablePropagationOutput, error) {
	return m.DisableTransitGatewayRouteTablePropagationMock(arg1, arg2, arg3...)
}

// DisableVgwRoutePropagation mocked function
func (m EC2ClientMock) DisableVgwRoutePropagation(arg1 context.Context, arg2 *ec2.DisableVgwRoutePropagationInput, arg3 ...func(*ec2.Options)) (*ec2.DisableVgwRoutePropagationOutput, error) {
	return m.DisableVgwRoutePropagationMock(arg1, arg2, arg3...)
}

// DisableVpcClassicLink mocked function
func (m EC2ClientMock) DisableVpcClassicLink(arg1 context.Context, arg2 *ec2.DisableVpcClassicLinkInput, arg3 ...func(*ec2.Options)) (*ec2.DisableVpcClassicLinkOutput, error) {
	return m.DisableVpcClassicLinkMock(arg1, arg2, arg3...)
}

// DisableVpcClassicLinkDnsSupport mocked function
func (m EC2ClientMock) DisableVpcClassicLinkDnsSupport(arg1 context.Context, arg2 *ec2.DisableVpcClassicLinkDnsSupportInput, arg3 ...func(*ec2.Options)) (*ec2.DisableVpcClassicLinkDnsSupportOutput, error) {
	return m.DisableVpcClassicLinkDnsSupportMock(arg1, arg2, arg3...)
}

// DisassociateAddress mocked function
func (m EC2ClientMock) DisassociateAddress(arg1 context.Context, arg2 *ec2.DisassociateAddressInput, arg3 ...func(*ec2.Options)) (*ec2.DisassociateAddressOutput, error) {
	return m.DisassociateAddressMock(arg1, arg2, arg3...)
}

// DisassociateClientVpnTargetNetwork mocked function
func (m EC2ClientMock) DisassociateClientVpnTargetNetwork(arg1 context.Context, arg2 *ec2.DisassociateClientVpnTargetNetworkInput, arg3 ...func(*ec2.Options)) (*ec2.DisassociateClientVpnTargetNetworkOutput, error) {
	return m.DisassociateClientVpnTargetNetworkMock(arg1, arg2, arg3...)
}

// DisassociateEnclaveCertificateIamRole mocked function
func (m EC2ClientMock) DisassociateEnclaveCertificateIamRole(arg1 context.Context, arg2 *ec2.DisassociateEnclaveCertificateIamRoleInput, arg3 ...func(*ec2.Options)) (*ec2.DisassociateEnclaveCertificateIamRoleOutput, error) {
	return m.DisassociateEnclaveCertificateIamRoleMock(arg1, arg2, arg3...)
}

// DisassociateIamInstanceProfile mocked function
func (m EC2ClientMock) DisassociateIamInstanceProfile(arg1 context.Context, arg2 *ec2.DisassociateIamInstanceProfileInput, arg3 ...func(*ec2.Options)) (*ec2.DisassociateIamInstanceProfileOutput, error) {
	return m.DisassociateIamInstanceProfileMock(arg1, arg2, arg3...)
}

// DisassociateInstanceEventWindow mocked function
func (m EC2ClientMock) DisassociateInstanceEventWindow(arg1 context.Context, arg2 *ec2.DisassociateInstanceEventWindowInput, arg3 ...func(*ec2.Options)) (*ec2.DisassociateInstanceEventWindowOutput, error) {
	return m.DisassociateInstanceEventWindowMock(arg1, arg2, arg3...)
}

// DisassociateRouteTable mocked function
func (m EC2ClientMock) DisassociateRouteTable(arg1 context.Context, arg2 *ec2.DisassociateRouteTableInput, arg3 ...func(*ec2.Options)) (*ec2.DisassociateRouteTableOutput, error) {
	return m.DisassociateRouteTableMock(arg1, arg2, arg3...)
}

// DisassociateSubnetCidrBlock mocked function
func (m EC2ClientMock) DisassociateSubnetCidrBlock(arg1 context.Context, arg2 *ec2.DisassociateSubnetCidrBlockInput, arg3 ...func(*ec2.Options)) (*ec2.DisassociateSubnetCidrBlockOutput, error) {
	return m.DisassociateSubnetCidrBlockMock(arg1, arg2, arg3...)
}

// DisassociateTransitGatewayMulticastDomain mocked function
func (m EC2ClientMock) DisassociateTransitGatewayMulticastDomain(arg1 context.Context, arg2 *ec2.DisassociateTransitGatewayMulticastDomainInput, arg3 ...func(*ec2.Options)) (*ec2.DisassociateTransitGatewayMulticastDomainOutput, error) {
	return m.DisassociateTransitGatewayMulticastDomainMock(arg1, arg2, arg3...)
}

// DisassociateTransitGatewayPolicyTable mocked function
func (m EC2ClientMock) DisassociateTransitGatewayPolicyTable(arg1 context.Context, arg2 *ec2.DisassociateTransitGatewayPolicyTableInput, arg3 ...func(*ec2.Options)) (*ec2.DisassociateTransitGatewayPolicyTableOutput, error) {
	return m.DisassociateTransitGatewayPolicyTableMock(arg1, arg2, arg3...)
}

// DisassociateTransitGatewayRouteTable mocked function
func (m EC2ClientMock) DisassociateTransitGatewayRouteTable(arg1 context.Context, arg2 *ec2.DisassociateTransitGatewayRouteTableInput, arg3 ...func(*ec2.Options)) (*ec2.DisassociateTransitGatewayRouteTableOutput, error) {
	return m.DisassociateTransitGatewayRouteTableMock(arg1, arg2, arg3...)
}

// DisassociateTrunkInterface mocked function
func (m EC2ClientMock) DisassociateTrunkInterface(arg1 context.Context, arg2 *ec2.DisassociateTrunkInterfaceInput, arg3 ...func(*ec2.Options)) (*ec2.DisassociateTrunkInterfaceOutput, error) {
	return m.DisassociateTrunkInterfaceMock(arg1, arg2, arg3...)
}

// DisassociateVpcCidrBlock mocked function
func (m EC2ClientMock) DisassociateVpcCidrBlock(arg1 context.Context, arg2 *ec2.DisassociateVpcCidrBlockInput, arg3 ...func(*ec2.Options)) (*ec2.DisassociateVpcCidrBlockOutput, error) {
	return m.DisassociateVpcCidrBlockMock(arg1, arg2, arg3...)
}

// EnableEbsEncryptionByDefault mocked function
func (m EC2ClientMock) EnableEbsEncryptionByDefault(arg1 context.Context, arg2 *ec2.EnableEbsEncryptionByDefaultInput, arg3 ...func(*ec2.Options)) (*ec2.EnableEbsEncryptionByDefaultOutput, error) {
	return m.EnableEbsEncryptionByDefaultMock(arg1, arg2, arg3...)
}

// EnableFastLaunch mocked function
func (m EC2ClientMock) EnableFastLaunch(arg1 context.Context, arg2 *ec2.EnableFastLaunchInput, arg3 ...func(*ec2.Options)) (*ec2.EnableFastLaunchOutput, error) {
	return m.EnableFastLaunchMock(arg1, arg2, arg3...)
}

// EnableFastSnapshotRestores mocked function
func (m EC2ClientMock) EnableFastSnapshotRestores(arg1 context.Context, arg2 *ec2.EnableFastSnapshotRestoresInput, arg3 ...func(*ec2.Options)) (*ec2.EnableFastSnapshotRestoresOutput, error) {
	return m.EnableFastSnapshotRestoresMock(arg1, arg2, arg3...)
}

// EnableImageDeprecation mocked function
func (m EC2ClientMock) EnableImageDeprecation(arg1 context.Context, arg2 *ec2.EnableImageDeprecationInput, arg3 ...func(*ec2.Options)) (*ec2.EnableImageDeprecationOutput, error) {
	return m.EnableImageDeprecationMock(arg1, arg2, arg3...)
}

// EnableIpamOrganizationAdminAccount mocked function
func (m EC2ClientMock) EnableIpamOrganizationAdminAccount(arg1 context.Context, arg2 *ec2.EnableIpamOrganizationAdminAccountInput, arg3 ...func(*ec2.Options)) (*ec2.EnableIpamOrganizationAdminAccountOutput, error) {
	return m.EnableIpamOrganizationAdminAccountMock(arg1, arg2, arg3...)
}

// EnableSerialConsoleAccess mocked function
func (m EC2ClientMock) EnableSerialConsoleAccess(arg1 context.Context, arg2 *ec2.EnableSerialConsoleAccessInput, arg3 ...func(*ec2.Options)) (*ec2.EnableSerialConsoleAccessOutput, error) {
	return m.EnableSerialConsoleAccessMock(arg1, arg2, arg3...)
}

// EnableTransitGatewayRouteTablePropagation mocked function
func (m EC2ClientMock) EnableTransitGatewayRouteTablePropagation(arg1 context.Context, arg2 *ec2.EnableTransitGatewayRouteTablePropagationInput, arg3 ...func(*ec2.Options)) (*ec2.EnableTransitGatewayRouteTablePropagationOutput, error) {
	return m.EnableTransitGatewayRouteTablePropagationMock(arg1, arg2, arg3...)
}

// EnableVgwRoutePropagation mocked function
func (m EC2ClientMock) EnableVgwRoutePropagation(arg1 context.Context, arg2 *ec2.EnableVgwRoutePropagationInput, arg3 ...func(*ec2.Options)) (*ec2.EnableVgwRoutePropagationOutput, error) {
	return m.EnableVgwRoutePropagationMock(arg1, arg2, arg3...)
}

// EnableVolumeIO mocked function
func (m EC2ClientMock) EnableVolumeIO(arg1 context.Context, arg2 *ec2.EnableVolumeIOInput, arg3 ...func(*ec2.Options)) (*ec2.EnableVolumeIOOutput, error) {
	return m.EnableVolumeIOMock(arg1, arg2, arg3...)
}

// EnableVpcClassicLink mocked function
func (m EC2ClientMock) EnableVpcClassicLink(arg1 context.Context, arg2 *ec2.EnableVpcClassicLinkInput, arg3 ...func(*ec2.Options)) (*ec2.EnableVpcClassicLinkOutput, error) {
	return m.EnableVpcClassicLinkMock(arg1, arg2, arg3...)
}

// EnableVpcClassicLinkDnsSupport mocked function
func (m EC2ClientMock) EnableVpcClassicLinkDnsSupport(arg1 context.Context, arg2 *ec2.EnableVpcClassicLinkDnsSupportInput, arg3 ...func(*ec2.Options)) (*ec2.EnableVpcClassicLinkDnsSupportOutput, error) {
	return m.EnableVpcClassicLinkDnsSupportMock(arg1, arg2, arg3...)
}

// ExportClientVpnClientCertificateRevocationList mocked function
func (m EC2ClientMock) ExportClientVpnClientCertificateRevocationList(arg1 context.Context, arg2 *ec2.ExportClientVpnClientCertificateRevocationListInput, arg3 ...func(*ec2.Options)) (*ec2.ExportClientVpnClientCertificateRevocationListOutput, error) {
	return m.ExportClientVpnClientCertificateRevocationListMock(arg1, arg2, arg3...)
}

// ExportClientVpnClientConfiguration mocked function
func (m EC2ClientMock) ExportClientVpnClientConfiguration(arg1 context.Context, arg2 *ec2.ExportClientVpnClientConfigurationInput, arg3 ...func(*ec2.Options)) (*ec2.ExportClientVpnClientConfigurationOutput, error) {
	return m.ExportClientVpnClientConfigurationMock(arg1, arg2, arg3...)
}

// ExportImage mocked function
func (m EC2ClientMock) ExportImage(arg1 context.Context, arg2 *ec2.ExportImageInput, arg3 ...func(*ec2.Options)) (*ec2.ExportImageOutput, error) {
	return m.ExportImageMock(arg1, arg2, arg3...)
}

// ExportTransitGatewayRoutes mocked function
func (m EC2ClientMock) ExportTransitGatewayRoutes(arg1 context.Context, arg2 *ec2.ExportTransitGatewayRoutesInput, arg3 ...func(*ec2.Options)) (*ec2.ExportTransitGatewayRoutesOutput, error) {
	return m.ExportTransitGatewayRoutesMock(arg1, arg2, arg3...)
}

// GetAssociatedEnclaveCertificateIamRoles mocked function
func (m EC2ClientMock) GetAssociatedEnclaveCertificateIamRoles(arg1 context.Context, arg2 *ec2.GetAssociatedEnclaveCertificateIamRolesInput, arg3 ...func(*ec2.Options)) (*ec2.GetAssociatedEnclaveCertificateIamRolesOutput, error) {
	return m.GetAssociatedEnclaveCertificateIamRolesMock(arg1, arg2, arg3...)
}

// GetAssociatedIpv6PoolCidrs mocked function
func (m EC2ClientMock) GetAssociatedIpv6PoolCidrs(arg1 context.Context, arg2 *ec2.GetAssociatedIpv6PoolCidrsInput, arg3 ...func(*ec2.Options)) (*ec2.GetAssociatedIpv6PoolCidrsOutput, error) {
	return m.GetAssociatedIpv6PoolCidrsMock(arg1, arg2, arg3...)
}

// GetCapacityReservationUsage mocked function
func (m EC2ClientMock) GetCapacityReservationUsage(arg1 context.Context, arg2 *ec2.GetCapacityReservationUsageInput, arg3 ...func(*ec2.Options)) (*ec2.GetCapacityReservationUsageOutput, error) {
	return m.GetCapacityReservationUsageMock(arg1, arg2, arg3...)
}

// GetCoipPoolUsage mocked function
func (m EC2ClientMock) GetCoipPoolUsage(arg1 context.Context, arg2 *ec2.GetCoipPoolUsageInput, arg3 ...func(*ec2.Options)) (*ec2.GetCoipPoolUsageOutput, error) {
	return m.GetCoipPoolUsageMock(arg1, arg2, arg3...)
}

// GetConsoleOutput mocked function
func (m EC2ClientMock) GetConsoleOutput(arg1 context.Context, arg2 *ec2.GetConsoleOutputInput, arg3 ...func(*ec2.Options)) (*ec2.GetConsoleOutputOutput, error) {
	return m.GetConsoleOutputMock(arg1, arg2, arg3...)
}

// GetConsoleScreenshot mocked function
func (m EC2ClientMock) GetConsoleScreenshot(arg1 context.Context, arg2 *ec2.GetConsoleScreenshotInput, arg3 ...func(*ec2.Options)) (*ec2.GetConsoleScreenshotOutput, error) {
	return m.GetConsoleScreenshotMock(arg1, arg2, arg3...)
}

// GetDefaultCreditSpecification mocked function
func (m EC2ClientMock) GetDefaultCreditSpecification(arg1 context.Context, arg2 *ec2.GetDefaultCreditSpecificationInput, arg3 ...func(*ec2.Options)) (*ec2.GetDefaultCreditSpecificationOutput, error) {
	return m.GetDefaultCreditSpecificationMock(arg1, arg2, arg3...)
}

// GetEbsDefaultKmsKeyId mocked function
func (m EC2ClientMock) GetEbsDefaultKmsKeyId(arg1 context.Context, arg2 *ec2.GetEbsDefaultKmsKeyIdInput, arg3 ...func(*ec2.Options)) (*ec2.GetEbsDefaultKmsKeyIdOutput, error) {
	return m.GetEbsDefaultKmsKeyIdMock(arg1, arg2, arg3...)
}

// GetEbsEncryptionByDefault mocked function
func (m EC2ClientMock) GetEbsEncryptionByDefault(arg1 context.Context, arg2 *ec2.GetEbsEncryptionByDefaultInput, arg3 ...func(*ec2.Options)) (*ec2.GetEbsEncryptionByDefaultOutput, error) {
	return m.GetEbsEncryptionByDefaultMock(arg1, arg2, arg3...)
}

// GetFlowLogsIntegrationTemplate mocked function
func (m EC2ClientMock) GetFlowLogsIntegrationTemplate(arg1 context.Context, arg2 *ec2.GetFlowLogsIntegrationTemplateInput, arg3 ...func(*ec2.Options)) (*ec2.GetFlowLogsIntegrationTemplateOutput, error) {
	return m.GetFlowLogsIntegrationTemplateMock(arg1, arg2, arg3...)
}

// GetGroupsForCapacityReservation mocked function
func (m EC2ClientMock) GetGroupsForCapacityReservation(arg1 context.Context, arg2 *ec2.GetGroupsForCapacityReservationInput, arg3 ...func(*ec2.Options)) (*ec2.GetGroupsForCapacityReservationOutput, error) {
	return m.GetGroupsForCapacityReservationMock(arg1, arg2, arg3...)
}

// GetHostReservationPurchasePreview mocked function
func (m EC2ClientMock) GetHostReservationPurchasePreview(arg1 context.Context, arg2 *ec2.GetHostReservationPurchasePreviewInput, arg3 ...func(*ec2.Options)) (*ec2.GetHostReservationPurchasePreviewOutput, error) {
	return m.GetHostReservationPurchasePreviewMock(arg1, arg2, arg3...)
}

// GetInstanceTypesFromInstanceRequirements mocked function
func (m EC2ClientMock) GetInstanceTypesFromInstanceRequirements(arg1 context.Context, arg2 *ec2.GetInstanceTypesFromInstanceRequirementsInput, arg3 ...func(*ec2.Options)) (*ec2.GetInstanceTypesFromInstanceRequirementsOutput, error) {
	return m.GetInstanceTypesFromInstanceRequirementsMock(arg1, arg2, arg3...)
}

// GetInstanceUefiData mocked function
func (m EC2ClientMock) GetInstanceUefiData(arg1 context.Context, arg2 *ec2.GetInstanceUefiDataInput, arg3 ...func(*ec2.Options)) (*ec2.GetInstanceUefiDataOutput, error) {
	return m.GetInstanceUefiDataMock(arg1, arg2, arg3...)
}

// GetIpamAddressHistory mocked function
func (m EC2ClientMock) GetIpamAddressHistory(arg1 context.Context, arg2 *ec2.GetIpamAddressHistoryInput, arg3 ...func(*ec2.Options)) (*ec2.GetIpamAddressHistoryOutput, error) {
	return m.GetIpamAddressHistoryMock(arg1, arg2, arg3...)
}

// GetIpamPoolAllocations mocked function
func (m EC2ClientMock) GetIpamPoolAllocations(arg1 context.Context, arg2 *ec2.GetIpamPoolAllocationsInput, arg3 ...func(*ec2.Options)) (*ec2.GetIpamPoolAllocationsOutput, error) {
	return m.GetIpamPoolAllocationsMock(arg1, arg2, arg3...)
}

// GetIpamPoolCidrs mocked function
func (m EC2ClientMock) GetIpamPoolCidrs(arg1 context.Context, arg2 *ec2.GetIpamPoolCidrsInput, arg3 ...func(*ec2.Options)) (*ec2.GetIpamPoolCidrsOutput, error) {
	return m.GetIpamPoolCidrsMock(arg1, arg2, arg3...)
}

// GetIpamResourceCidrs mocked function
func (m EC2ClientMock) GetIpamResourceCidrs(arg1 context.Context, arg2 *ec2.GetIpamResourceCidrsInput, arg3 ...func(*ec2.Options)) (*ec2.GetIpamResourceCidrsOutput, error) {
	return m.GetIpamResourceCidrsMock(arg1, arg2, arg3...)
}

// GetLaunchTemplateData mocked function
func (m EC2ClientMock) GetLaunchTemplateData(arg1 context.Context, arg2 *ec2.GetLaunchTemplateDataInput, arg3 ...func(*ec2.Options)) (*ec2.GetLaunchTemplateDataOutput, error) {
	return m.GetLaunchTemplateDataMock(arg1, arg2, arg3...)
}

// GetManagedPrefixListAssociations mocked function
func (m EC2ClientMock) GetManagedPrefixListAssociations(arg1 context.Context, arg2 *ec2.GetManagedPrefixListAssociationsInput, arg3 ...func(*ec2.Options)) (*ec2.GetManagedPrefixListAssociationsOutput, error) {
	return m.GetManagedPrefixListAssociationsMock(arg1, arg2, arg3...)
}

// GetManagedPrefixListEntries mocked function
func (m EC2ClientMock) GetManagedPrefixListEntries(arg1 context.Context, arg2 *ec2.GetManagedPrefixListEntriesInput, arg3 ...func(*ec2.Options)) (*ec2.GetManagedPrefixListEntriesOutput, error) {
	return m.GetManagedPrefixListEntriesMock(arg1, arg2, arg3...)
}

// GetNetworkInsightsAccessScopeAnalysisFindings mocked function
func (m EC2ClientMock) GetNetworkInsightsAccessScopeAnalysisFindings(arg1 context.Context, arg2 *ec2.GetNetworkInsightsAccessScopeAnalysisFindingsInput, arg3 ...func(*ec2.Options)) (*ec2.GetNetworkInsightsAccessScopeAnalysisFindingsOutput, error) {
	return m.GetNetworkInsightsAccessScopeAnalysisFindingsMock(arg1, arg2, arg3...)
}

// GetNetworkInsightsAccessScopeContent mocked function
func (m EC2ClientMock) GetNetworkInsightsAccessScopeContent(arg1 context.Context, arg2 *ec2.GetNetworkInsightsAccessScopeContentInput, arg3 ...func(*ec2.Options)) (*ec2.GetNetworkInsightsAccessScopeContentOutput, error) {
	return m.GetNetworkInsightsAccessScopeContentMock(arg1, arg2, arg3...)
}

// GetPasswordData mocked function
func (m EC2ClientMock) GetPasswordData(arg1 context.Context, arg2 *ec2.GetPasswordDataInput, arg3 ...func(*ec2.Options)) (*ec2.GetPasswordDataOutput, error) {
	return m.GetPasswordDataMock(arg1, arg2, arg3...)
}

// GetReservedInstancesExchangeQuote mocked function
func (m EC2ClientMock) GetReservedInstancesExchangeQuote(arg1 context.Context, arg2 *ec2.GetReservedInstancesExchangeQuoteInput, arg3 ...func(*ec2.Options)) (*ec2.GetReservedInstancesExchangeQuoteOutput, error) {
	return m.GetReservedInstancesExchangeQuoteMock(arg1, arg2, arg3...)
}

// GetSerialConsoleAccessStatus mocked function
func (m EC2ClientMock) GetSerialConsoleAccessStatus(arg1 context.Context, arg2 *ec2.GetSerialConsoleAccessStatusInput, arg3 ...func(*ec2.Options)) (*ec2.GetSerialConsoleAccessStatusOutput, error) {
	return m.GetSerialConsoleAccessStatusMock(arg1, arg2, arg3...)
}

// GetSpotPlacementScores mocked function
func (m EC2ClientMock) GetSpotPlacementScores(arg1 context.Context, arg2 *ec2.GetSpotPlacementScoresInput, arg3 ...func(*ec2.Options)) (*ec2.GetSpotPlacementScoresOutput, error) {
	return m.GetSpotPlacementScoresMock(arg1, arg2, arg3...)
}

// GetSubnetCidrReservations mocked function
func (m EC2ClientMock) GetSubnetCidrReservations(arg1 context.Context, arg2 *ec2.GetSubnetCidrReservationsInput, arg3 ...func(*ec2.Options)) (*ec2.GetSubnetCidrReservationsOutput, error) {
	return m.GetSubnetCidrReservationsMock(arg1, arg2, arg3...)
}

// GetTransitGatewayAttachmentPropagations mocked function
func (m EC2ClientMock) GetTransitGatewayAttachmentPropagations(arg1 context.Context, arg2 *ec2.GetTransitGatewayAttachmentPropagationsInput, arg3 ...func(*ec2.Options)) (*ec2.GetTransitGatewayAttachmentPropagationsOutput, error) {
	return m.GetTransitGatewayAttachmentPropagationsMock(arg1, arg2, arg3...)
}

// GetTransitGatewayMulticastDomainAssociations mocked function
func (m EC2ClientMock) GetTransitGatewayMulticastDomainAssociations(arg1 context.Context, arg2 *ec2.GetTransitGatewayMulticastDomainAssociationsInput, arg3 ...func(*ec2.Options)) (*ec2.GetTransitGatewayMulticastDomainAssociationsOutput, error) {
	return m.GetTransitGatewayMulticastDomainAssociationsMock(arg1, arg2, arg3...)
}

// GetTransitGatewayPolicyTableAssociations mocked function
func (m EC2ClientMock) GetTransitGatewayPolicyTableAssociations(arg1 context.Context, arg2 *ec2.GetTransitGatewayPolicyTableAssociationsInput, arg3 ...func(*ec2.Options)) (*ec2.GetTransitGatewayPolicyTableAssociationsOutput, error) {
	return m.GetTransitGatewayPolicyTableAssociationsMock(arg1, arg2, arg3...)
}

// GetTransitGatewayPolicyTableEntries mocked function
func (m EC2ClientMock) GetTransitGatewayPolicyTableEntries(arg1 context.Context, arg2 *ec2.GetTransitGatewayPolicyTableEntriesInput, arg3 ...func(*ec2.Options)) (*ec2.GetTransitGatewayPolicyTableEntriesOutput, error) {
	return m.GetTransitGatewayPolicyTableEntriesMock(arg1, arg2, arg3...)
}

// GetTransitGatewayPrefixListReferences mocked function
func (m EC2ClientMock) GetTransitGatewayPrefixListReferences(arg1 context.Context, arg2 *ec2.GetTransitGatewayPrefixListReferencesInput, arg3 ...func(*ec2.Options)) (*ec2.GetTransitGatewayPrefixListReferencesOutput, error) {
	return m.GetTransitGatewayPrefixListReferencesMock(arg1, arg2, arg3...)
}

// GetTransitGatewayRouteTableAssociations mocked function
func (m EC2ClientMock) GetTransitGatewayRouteTableAssociations(arg1 context.Context, arg2 *ec2.GetTransitGatewayRouteTableAssociationsInput, arg3 ...func(*ec2.Options)) (*ec2.GetTransitGatewayRouteTableAssociationsOutput, error) {
	return m.GetTransitGatewayRouteTableAssociationsMock(arg1, arg2, arg3...)
}

// GetTransitGatewayRouteTablePropagations mocked function
func (m EC2ClientMock) GetTransitGatewayRouteTablePropagations(arg1 context.Context, arg2 *ec2.GetTransitGatewayRouteTablePropagationsInput, arg3 ...func(*ec2.Options)) (*ec2.GetTransitGatewayRouteTablePropagationsOutput, error) {
	return m.GetTransitGatewayRouteTablePropagationsMock(arg1, arg2, arg3...)
}

// GetVpnConnectionDeviceSampleConfiguration mocked function
func (m EC2ClientMock) GetVpnConnectionDeviceSampleConfiguration(arg1 context.Context, arg2 *ec2.GetVpnConnectionDeviceSampleConfigurationInput, arg3 ...func(*ec2.Options)) (*ec2.GetVpnConnectionDeviceSampleConfigurationOutput, error) {
	return m.GetVpnConnectionDeviceSampleConfigurationMock(arg1, arg2, arg3...)
}

// GetVpnConnectionDeviceTypes mocked function
func (m EC2ClientMock) GetVpnConnectionDeviceTypes(arg1 context.Context, arg2 *ec2.GetVpnConnectionDeviceTypesInput, arg3 ...func(*ec2.Options)) (*ec2.GetVpnConnectionDeviceTypesOutput, error) {
	return m.GetVpnConnectionDeviceTypesMock(arg1, arg2, arg3...)
}

// ImportClientVpnClientCertificateRevocationList mocked function
func (m EC2ClientMock) ImportClientVpnClientCertificateRevocationList(arg1 context.Context, arg2 *ec2.ImportClientVpnClientCertificateRevocationListInput, arg3 ...func(*ec2.Options)) (*ec2.ImportClientVpnClientCertificateRevocationListOutput, error) {
	return m.ImportClientVpnClientCertificateRevocationListMock(arg1, arg2, arg3...)
}

// ImportImage mocked function
func (m EC2ClientMock) ImportImage(arg1 context.Context, arg2 *ec2.ImportImageInput, arg3 ...func(*ec2.Options)) (*ec2.ImportImageOutput, error) {
	return m.ImportImageMock(arg1, arg2, arg3...)
}

// ImportInstance mocked function
func (m EC2ClientMock) ImportInstance(arg1 context.Context, arg2 *ec2.ImportInstanceInput, arg3 ...func(*ec2.Options)) (*ec2.ImportInstanceOutput, error) {
	return m.ImportInstanceMock(arg1, arg2, arg3...)
}

// ImportKeyPair mocked function
func (m EC2ClientMock) ImportKeyPair(arg1 context.Context, arg2 *ec2.ImportKeyPairInput, arg3 ...func(*ec2.Options)) (*ec2.ImportKeyPairOutput, error) {
	return m.ImportKeyPairMock(arg1, arg2, arg3...)
}

// ImportSnapshot mocked function
func (m EC2ClientMock) ImportSnapshot(arg1 context.Context, arg2 *ec2.ImportSnapshotInput, arg3 ...func(*ec2.Options)) (*ec2.ImportSnapshotOutput, error) {
	return m.ImportSnapshotMock(arg1, arg2, arg3...)
}

// ImportVolume mocked function
func (m EC2ClientMock) ImportVolume(arg1 context.Context, arg2 *ec2.ImportVolumeInput, arg3 ...func(*ec2.Options)) (*ec2.ImportVolumeOutput, error) {
	return m.ImportVolumeMock(arg1, arg2, arg3...)
}

// ListImagesInRecycleBin mocked function
func (m EC2ClientMock) ListImagesInRecycleBin(arg1 context.Context, arg2 *ec2.ListImagesInRecycleBinInput, arg3 ...func(*ec2.Options)) (*ec2.ListImagesInRecycleBinOutput, error) {
	return m.ListImagesInRecycleBinMock(arg1, arg2, arg3...)
}

// ListSnapshotsInRecycleBin mocked function
func (m EC2ClientMock) ListSnapshotsInRecycleBin(arg1 context.Context, arg2 *ec2.ListSnapshotsInRecycleBinInput, arg3 ...func(*ec2.Options)) (*ec2.ListSnapshotsInRecycleBinOutput, error) {
	return m.ListSnapshotsInRecycleBinMock(arg1, arg2, arg3...)
}

// ModifyAddressAttribute mocked function
func (m EC2ClientMock) ModifyAddressAttribute(arg1 context.Context, arg2 *ec2.ModifyAddressAttributeInput, arg3 ...func(*ec2.Options)) (*ec2.ModifyAddressAttributeOutput, error) {
	return m.ModifyAddressAttributeMock(arg1, arg2, arg3...)
}

// ModifyAvailabilityZoneGroup mocked function
func (m EC2ClientMock) ModifyAvailabilityZoneGroup(arg1 context.Context, arg2 *ec2.ModifyAvailabilityZoneGroupInput, arg3 ...func(*ec2.Options)) (*ec2.ModifyAvailabilityZoneGroupOutput, error) {
	return m.ModifyAvailabilityZoneGroupMock(arg1, arg2, arg3...)
}

// ModifyCapacityReservation mocked function
func (m EC2ClientMock) ModifyCapacityReservation(arg1 context.Context, arg2 *ec2.ModifyCapacityReservationInput, arg3 ...func(*ec2.Options)) (*ec2.ModifyCapacityReservationOutput, error) {
	return m.ModifyCapacityReservationMock(arg1, arg2, arg3...)
}

// ModifyCapacityReservationFleet mocked function
func (m EC2ClientMock) ModifyCapacityReservationFleet(arg1 context.Context, arg2 *ec2.ModifyCapacityReservationFleetInput, arg3 ...func(*ec2.Options)) (*ec2.ModifyCapacityReservationFleetOutput, error) {
	return m.ModifyCapacityReservationFleetMock(arg1, arg2, arg3...)
}

// ModifyClientVpnEndpoint mocked function
func (m EC2ClientMock) ModifyClientVpnEndpoint(arg1 context.Context, arg2 *ec2.ModifyClientVpnEndpointInput, arg3 ...func(*ec2.Options)) (*ec2.ModifyClientVpnEndpointOutput, error) {
	return m.ModifyClientVpnEndpointMock(arg1, arg2, arg3...)
}

// ModifyDefaultCreditSpecification mocked function
func (m EC2ClientMock) ModifyDefaultCreditSpecification(arg1 context.Context, arg2 *ec2.ModifyDefaultCreditSpecificationInput, arg3 ...func(*ec2.Options)) (*ec2.ModifyDefaultCreditSpecificationOutput, error) {
	return m.ModifyDefaultCreditSpecificationMock(arg1, arg2, arg3...)
}

// ModifyEbsDefaultKmsKeyId mocked function
func (m EC2ClientMock) ModifyEbsDefaultKmsKeyId(arg1 context.Context, arg2 *ec2.ModifyEbsDefaultKmsKeyIdInput, arg3 ...func(*ec2.Options)) (*ec2.ModifyEbsDefaultKmsKeyIdOutput, error) {
	return m.ModifyEbsDefaultKmsKeyIdMock(arg1, arg2, arg3...)
}

// ModifyFleet mocked function
func (m EC2ClientMock) ModifyFleet(arg1 context.Context, arg2 *ec2.ModifyFleetInput, arg3 ...func(*ec2.Options)) (*ec2.ModifyFleetOutput, error) {
	return m.ModifyFleetMock(arg1, arg2, arg3...)
}

// ModifyFpgaImageAttribute mocked function
func (m EC2ClientMock) ModifyFpgaImageAttribute(arg1 context.Context, arg2 *ec2.ModifyFpgaImageAttributeInput, arg3 ...func(*ec2.Options)) (*ec2.ModifyFpgaImageAttributeOutput, error) {
	return m.ModifyFpgaImageAttributeMock(arg1, arg2, arg3...)
}

// ModifyHosts mocked function
func (m EC2ClientMock) ModifyHosts(arg1 context.Context, arg2 *ec2.ModifyHostsInput, arg3 ...func(*ec2.Options)) (*ec2.ModifyHostsOutput, error) {
	return m.ModifyHostsMock(arg1, arg2, arg3...)
}

// ModifyIdFormat mocked function
func (m EC2ClientMock) ModifyIdFormat(arg1 context.Context, arg2 *ec2.ModifyIdFormatInput, arg3 ...func(*ec2.Options)) (*ec2.ModifyIdFormatOutput, error) {
	return m.ModifyIdFormatMock(arg1, arg2, arg3...)
}

// ModifyIdentityIdFormat mocked function
func (m EC2ClientMock) ModifyIdentityIdFormat(arg1 context.Context, arg2 *ec2.ModifyIdentityIdFormatInput, arg3 ...func(*ec2.Options)) (*ec2.ModifyIdentityIdFormatOutput, error) {
	return m.ModifyIdentityIdFormatMock(arg1, arg2, arg3...)
}

// ModifyImageAttribute mocked function
func (m EC2ClientMock) ModifyImageAttribute(arg1 context.Context, arg2 *ec2.ModifyImageAttributeInput, arg3 ...func(*ec2.Options)) (*ec2.ModifyImageAttributeOutput, error) {
	return m.ModifyImageAttributeMock(arg1, arg2, arg3...)
}

// ModifyInstanceAttribute mocked function
func (m EC2ClientMock) ModifyInstanceAttribute(arg1 context.Context, arg2 *ec2.ModifyInstanceAttributeInput, arg3 ...func(*ec2.Options)) (*ec2.ModifyInstanceAttributeOutput, error) {
	return m.ModifyInstanceAttributeMock(arg1, arg2, arg3...)
}

// ModifyInstanceCapacityReservationAttributes mocked function
func (m EC2ClientMock) ModifyInstanceCapacityReservationAttributes(arg1 context.Context, arg2 *ec2.ModifyInstanceCapacityReservationAttributesInput, arg3 ...func(*ec2.Options)) (*ec2.ModifyInstanceCapacityReservationAttributesOutput, error) {
	return m.ModifyInstanceCapacityReservationAttributesMock(arg1, arg2, arg3...)
}

// ModifyInstanceCreditSpecification mocked function
func (m EC2ClientMock) ModifyInstanceCreditSpecification(arg1 context.Context, arg2 *ec2.ModifyInstanceCreditSpecificationInput, arg3 ...func(*ec2.Options)) (*ec2.ModifyInstanceCreditSpecificationOutput, error) {
	return m.ModifyInstanceCreditSpecificationMock(arg1, arg2, arg3...)
}

// ModifyInstanceEventStartTime mocked function
func (m EC2ClientMock) ModifyInstanceEventStartTime(arg1 context.Context, arg2 *ec2.ModifyInstanceEventStartTimeInput, arg3 ...func(*ec2.Options)) (*ec2.ModifyInstanceEventStartTimeOutput, error) {
	return m.ModifyInstanceEventStartTimeMock(arg1, arg2, arg3...)
}

// ModifyInstanceEventWindow mocked function
func (m EC2ClientMock) ModifyInstanceEventWindow(arg1 context.Context, arg2 *ec2.ModifyInstanceEventWindowInput, arg3 ...func(*ec2.Options)) (*ec2.ModifyInstanceEventWindowOutput, error) {
	return m.ModifyInstanceEventWindowMock(arg1, arg2, arg3...)
}

// ModifyInstanceMaintenanceOptions mocked function
func (m EC2ClientMock) ModifyInstanceMaintenanceOptions(arg1 context.Context, arg2 *ec2.ModifyInstanceMaintenanceOptionsInput, arg3 ...func(*ec2.Options)) (*ec2.ModifyInstanceMaintenanceOptionsOutput, error) {
	return m.ModifyInstanceMaintenanceOptionsMock(arg1, arg2, arg3...)
}

// ModifyInstanceMetadataOptions mocked function
func (m EC2ClientMock) ModifyInstanceMetadataOptions(arg1 context.Context, arg2 *ec2.ModifyInstanceMetadataOptionsInput, arg3 ...func(*ec2.Options)) (*ec2.ModifyInstanceMetadataOptionsOutput, error) {
	return m.ModifyInstanceMetadataOptionsMock(arg1, arg2, arg3...)
}

// ModifyInstancePlacement mocked function
func (m EC2ClientMock) ModifyInstancePlacement(arg1 context.Context, arg2 *ec2.ModifyInstancePlacementInput, arg3 ...func(*ec2.Options)) (*ec2.ModifyInstancePlacementOutput, error) {
	return m.ModifyInstancePlacementMock(arg1, arg2, arg3...)
}

// ModifyIpam mocked function
func (m EC2ClientMock) ModifyIpam(arg1 context.Context, arg2 *ec2.ModifyIpamInput, arg3 ...func(*ec2.Options)) (*ec2.ModifyIpamOutput, error) {
	return m.ModifyIpamMock(arg1, arg2, arg3...)
}

// ModifyIpamPool mocked function
func (m EC2ClientMock) ModifyIpamPool(arg1 context.Context, arg2 *ec2.ModifyIpamPoolInput, arg3 ...func(*ec2.Options)) (*ec2.ModifyIpamPoolOutput, error) {
	return m.ModifyIpamPoolMock(arg1, arg2, arg3...)
}

// ModifyIpamResourceCidr mocked function
func (m EC2ClientMock) ModifyIpamResourceCidr(arg1 context.Context, arg2 *ec2.ModifyIpamResourceCidrInput, arg3 ...func(*ec2.Options)) (*ec2.ModifyIpamResourceCidrOutput, error) {
	return m.ModifyIpamResourceCidrMock(arg1, arg2, arg3...)
}

// ModifyIpamScope mocked function
func (m EC2ClientMock) ModifyIpamScope(arg1 context.Context, arg2 *ec2.ModifyIpamScopeInput, arg3 ...func(*ec2.Options)) (*ec2.ModifyIpamScopeOutput, error) {
	return m.ModifyIpamScopeMock(arg1, arg2, arg3...)
}

// ModifyLaunchTemplate mocked function
func (m EC2ClientMock) ModifyLaunchTemplate(arg1 context.Context, arg2 *ec2.ModifyLaunchTemplateInput, arg3 ...func(*ec2.Options)) (*ec2.ModifyLaunchTemplateOutput, error) {
	return m.ModifyLaunchTemplateMock(arg1, arg2, arg3...)
}

// ModifyManagedPrefixList mocked function
func (m EC2ClientMock) ModifyManagedPrefixList(arg1 context.Context, arg2 *ec2.ModifyManagedPrefixListInput, arg3 ...func(*ec2.Options)) (*ec2.ModifyManagedPrefixListOutput, error) {
	return m.ModifyManagedPrefixListMock(arg1, arg2, arg3...)
}

// ModifyNetworkInterfaceAttribute mocked function
func (m EC2ClientMock) ModifyNetworkInterfaceAttribute(arg1 context.Context, arg2 *ec2.ModifyNetworkInterfaceAttributeInput, arg3 ...func(*ec2.Options)) (*ec2.ModifyNetworkInterfaceAttributeOutput, error) {
	return m.ModifyNetworkInterfaceAttributeMock(arg1, arg2, arg3...)
}

// ModifyPrivateDnsNameOptions mocked function
func (m EC2ClientMock) ModifyPrivateDnsNameOptions(arg1 context.Context, arg2 *ec2.ModifyPrivateDnsNameOptionsInput, arg3 ...func(*ec2.Options)) (*ec2.ModifyPrivateDnsNameOptionsOutput, error) {
	return m.ModifyPrivateDnsNameOptionsMock(arg1, arg2, arg3...)
}

// ModifyReservedInstances mocked function
func (m EC2ClientMock) ModifyReservedInstances(arg1 context.Context, arg2 *ec2.ModifyReservedInstancesInput, arg3 ...func(*ec2.Options)) (*ec2.ModifyReservedInstancesOutput, error) {
	return m.ModifyReservedInstancesMock(arg1, arg2, arg3...)
}

// ModifySecurityGroupRules mocked function
func (m EC2ClientMock) ModifySecurityGroupRules(arg1 context.Context, arg2 *ec2.ModifySecurityGroupRulesInput, arg3 ...func(*ec2.Options)) (*ec2.ModifySecurityGroupRulesOutput, error) {
	return m.ModifySecurityGroupRulesMock(arg1, arg2, arg3...)
}

// ModifySnapshotAttribute mocked function
func (m EC2ClientMock) ModifySnapshotAttribute(arg1 context.Context, arg2 *ec2.ModifySnapshotAttributeInput, arg3 ...func(*ec2.Options)) (*ec2.ModifySnapshotAttributeOutput, error) {
	return m.ModifySnapshotAttributeMock(arg1, arg2, arg3...)
}

// ModifySnapshotTier mocked function
func (m EC2ClientMock) ModifySnapshotTier(arg1 context.Context, arg2 *ec2.ModifySnapshotTierInput, arg3 ...func(*ec2.Options)) (*ec2.ModifySnapshotTierOutput, error) {
	return m.ModifySnapshotTierMock(arg1, arg2, arg3...)
}

// ModifySpotFleetRequest mocked function
func (m EC2ClientMock) ModifySpotFleetRequest(arg1 context.Context, arg2 *ec2.ModifySpotFleetRequestInput, arg3 ...func(*ec2.Options)) (*ec2.ModifySpotFleetRequestOutput, error) {
	return m.ModifySpotFleetRequestMock(arg1, arg2, arg3...)
}

// ModifySubnetAttribute mocked function
func (m EC2ClientMock) ModifySubnetAttribute(arg1 context.Context, arg2 *ec2.ModifySubnetAttributeInput, arg3 ...func(*ec2.Options)) (*ec2.ModifySubnetAttributeOutput, error) {
	return m.ModifySubnetAttributeMock(arg1, arg2, arg3...)
}

// ModifyTrafficMirrorFilterNetworkServices mocked function
func (m EC2ClientMock) ModifyTrafficMirrorFilterNetworkServices(arg1 context.Context, arg2 *ec2.ModifyTrafficMirrorFilterNetworkServicesInput, arg3 ...func(*ec2.Options)) (*ec2.ModifyTrafficMirrorFilterNetworkServicesOutput, error) {
	return m.ModifyTrafficMirrorFilterNetworkServicesMock(arg1, arg2, arg3...)
}

// ModifyTrafficMirrorFilterRule mocked function
func (m EC2ClientMock) ModifyTrafficMirrorFilterRule(arg1 context.Context, arg2 *ec2.ModifyTrafficMirrorFilterRuleInput, arg3 ...func(*ec2.Options)) (*ec2.ModifyTrafficMirrorFilterRuleOutput, error) {
	return m.ModifyTrafficMirrorFilterRuleMock(arg1, arg2, arg3...)
}

// ModifyTrafficMirrorSession mocked function
func (m EC2ClientMock) ModifyTrafficMirrorSession(arg1 context.Context, arg2 *ec2.ModifyTrafficMirrorSessionInput, arg3 ...func(*ec2.Options)) (*ec2.ModifyTrafficMirrorSessionOutput, error) {
	return m.ModifyTrafficMirrorSessionMock(arg1, arg2, arg3...)
}

// ModifyTransitGateway mocked function
func (m EC2ClientMock) ModifyTransitGateway(arg1 context.Context, arg2 *ec2.ModifyTransitGatewayInput, arg3 ...func(*ec2.Options)) (*ec2.ModifyTransitGatewayOutput, error) {
	return m.ModifyTransitGatewayMock(arg1, arg2, arg3...)
}

// ModifyTransitGatewayPrefixListReference mocked function
func (m EC2ClientMock) ModifyTransitGatewayPrefixListReference(arg1 context.Context, arg2 *ec2.ModifyTransitGatewayPrefixListReferenceInput, arg3 ...func(*ec2.Options)) (*ec2.ModifyTransitGatewayPrefixListReferenceOutput, error) {
	return m.ModifyTransitGatewayPrefixListReferenceMock(arg1, arg2, arg3...)
}

// ModifyTransitGatewayVpcAttachment mocked function
func (m EC2ClientMock) ModifyTransitGatewayVpcAttachment(arg1 context.Context, arg2 *ec2.ModifyTransitGatewayVpcAttachmentInput, arg3 ...func(*ec2.Options)) (*ec2.ModifyTransitGatewayVpcAttachmentOutput, error) {
	return m.ModifyTransitGatewayVpcAttachmentMock(arg1, arg2, arg3...)
}

// ModifyVolume mocked function
func (m EC2ClientMock) ModifyVolume(arg1 context.Context, arg2 *ec2.ModifyVolumeInput, arg3 ...func(*ec2.Options)) (*ec2.ModifyVolumeOutput, error) {
	return m.ModifyVolumeMock(arg1, arg2, arg3...)
}

// ModifyVolumeAttribute mocked function
func (m EC2ClientMock) ModifyVolumeAttribute(arg1 context.Context, arg2 *ec2.ModifyVolumeAttributeInput, arg3 ...func(*ec2.Options)) (*ec2.ModifyVolumeAttributeOutput, error) {
	return m.ModifyVolumeAttributeMock(arg1, arg2, arg3...)
}

// ModifyVpcAttribute mocked function
func (m EC2ClientMock) ModifyVpcAttribute(arg1 context.Context, arg2 *ec2.ModifyVpcAttributeInput, arg3 ...func(*ec2.Options)) (*ec2.ModifyVpcAttributeOutput, error) {
	return m.ModifyVpcAttributeMock(arg1, arg2, arg3...)
}

// ModifyVpcEndpoint mocked function
func (m EC2ClientMock) ModifyVpcEndpoint(arg1 context.Context, arg2 *ec2.ModifyVpcEndpointInput, arg3 ...func(*ec2.Options)) (*ec2.ModifyVpcEndpointOutput, error) {
	return m.ModifyVpcEndpointMock(arg1, arg2, arg3...)
}

// ModifyVpcEndpointConnectionNotification mocked function
func (m EC2ClientMock) ModifyVpcEndpointConnectionNotification(arg1 context.Context, arg2 *ec2.ModifyVpcEndpointConnectionNotificationInput, arg3 ...func(*ec2.Options)) (*ec2.ModifyVpcEndpointConnectionNotificationOutput, error) {
	return m.ModifyVpcEndpointConnectionNotificationMock(arg1, arg2, arg3...)
}

// ModifyVpcEndpointServiceConfiguration mocked function
func (m EC2ClientMock) ModifyVpcEndpointServiceConfiguration(arg1 context.Context, arg2 *ec2.ModifyVpcEndpointServiceConfigurationInput, arg3 ...func(*ec2.Options)) (*ec2.ModifyVpcEndpointServiceConfigurationOutput, error) {
	return m.ModifyVpcEndpointServiceConfigurationMock(arg1, arg2, arg3...)
}

// ModifyVpcEndpointServicePayerResponsibility mocked function
func (m EC2ClientMock) ModifyVpcEndpointServicePayerResponsibility(arg1 context.Context, arg2 *ec2.ModifyVpcEndpointServicePayerResponsibilityInput, arg3 ...func(*ec2.Options)) (*ec2.ModifyVpcEndpointServicePayerResponsibilityOutput, error) {
	return m.ModifyVpcEndpointServicePayerResponsibilityMock(arg1, arg2, arg3...)
}

// ModifyVpcEndpointServicePermissions mocked function
func (m EC2ClientMock) ModifyVpcEndpointServicePermissions(arg1 context.Context, arg2 *ec2.ModifyVpcEndpointServicePermissionsInput, arg3 ...func(*ec2.Options)) (*ec2.ModifyVpcEndpointServicePermissionsOutput, error) {
	return m.ModifyVpcEndpointServicePermissionsMock(arg1, arg2, arg3...)
}

// ModifyVpcPeeringConnectionOptions mocked function
func (m EC2ClientMock) ModifyVpcPeeringConnectionOptions(arg1 context.Context, arg2 *ec2.ModifyVpcPeeringConnectionOptionsInput, arg3 ...func(*ec2.Options)) (*ec2.ModifyVpcPeeringConnectionOptionsOutput, error) {
	return m.ModifyVpcPeeringConnectionOptionsMock(arg1, arg2, arg3...)
}

// ModifyVpcTenancy mocked function
func (m EC2ClientMock) ModifyVpcTenancy(arg1 context.Context, arg2 *ec2.ModifyVpcTenancyInput, arg3 ...func(*ec2.Options)) (*ec2.ModifyVpcTenancyOutput, error) {
	return m.ModifyVpcTenancyMock(arg1, arg2, arg3...)
}

// ModifyVpnConnection mocked function
func (m EC2ClientMock) ModifyVpnConnection(arg1 context.Context, arg2 *ec2.ModifyVpnConnectionInput, arg3 ...func(*ec2.Options)) (*ec2.ModifyVpnConnectionOutput, error) {
	return m.ModifyVpnConnectionMock(arg1, arg2, arg3...)
}

// ModifyVpnConnectionOptions mocked function
func (m EC2ClientMock) ModifyVpnConnectionOptions(arg1 context.Context, arg2 *ec2.ModifyVpnConnectionOptionsInput, arg3 ...func(*ec2.Options)) (*ec2.ModifyVpnConnectionOptionsOutput, error) {
	return m.ModifyVpnConnectionOptionsMock(arg1, arg2, arg3...)
}

// ModifyVpnTunnelCertificate mocked function
func (m EC2ClientMock) ModifyVpnTunnelCertificate(arg1 context.Context, arg2 *ec2.ModifyVpnTunnelCertificateInput, arg3 ...func(*ec2.Options)) (*ec2.ModifyVpnTunnelCertificateOutput, error) {
	return m.ModifyVpnTunnelCertificateMock(arg1, arg2, arg3...)
}

// ModifyVpnTunnelOptions mocked function
func (m EC2ClientMock) ModifyVpnTunnelOptions(arg1 context.Context, arg2 *ec2.ModifyVpnTunnelOptionsInput, arg3 ...func(*ec2.Options)) (*ec2.ModifyVpnTunnelOptionsOutput, error) {
	return m.ModifyVpnTunnelOptionsMock(arg1, arg2, arg3...)
}

// MonitorInstances mocked function
func (m EC2ClientMock) MonitorInstances(arg1 context.Context, arg2 *ec2.MonitorInstancesInput, arg3 ...func(*ec2.Options)) (*ec2.MonitorInstancesOutput, error) {
	return m.MonitorInstancesMock(arg1, arg2, arg3...)
}

// MoveAddressToVpc mocked function
func (m EC2ClientMock) MoveAddressToVpc(arg1 context.Context, arg2 *ec2.MoveAddressToVpcInput, arg3 ...func(*ec2.Options)) (*ec2.MoveAddressToVpcOutput, error) {
	return m.MoveAddressToVpcMock(arg1, arg2, arg3...)
}

// MoveByoipCidrToIpam mocked function
func (m EC2ClientMock) MoveByoipCidrToIpam(arg1 context.Context, arg2 *ec2.MoveByoipCidrToIpamInput, arg3 ...func(*ec2.Options)) (*ec2.MoveByoipCidrToIpamOutput, error) {
	return m.MoveByoipCidrToIpamMock(arg1, arg2, arg3...)
}

// ProvisionByoipCidr mocked function
func (m EC2ClientMock) ProvisionByoipCidr(arg1 context.Context, arg2 *ec2.ProvisionByoipCidrInput, arg3 ...func(*ec2.Options)) (*ec2.ProvisionByoipCidrOutput, error) {
	return m.ProvisionByoipCidrMock(arg1, arg2, arg3...)
}

// ProvisionIpamPoolCidr mocked function
func (m EC2ClientMock) ProvisionIpamPoolCidr(arg1 context.Context, arg2 *ec2.ProvisionIpamPoolCidrInput, arg3 ...func(*ec2.Options)) (*ec2.ProvisionIpamPoolCidrOutput, error) {
	return m.ProvisionIpamPoolCidrMock(arg1, arg2, arg3...)
}

// ProvisionPublicIpv4PoolCidr mocked function
func (m EC2ClientMock) ProvisionPublicIpv4PoolCidr(arg1 context.Context, arg2 *ec2.ProvisionPublicIpv4PoolCidrInput, arg3 ...func(*ec2.Options)) (*ec2.ProvisionPublicIpv4PoolCidrOutput, error) {
	return m.ProvisionPublicIpv4PoolCidrMock(arg1, arg2, arg3...)
}

// PurchaseHostReservation mocked function
func (m EC2ClientMock) PurchaseHostReservation(arg1 context.Context, arg2 *ec2.PurchaseHostReservationInput, arg3 ...func(*ec2.Options)) (*ec2.PurchaseHostReservationOutput, error) {
	return m.PurchaseHostReservationMock(arg1, arg2, arg3...)
}

// PurchaseReservedInstancesOffering mocked function
func (m EC2ClientMock) PurchaseReservedInstancesOffering(arg1 context.Context, arg2 *ec2.PurchaseReservedInstancesOfferingInput, arg3 ...func(*ec2.Options)) (*ec2.PurchaseReservedInstancesOfferingOutput, error) {
	return m.PurchaseReservedInstancesOfferingMock(arg1, arg2, arg3...)
}

// PurchaseScheduledInstances mocked function
func (m EC2ClientMock) PurchaseScheduledInstances(arg1 context.Context, arg2 *ec2.PurchaseScheduledInstancesInput, arg3 ...func(*ec2.Options)) (*ec2.PurchaseScheduledInstancesOutput, error) {
	return m.PurchaseScheduledInstancesMock(arg1, arg2, arg3...)
}

// RebootInstances mocked function
func (m EC2ClientMock) RebootInstances(arg1 context.Context, arg2 *ec2.RebootInstancesInput, arg3 ...func(*ec2.Options)) (*ec2.RebootInstancesOutput, error) {
	return m.RebootInstancesMock(arg1, arg2, arg3...)
}

// RegisterImage mocked function
func (m EC2ClientMock) RegisterImage(arg1 context.Context, arg2 *ec2.RegisterImageInput, arg3 ...func(*ec2.Options)) (*ec2.RegisterImageOutput, error) {
	return m.RegisterImageMock(arg1, arg2, arg3...)
}

// RegisterInstanceEventNotificationAttributes mocked function
func (m EC2ClientMock) RegisterInstanceEventNotificationAttributes(arg1 context.Context, arg2 *ec2.RegisterInstanceEventNotificationAttributesInput, arg3 ...func(*ec2.Options)) (*ec2.RegisterInstanceEventNotificationAttributesOutput, error) {
	return m.RegisterInstanceEventNotificationAttributesMock(arg1, arg2, arg3...)
}

// RegisterTransitGatewayMulticastGroupMembers mocked function
func (m EC2ClientMock) RegisterTransitGatewayMulticastGroupMembers(arg1 context.Context, arg2 *ec2.RegisterTransitGatewayMulticastGroupMembersInput, arg3 ...func(*ec2.Options)) (*ec2.RegisterTransitGatewayMulticastGroupMembersOutput, error) {
	return m.RegisterTransitGatewayMulticastGroupMembersMock(arg1, arg2, arg3...)
}

// RegisterTransitGatewayMulticastGroupSources mocked function
func (m EC2ClientMock) RegisterTransitGatewayMulticastGroupSources(arg1 context.Context, arg2 *ec2.RegisterTransitGatewayMulticastGroupSourcesInput, arg3 ...func(*ec2.Options)) (*ec2.RegisterTransitGatewayMulticastGroupSourcesOutput, error) {
	return m.RegisterTransitGatewayMulticastGroupSourcesMock(arg1, arg2, arg3...)
}

// RejectTransitGatewayMulticastDomainAssociations mocked function
func (m EC2ClientMock) RejectTransitGatewayMulticastDomainAssociations(arg1 context.Context, arg2 *ec2.RejectTransitGatewayMulticastDomainAssociationsInput, arg3 ...func(*ec2.Options)) (*ec2.RejectTransitGatewayMulticastDomainAssociationsOutput, error) {
	return m.RejectTransitGatewayMulticastDomainAssociationsMock(arg1, arg2, arg3...)
}

// RejectTransitGatewayPeeringAttachment mocked function
func (m EC2ClientMock) RejectTransitGatewayPeeringAttachment(arg1 context.Context, arg2 *ec2.RejectTransitGatewayPeeringAttachmentInput, arg3 ...func(*ec2.Options)) (*ec2.RejectTransitGatewayPeeringAttachmentOutput, error) {
	return m.RejectTransitGatewayPeeringAttachmentMock(arg1, arg2, arg3...)
}

// RejectTransitGatewayVpcAttachment mocked function
func (m EC2ClientMock) RejectTransitGatewayVpcAttachment(arg1 context.Context, arg2 *ec2.RejectTransitGatewayVpcAttachmentInput, arg3 ...func(*ec2.Options)) (*ec2.RejectTransitGatewayVpcAttachmentOutput, error) {
	return m.RejectTransitGatewayVpcAttachmentMock(arg1, arg2, arg3...)
}

// RejectVpcEndpointConnections mocked function
func (m EC2ClientMock) RejectVpcEndpointConnections(arg1 context.Context, arg2 *ec2.RejectVpcEndpointConnectionsInput, arg3 ...func(*ec2.Options)) (*ec2.RejectVpcEndpointConnectionsOutput, error) {
	return m.RejectVpcEndpointConnectionsMock(arg1, arg2, arg3...)
}

// RejectVpcPeeringConnection mocked function
func (m EC2ClientMock) RejectVpcPeeringConnection(arg1 context.Context, arg2 *ec2.RejectVpcPeeringConnectionInput, arg3 ...func(*ec2.Options)) (*ec2.RejectVpcPeeringConnectionOutput, error) {
	return m.RejectVpcPeeringConnectionMock(arg1, arg2, arg3...)
}

// ReleaseAddress mocked function
func (m EC2ClientMock) ReleaseAddress(arg1 context.Context, arg2 *ec2.ReleaseAddressInput, arg3 ...func(*ec2.Options)) (*ec2.ReleaseAddressOutput, error) {
	return m.ReleaseAddressMock(arg1, arg2, arg3...)
}

// ReleaseHosts mocked function
func (m EC2ClientMock) ReleaseHosts(arg1 context.Context, arg2 *ec2.ReleaseHostsInput, arg3 ...func(*ec2.Options)) (*ec2.ReleaseHostsOutput, error) {
	return m.ReleaseHostsMock(arg1, arg2, arg3...)
}

// ReleaseIpamPoolAllocation mocked function
func (m EC2ClientMock) ReleaseIpamPoolAllocation(arg1 context.Context, arg2 *ec2.ReleaseIpamPoolAllocationInput, arg3 ...func(*ec2.Options)) (*ec2.ReleaseIpamPoolAllocationOutput, error) {
	return m.ReleaseIpamPoolAllocationMock(arg1, arg2, arg3...)
}

// ReplaceIamInstanceProfileAssociation mocked function
func (m EC2ClientMock) ReplaceIamInstanceProfileAssociation(arg1 context.Context, arg2 *ec2.ReplaceIamInstanceProfileAssociationInput, arg3 ...func(*ec2.Options)) (*ec2.ReplaceIamInstanceProfileAssociationOutput, error) {
	return m.ReplaceIamInstanceProfileAssociationMock(arg1, arg2, arg3...)
}

// ReplaceNetworkAclAssociation mocked function
func (m EC2ClientMock) ReplaceNetworkAclAssociation(arg1 context.Context, arg2 *ec2.ReplaceNetworkAclAssociationInput, arg3 ...func(*ec2.Options)) (*ec2.ReplaceNetworkAclAssociationOutput, error) {
	return m.ReplaceNetworkAclAssociationMock(arg1, arg2, arg3...)
}

// ReplaceNetworkAclEntry mocked function
func (m EC2ClientMock) ReplaceNetworkAclEntry(arg1 context.Context, arg2 *ec2.ReplaceNetworkAclEntryInput, arg3 ...func(*ec2.Options)) (*ec2.ReplaceNetworkAclEntryOutput, error) {
	return m.ReplaceNetworkAclEntryMock(arg1, arg2, arg3...)
}

// ReplaceRoute mocked function
func (m EC2ClientMock) ReplaceRoute(arg1 context.Context, arg2 *ec2.ReplaceRouteInput, arg3 ...func(*ec2.Options)) (*ec2.ReplaceRouteOutput, error) {
	return m.ReplaceRouteMock(arg1, arg2, arg3...)
}

// ReplaceRouteTableAssociation mocked function
func (m EC2ClientMock) ReplaceRouteTableAssociation(arg1 context.Context, arg2 *ec2.ReplaceRouteTableAssociationInput, arg3 ...func(*ec2.Options)) (*ec2.ReplaceRouteTableAssociationOutput, error) {
	return m.ReplaceRouteTableAssociationMock(arg1, arg2, arg3...)
}

// ReplaceTransitGatewayRoute mocked function
func (m EC2ClientMock) ReplaceTransitGatewayRoute(arg1 context.Context, arg2 *ec2.ReplaceTransitGatewayRouteInput, arg3 ...func(*ec2.Options)) (*ec2.ReplaceTransitGatewayRouteOutput, error) {
	return m.ReplaceTransitGatewayRouteMock(arg1, arg2, arg3...)
}

// ReportInstanceStatus mocked function
func (m EC2ClientMock) ReportInstanceStatus(arg1 context.Context, arg2 *ec2.ReportInstanceStatusInput, arg3 ...func(*ec2.Options)) (*ec2.ReportInstanceStatusOutput, error) {
	return m.ReportInstanceStatusMock(arg1, arg2, arg3...)
}

// RequestSpotFleet mocked function
func (m EC2ClientMock) RequestSpotFleet(arg1 context.Context, arg2 *ec2.RequestSpotFleetInput, arg3 ...func(*ec2.Options)) (*ec2.RequestSpotFleetOutput, error) {
	return m.RequestSpotFleetMock(arg1, arg2, arg3...)
}

// RequestSpotInstances mocked function
func (m EC2ClientMock) RequestSpotInstances(arg1 context.Context, arg2 *ec2.RequestSpotInstancesInput, arg3 ...func(*ec2.Options)) (*ec2.RequestSpotInstancesOutput, error) {
	return m.RequestSpotInstancesMock(arg1, arg2, arg3...)
}

// ResetAddressAttribute mocked function
func (m EC2ClientMock) ResetAddressAttribute(arg1 context.Context, arg2 *ec2.ResetAddressAttributeInput, arg3 ...func(*ec2.Options)) (*ec2.ResetAddressAttributeOutput, error) {
	return m.ResetAddressAttributeMock(arg1, arg2, arg3...)
}

// ResetEbsDefaultKmsKeyId mocked function
func (m EC2ClientMock) ResetEbsDefaultKmsKeyId(arg1 context.Context, arg2 *ec2.ResetEbsDefaultKmsKeyIdInput, arg3 ...func(*ec2.Options)) (*ec2.ResetEbsDefaultKmsKeyIdOutput, error) {
	return m.ResetEbsDefaultKmsKeyIdMock(arg1, arg2, arg3...)
}

// ResetFpgaImageAttribute mocked function
func (m EC2ClientMock) ResetFpgaImageAttribute(arg1 context.Context, arg2 *ec2.ResetFpgaImageAttributeInput, arg3 ...func(*ec2.Options)) (*ec2.ResetFpgaImageAttributeOutput, error) {
	return m.ResetFpgaImageAttributeMock(arg1, arg2, arg3...)
}

// ResetImageAttribute mocked function
func (m EC2ClientMock) ResetImageAttribute(arg1 context.Context, arg2 *ec2.ResetImageAttributeInput, arg3 ...func(*ec2.Options)) (*ec2.ResetImageAttributeOutput, error) {
	return m.ResetImageAttributeMock(arg1, arg2, arg3...)
}

// ResetInstanceAttribute mocked function
func (m EC2ClientMock) ResetInstanceAttribute(arg1 context.Context, arg2 *ec2.ResetInstanceAttributeInput, arg3 ...func(*ec2.Options)) (*ec2.ResetInstanceAttributeOutput, error) {
	return m.ResetInstanceAttributeMock(arg1, arg2, arg3...)
}

// ResetNetworkInterfaceAttribute mocked function
func (m EC2ClientMock) ResetNetworkInterfaceAttribute(arg1 context.Context, arg2 *ec2.ResetNetworkInterfaceAttributeInput, arg3 ...func(*ec2.Options)) (*ec2.ResetNetworkInterfaceAttributeOutput, error) {
	return m.ResetNetworkInterfaceAttributeMock(arg1, arg2, arg3...)
}

// ResetSnapshotAttribute mocked function
func (m EC2ClientMock) ResetSnapshotAttribute(arg1 context.Context, arg2 *ec2.ResetSnapshotAttributeInput, arg3 ...func(*ec2.Options)) (*ec2.ResetSnapshotAttributeOutput, error) {
	return m.ResetSnapshotAttributeMock(arg1, arg2, arg3...)
}

// RestoreAddressToClassic mocked function
func (m EC2ClientMock) RestoreAddressToClassic(arg1 context.Context, arg2 *ec2.RestoreAddressToClassicInput, arg3 ...func(*ec2.Options)) (*ec2.RestoreAddressToClassicOutput, error) {
	return m.RestoreAddressToClassicMock(arg1, arg2, arg3...)
}

// RestoreImageFromRecycleBin mocked function
func (m EC2ClientMock) RestoreImageFromRecycleBin(arg1 context.Context, arg2 *ec2.RestoreImageFromRecycleBinInput, arg3 ...func(*ec2.Options)) (*ec2.RestoreImageFromRecycleBinOutput, error) {
	return m.RestoreImageFromRecycleBinMock(arg1, arg2, arg3...)
}

// RestoreManagedPrefixListVersion mocked function
func (m EC2ClientMock) RestoreManagedPrefixListVersion(arg1 context.Context, arg2 *ec2.RestoreManagedPrefixListVersionInput, arg3 ...func(*ec2.Options)) (*ec2.RestoreManagedPrefixListVersionOutput, error) {
	return m.RestoreManagedPrefixListVersionMock(arg1, arg2, arg3...)
}

// RestoreSnapshotFromRecycleBin mocked function
func (m EC2ClientMock) RestoreSnapshotFromRecycleBin(arg1 context.Context, arg2 *ec2.RestoreSnapshotFromRecycleBinInput, arg3 ...func(*ec2.Options)) (*ec2.RestoreSnapshotFromRecycleBinOutput, error) {
	return m.RestoreSnapshotFromRecycleBinMock(arg1, arg2, arg3...)
}

// RestoreSnapshotTier mocked function
func (m EC2ClientMock) RestoreSnapshotTier(arg1 context.Context, arg2 *ec2.RestoreSnapshotTierInput, arg3 ...func(*ec2.Options)) (*ec2.RestoreSnapshotTierOutput, error) {
	return m.RestoreSnapshotTierMock(arg1, arg2, arg3...)
}

// RevokeClientVpnIngress mocked function
func (m EC2ClientMock) RevokeClientVpnIngress(arg1 context.Context, arg2 *ec2.RevokeClientVpnIngressInput, arg3 ...func(*ec2.Options)) (*ec2.RevokeClientVpnIngressOutput, error) {
	return m.RevokeClientVpnIngressMock(arg1, arg2, arg3...)
}

// RevokeSecurityGroupEgress mocked function
func (m EC2ClientMock) RevokeSecurityGroupEgress(arg1 context.Context, arg2 *ec2.RevokeSecurityGroupEgressInput, arg3 ...func(*ec2.Options)) (*ec2.RevokeSecurityGroupEgressOutput, error) {
	return m.RevokeSecurityGroupEgressMock(arg1, arg2, arg3...)
}

// RevokeSecurityGroupIngress mocked function
func (m EC2ClientMock) RevokeSecurityGroupIngress(arg1 context.Context, arg2 *ec2.RevokeSecurityGroupIngressInput, arg3 ...func(*ec2.Options)) (*ec2.RevokeSecurityGroupIngressOutput, error) {
	return m.RevokeSecurityGroupIngressMock(arg1, arg2, arg3...)
}

// RunInstances mocked function
func (m EC2ClientMock) RunInstances(arg1 context.Context, arg2 *ec2.RunInstancesInput, arg3 ...func(*ec2.Options)) (*ec2.RunInstancesOutput, error) {
	return m.RunInstancesMock(arg1, arg2, arg3...)
}

// RunScheduledInstances mocked function
func (m EC2ClientMock) RunScheduledInstances(arg1 context.Context, arg2 *ec2.RunScheduledInstancesInput, arg3 ...func(*ec2.Options)) (*ec2.RunScheduledInstancesOutput, error) {
	return m.RunScheduledInstancesMock(arg1, arg2, arg3...)
}

// SearchLocalGatewayRoutes mocked function
func (m EC2ClientMock) SearchLocalGatewayRoutes(arg1 context.Context, arg2 *ec2.SearchLocalGatewayRoutesInput, arg3 ...func(*ec2.Options)) (*ec2.SearchLocalGatewayRoutesOutput, error) {
	return m.SearchLocalGatewayRoutesMock(arg1, arg2, arg3...)
}

// SearchTransitGatewayMulticastGroups mocked function
func (m EC2ClientMock) SearchTransitGatewayMulticastGroups(arg1 context.Context, arg2 *ec2.SearchTransitGatewayMulticastGroupsInput, arg3 ...func(*ec2.Options)) (*ec2.SearchTransitGatewayMulticastGroupsOutput, error) {
	return m.SearchTransitGatewayMulticastGroupsMock(arg1, arg2, arg3...)
}

// SearchTransitGatewayRoutes mocked function
func (m EC2ClientMock) SearchTransitGatewayRoutes(arg1 context.Context, arg2 *ec2.SearchTransitGatewayRoutesInput, arg3 ...func(*ec2.Options)) (*ec2.SearchTransitGatewayRoutesOutput, error) {
	return m.SearchTransitGatewayRoutesMock(arg1, arg2, arg3...)
}

// SendDiagnosticInterrupt mocked function
func (m EC2ClientMock) SendDiagnosticInterrupt(arg1 context.Context, arg2 *ec2.SendDiagnosticInterruptInput, arg3 ...func(*ec2.Options)) (*ec2.SendDiagnosticInterruptOutput, error) {
	return m.SendDiagnosticInterruptMock(arg1, arg2, arg3...)
}

// StartInstances mocked function
func (m EC2ClientMock) StartInstances(arg1 context.Context, arg2 *ec2.StartInstancesInput, arg3 ...func(*ec2.Options)) (*ec2.StartInstancesOutput, error) {
	return m.StartInstancesMock(arg1, arg2, arg3...)
}

// StartNetworkInsightsAccessScopeAnalysis mocked function
func (m EC2ClientMock) StartNetworkInsightsAccessScopeAnalysis(arg1 context.Context, arg2 *ec2.StartNetworkInsightsAccessScopeAnalysisInput, arg3 ...func(*ec2.Options)) (*ec2.StartNetworkInsightsAccessScopeAnalysisOutput, error) {
	return m.StartNetworkInsightsAccessScopeAnalysisMock(arg1, arg2, arg3...)
}

// StartNetworkInsightsAnalysis mocked function
func (m EC2ClientMock) StartNetworkInsightsAnalysis(arg1 context.Context, arg2 *ec2.StartNetworkInsightsAnalysisInput, arg3 ...func(*ec2.Options)) (*ec2.StartNetworkInsightsAnalysisOutput, error) {
	return m.StartNetworkInsightsAnalysisMock(arg1, arg2, arg3...)
}

// StartVpcEndpointServicePrivateDnsVerification mocked function
func (m EC2ClientMock) StartVpcEndpointServicePrivateDnsVerification(arg1 context.Context, arg2 *ec2.StartVpcEndpointServicePrivateDnsVerificationInput, arg3 ...func(*ec2.Options)) (*ec2.StartVpcEndpointServicePrivateDnsVerificationOutput, error) {
	return m.StartVpcEndpointServicePrivateDnsVerificationMock(arg1, arg2, arg3...)
}

// StopInstances mocked function
func (m EC2ClientMock) StopInstances(arg1 context.Context, arg2 *ec2.StopInstancesInput, arg3 ...func(*ec2.Options)) (*ec2.StopInstancesOutput, error) {
	return m.StopInstancesMock(arg1, arg2, arg3...)
}

// TerminateClientVpnConnections mocked function
func (m EC2ClientMock) TerminateClientVpnConnections(arg1 context.Context, arg2 *ec2.TerminateClientVpnConnectionsInput, arg3 ...func(*ec2.Options)) (*ec2.TerminateClientVpnConnectionsOutput, error) {
	return m.TerminateClientVpnConnectionsMock(arg1, arg2, arg3...)
}

// TerminateInstances mocked function
func (m EC2ClientMock) TerminateInstances(arg1 context.Context, arg2 *ec2.TerminateInstancesInput, arg3 ...func(*ec2.Options)) (*ec2.TerminateInstancesOutput, error) {
	return m.TerminateInstancesMock(arg1, arg2, arg3...)
}

// UnassignIpv6Addresses mocked function
func (m EC2ClientMock) UnassignIpv6Addresses(arg1 context.Context, arg2 *ec2.UnassignIpv6AddressesInput, arg3 ...func(*ec2.Options)) (*ec2.UnassignIpv6AddressesOutput, error) {
	return m.UnassignIpv6AddressesMock(arg1, arg2, arg3...)
}

// UnassignPrivateIpAddresses mocked function
func (m EC2ClientMock) UnassignPrivateIpAddresses(arg1 context.Context, arg2 *ec2.UnassignPrivateIpAddressesInput, arg3 ...func(*ec2.Options)) (*ec2.UnassignPrivateIpAddressesOutput, error) {
	return m.UnassignPrivateIpAddressesMock(arg1, arg2, arg3...)
}

// UnmonitorInstances mocked function
func (m EC2ClientMock) UnmonitorInstances(arg1 context.Context, arg2 *ec2.UnmonitorInstancesInput, arg3 ...func(*ec2.Options)) (*ec2.UnmonitorInstancesOutput, error) {
	return m.UnmonitorInstancesMock(arg1, arg2, arg3...)
}

// UpdateSecurityGroupRuleDescriptionsEgress mocked function
func (m EC2ClientMock) UpdateSecurityGroupRuleDescriptionsEgress(arg1 context.Context, arg2 *ec2.UpdateSecurityGroupRuleDescriptionsEgressInput, arg3 ...func(*ec2.Options)) (*ec2.UpdateSecurityGroupRuleDescriptionsEgressOutput, error) {
	return m.UpdateSecurityGroupRuleDescriptionsEgressMock(arg1, arg2, arg3...)
}

// UpdateSecurityGroupRuleDescriptionsIngress mocked function
func (m EC2ClientMock) UpdateSecurityGroupRuleDescriptionsIngress(arg1 context.Context, arg2 *ec2.UpdateSecurityGroupRuleDescriptionsIngressInput, arg3 ...func(*ec2.Options)) (*ec2.UpdateSecurityGroupRuleDescriptionsIngressOutput, error) {
	return m.UpdateSecurityGroupRuleDescriptionsIngressMock(arg1, arg2, arg3...)
}

// WithdrawByoipCidr mocked function
func (m EC2ClientMock) WithdrawByoipCidr(arg1 context.Context, arg2 *ec2.WithdrawByoipCidrInput, arg3 ...func(*ec2.Options)) (*ec2.WithdrawByoipCidrOutput, error) {
	return m.WithdrawByoipCidrMock(arg1, arg2, arg3...)
}
