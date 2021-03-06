# KubernetesClusterCertificateConfigurationAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | Pointer to **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "kubernetes.ClusterCertificateConfiguration"]
**ObjectType** | Pointer to **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "kubernetes.ClusterCertificateConfiguration"]
**CaCert** | Pointer to **string** | Certificate for the Kubernetes API server. | [optional] 
**CaKey** | Pointer to **string** | Private Key for the Kubernetes API server. | [optional] 
**EtcdCert** | Pointer to **string** | Certificate for the etcd cluster. | [optional] 
**EtcdEncryptionKey** | Pointer to **[]string** |  | [optional] 
**EtcdKey** | Pointer to **string** | Private key for the etcd cluster. | [optional] 
**FrontProxyCert** | Pointer to **string** | Certificate for the front proxy to support Kubernetes API aggregators. | [optional] 
**FrontProxyKey** | Pointer to **string** | Private key for the front proxy to support Kubernetes API aggregators. | [optional] 
**SaPrivateKey** | Pointer to **string** | Service account private key used by Kubernetes Token Controller to sign generated service account tokens. | [optional] 
**SaPublicKey** | Pointer to **string** | Service account public key used by Kubernetes API Server to validate service account tokens generated by the Token Controller. | [optional] 

## Methods

### NewKubernetesClusterCertificateConfigurationAllOf

`func NewKubernetesClusterCertificateConfigurationAllOf(classId string, objectType string, ) *KubernetesClusterCertificateConfigurationAllOf`

NewKubernetesClusterCertificateConfigurationAllOf instantiates a new KubernetesClusterCertificateConfigurationAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesClusterCertificateConfigurationAllOfWithDefaults

`func NewKubernetesClusterCertificateConfigurationAllOfWithDefaults() *KubernetesClusterCertificateConfigurationAllOf`

NewKubernetesClusterCertificateConfigurationAllOfWithDefaults instantiates a new KubernetesClusterCertificateConfigurationAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *KubernetesClusterCertificateConfigurationAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *KubernetesClusterCertificateConfigurationAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *KubernetesClusterCertificateConfigurationAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *KubernetesClusterCertificateConfigurationAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *KubernetesClusterCertificateConfigurationAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *KubernetesClusterCertificateConfigurationAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCaCert

`func (o *KubernetesClusterCertificateConfigurationAllOf) GetCaCert() string`

GetCaCert returns the CaCert field if non-nil, zero value otherwise.

### GetCaCertOk

`func (o *KubernetesClusterCertificateConfigurationAllOf) GetCaCertOk() (*string, bool)`

GetCaCertOk returns a tuple with the CaCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaCert

`func (o *KubernetesClusterCertificateConfigurationAllOf) SetCaCert(v string)`

SetCaCert sets CaCert field to given value.

### HasCaCert

`func (o *KubernetesClusterCertificateConfigurationAllOf) HasCaCert() bool`

HasCaCert returns a boolean if a field has been set.

### GetCaKey

`func (o *KubernetesClusterCertificateConfigurationAllOf) GetCaKey() string`

GetCaKey returns the CaKey field if non-nil, zero value otherwise.

### GetCaKeyOk

`func (o *KubernetesClusterCertificateConfigurationAllOf) GetCaKeyOk() (*string, bool)`

GetCaKeyOk returns a tuple with the CaKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaKey

`func (o *KubernetesClusterCertificateConfigurationAllOf) SetCaKey(v string)`

SetCaKey sets CaKey field to given value.

### HasCaKey

`func (o *KubernetesClusterCertificateConfigurationAllOf) HasCaKey() bool`

HasCaKey returns a boolean if a field has been set.

### GetEtcdCert

`func (o *KubernetesClusterCertificateConfigurationAllOf) GetEtcdCert() string`

GetEtcdCert returns the EtcdCert field if non-nil, zero value otherwise.

### GetEtcdCertOk

`func (o *KubernetesClusterCertificateConfigurationAllOf) GetEtcdCertOk() (*string, bool)`

GetEtcdCertOk returns a tuple with the EtcdCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEtcdCert

`func (o *KubernetesClusterCertificateConfigurationAllOf) SetEtcdCert(v string)`

SetEtcdCert sets EtcdCert field to given value.

### HasEtcdCert

`func (o *KubernetesClusterCertificateConfigurationAllOf) HasEtcdCert() bool`

HasEtcdCert returns a boolean if a field has been set.

### GetEtcdEncryptionKey

`func (o *KubernetesClusterCertificateConfigurationAllOf) GetEtcdEncryptionKey() []string`

GetEtcdEncryptionKey returns the EtcdEncryptionKey field if non-nil, zero value otherwise.

### GetEtcdEncryptionKeyOk

`func (o *KubernetesClusterCertificateConfigurationAllOf) GetEtcdEncryptionKeyOk() (*[]string, bool)`

GetEtcdEncryptionKeyOk returns a tuple with the EtcdEncryptionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEtcdEncryptionKey

`func (o *KubernetesClusterCertificateConfigurationAllOf) SetEtcdEncryptionKey(v []string)`

SetEtcdEncryptionKey sets EtcdEncryptionKey field to given value.

### HasEtcdEncryptionKey

`func (o *KubernetesClusterCertificateConfigurationAllOf) HasEtcdEncryptionKey() bool`

HasEtcdEncryptionKey returns a boolean if a field has been set.

### SetEtcdEncryptionKeyNil

`func (o *KubernetesClusterCertificateConfigurationAllOf) SetEtcdEncryptionKeyNil(b bool)`

 SetEtcdEncryptionKeyNil sets the value for EtcdEncryptionKey to be an explicit nil

### UnsetEtcdEncryptionKey
`func (o *KubernetesClusterCertificateConfigurationAllOf) UnsetEtcdEncryptionKey()`

UnsetEtcdEncryptionKey ensures that no value is present for EtcdEncryptionKey, not even an explicit nil
### GetEtcdKey

`func (o *KubernetesClusterCertificateConfigurationAllOf) GetEtcdKey() string`

GetEtcdKey returns the EtcdKey field if non-nil, zero value otherwise.

### GetEtcdKeyOk

`func (o *KubernetesClusterCertificateConfigurationAllOf) GetEtcdKeyOk() (*string, bool)`

GetEtcdKeyOk returns a tuple with the EtcdKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEtcdKey

`func (o *KubernetesClusterCertificateConfigurationAllOf) SetEtcdKey(v string)`

SetEtcdKey sets EtcdKey field to given value.

### HasEtcdKey

`func (o *KubernetesClusterCertificateConfigurationAllOf) HasEtcdKey() bool`

HasEtcdKey returns a boolean if a field has been set.

### GetFrontProxyCert

`func (o *KubernetesClusterCertificateConfigurationAllOf) GetFrontProxyCert() string`

GetFrontProxyCert returns the FrontProxyCert field if non-nil, zero value otherwise.

### GetFrontProxyCertOk

`func (o *KubernetesClusterCertificateConfigurationAllOf) GetFrontProxyCertOk() (*string, bool)`

GetFrontProxyCertOk returns a tuple with the FrontProxyCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrontProxyCert

`func (o *KubernetesClusterCertificateConfigurationAllOf) SetFrontProxyCert(v string)`

SetFrontProxyCert sets FrontProxyCert field to given value.

### HasFrontProxyCert

`func (o *KubernetesClusterCertificateConfigurationAllOf) HasFrontProxyCert() bool`

HasFrontProxyCert returns a boolean if a field has been set.

### GetFrontProxyKey

`func (o *KubernetesClusterCertificateConfigurationAllOf) GetFrontProxyKey() string`

GetFrontProxyKey returns the FrontProxyKey field if non-nil, zero value otherwise.

### GetFrontProxyKeyOk

`func (o *KubernetesClusterCertificateConfigurationAllOf) GetFrontProxyKeyOk() (*string, bool)`

GetFrontProxyKeyOk returns a tuple with the FrontProxyKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrontProxyKey

`func (o *KubernetesClusterCertificateConfigurationAllOf) SetFrontProxyKey(v string)`

SetFrontProxyKey sets FrontProxyKey field to given value.

### HasFrontProxyKey

`func (o *KubernetesClusterCertificateConfigurationAllOf) HasFrontProxyKey() bool`

HasFrontProxyKey returns a boolean if a field has been set.

### GetSaPrivateKey

`func (o *KubernetesClusterCertificateConfigurationAllOf) GetSaPrivateKey() string`

GetSaPrivateKey returns the SaPrivateKey field if non-nil, zero value otherwise.

### GetSaPrivateKeyOk

`func (o *KubernetesClusterCertificateConfigurationAllOf) GetSaPrivateKeyOk() (*string, bool)`

GetSaPrivateKeyOk returns a tuple with the SaPrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaPrivateKey

`func (o *KubernetesClusterCertificateConfigurationAllOf) SetSaPrivateKey(v string)`

SetSaPrivateKey sets SaPrivateKey field to given value.

### HasSaPrivateKey

`func (o *KubernetesClusterCertificateConfigurationAllOf) HasSaPrivateKey() bool`

HasSaPrivateKey returns a boolean if a field has been set.

### GetSaPublicKey

`func (o *KubernetesClusterCertificateConfigurationAllOf) GetSaPublicKey() string`

GetSaPublicKey returns the SaPublicKey field if non-nil, zero value otherwise.

### GetSaPublicKeyOk

`func (o *KubernetesClusterCertificateConfigurationAllOf) GetSaPublicKeyOk() (*string, bool)`

GetSaPublicKeyOk returns a tuple with the SaPublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaPublicKey

`func (o *KubernetesClusterCertificateConfigurationAllOf) SetSaPublicKey(v string)`

SetSaPublicKey sets SaPublicKey field to given value.

### HasSaPublicKey

`func (o *KubernetesClusterCertificateConfigurationAllOf) HasSaPublicKey() bool`

HasSaPublicKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


